package chat

import (
	"github.com/nqxcode/chat_microservice/internal/repository"
	def "github.com/nqxcode/chat_microservice/internal/service"

	"github.com/nqxcode/platform_common/client/db"
)

type service struct {
	chatRepository       repository.ChatRepository
	chatToUserRepository repository.ChatToUserRepository
	messageRepository    repository.MessageRepository
	auditLogService      def.AuditLogService
	txManager            db.TxManager
}

// NewService new chat service
func NewService(
	chatRepository repository.ChatRepository,
	chatToUserRepository repository.ChatToUserRepository,
	messageRepository repository.MessageRepository,
	auditLogService def.AuditLogService,
	txManager db.TxManager,
) def.ChatService {
	return &service{
		chatRepository:       chatRepository,
		chatToUserRepository: chatToUserRepository,
		messageRepository:    messageRepository,
		auditLogService:      auditLogService,
		txManager:            txManager,
	}
}
