package chat

import (
	"github.com/nqxcode/chat_microservice/internal/service"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

// Implementation chat api implementation
type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewImplementation new chat service implementation
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
