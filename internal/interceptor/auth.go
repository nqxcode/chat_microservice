package interceptor

import (
	"context"
	descAccess "github.com/nqxcode/chat_microservice/pkg/auth_v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

const authPrefix = "Bearer "

// AuthInterceptor auth interceptor for requests
type AuthInterceptor interface {
	Intercept(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
}

type authInterceptor struct {
	authClient descAccess.AuthV1Client
}

func NewAuthInterceptor(authClient descAccess.AuthV1Client) *authInterceptor {
	return &authInterceptor{
		authClient: authClient,
	}
}

// Intercept request
func (ai *authInterceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return false, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return false, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	md = metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := ai.authClient.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: info.FullMethod,
	})
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
