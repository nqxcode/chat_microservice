package chat

import (
	"github.com/nqxcode/chat_microservice/internal/client/db"
	"github.com/nqxcode/chat_microservice/internal/repository"
	def "github.com/nqxcode/chat_microservice/internal/service"
)

type service struct {
	chatRepository       repository.ChatRepository
	chatToUserRepository repository.ChatToUserRepository
	messageRepository    repository.MessageRepository
	txManager            db.TxManager
}

func NewService(
	chatRepository repository.ChatRepository,
	chatToUserRepository repository.ChatToUserRepository,
	messageRepository repository.MessageRepository,
	txManager db.TxManager,
) def.ChatService {
	return &service{
		chatRepository:       chatRepository,
		chatToUserRepository: chatToUserRepository,
		messageRepository:    messageRepository,
		txManager:            txManager,
	}
}
