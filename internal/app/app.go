package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	descAuth "github.com/nqxcode/chat_microservice/pkg/auth_v1"
	"github.com/nqxcode/platform_common/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/nqxcode/chat_microservice/internal/config"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

// App application
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	authClient      descAuth.AuthV1Client
}

// NewApp new application
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run run application
func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.runGRPCServer(ctx)
		if err != nil {
			log.Printf("failed to run GRPC server: %v", err)
		}
	}()

	gracefulShutdown(ctx, cancel, wg)

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initAuthClient,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(configPath)
	if err != nil {
		log.Printf("No %s file found, using environment variables: %v", configPath, err)
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	tlsCert, err := tls.X509KeyPair(a.serviceProvider.GRPCConfig().Cert(), a.serviceProvider.GRPCConfig().Key())
	if err != nil {
		return err
	}
	creds := credentials.NewServerTLSFromCert(&tlsCert)

	a.grpcServer = grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(a.serviceProvider.AuthInterceptor(a.authClient).Intercept))

	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImpl(ctx))

	return nil
}

func (a *App) initAuthClient(ctx context.Context) error {
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(a.serviceProvider.AuthConfig().Cert()) {
		return fmt.Errorf("credentials: failed to append certificates")
	}

	creds := credentials.NewTLS(&tls.Config{ServerName: "", RootCAs: cp, MinVersion: tls.VersionTLS12})

	conn, err := grpc.Dial(a.serviceProvider.AuthConfig().Address(), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to dial Auth client: %v", err)
	}

	a.authClient = descAuth.NewAuthV1Client(conn)

	return nil
}

func (a *App) runGRPCServer(ctx context.Context) error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	go func() {
		<-ctx.Done()
		a.grpcServer.GracefulStop()
		log.Printf("GRPC server gracefully stopped")
	}()

	listener, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-waitSignal():
		log.Println("terminating: via signal")
	}

	cancel()
	if wg != nil {
		wg.Wait()
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}
