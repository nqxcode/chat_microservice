package app

import (
	"context"
	"log"

	"github.com/nqxcode/platform_common/client/db"
	"github.com/nqxcode/platform_common/client/db/pg"
	"github.com/nqxcode/platform_common/client/db/transaction"
	"github.com/nqxcode/platform_common/closer"

	"github.com/nqxcode/chat_microservice/internal/api/chat"
	"github.com/nqxcode/chat_microservice/internal/config"
	"github.com/nqxcode/chat_microservice/internal/repository"
	chatRepository "github.com/nqxcode/chat_microservice/internal/repository/chat"
	chatToUserRepository "github.com/nqxcode/chat_microservice/internal/repository/chat_to_user"
	logRepository "github.com/nqxcode/chat_microservice/internal/repository/log"
	messageRepository "github.com/nqxcode/chat_microservice/internal/repository/message"
	"github.com/nqxcode/chat_microservice/internal/service"
	chatService "github.com/nqxcode/chat_microservice/internal/service/chat"
	logService "github.com/nqxcode/chat_microservice/internal/service/log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient             db.Client
	txManager            db.TxManager
	chatRepository       repository.ChatRepository
	chatToUserRepository repository.ChatToUserRepository
	messageRepository    repository.MessageRepository
	logRepository        repository.LogRepository

	logService  service.LogService
	chatService service.ChatService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatToUserRepository(ctx context.Context) repository.ChatToUserRepository {
	if s.chatToUserRepository == nil {
		s.chatToUserRepository = chatToUserRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatToUserRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.DBClient(ctx))
	}

	return s.logRepository
}

func (s *serviceProvider) LogService(ctx context.Context) service.LogService {
	if s.logService == nil {
		s.logService = logService.NewService(
			s.LogRepository(ctx),
		)
	}

	return s.logService
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.ChatToUserRepository(ctx),
			s.MessageRepository(ctx),
			s.LogService(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
