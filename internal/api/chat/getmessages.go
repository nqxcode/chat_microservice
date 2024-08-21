package chat

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/converter"
	"log"

	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

// GetMessages get messages
func (i *Implementation) GetMessages(ctx context.Context, req *desc.GetMessagesRequest) (*desc.GetMessagesResponse, error) {
	log.Printf("Get messages for chat: %d", req.GetChatId())

	messages, err := i.chatService.GetMessages(ctx, req.GetChatId(), converter.ToLimitFromDesc(req.GetLimit()))
	if err != nil {
		return nil, err
	}

	return &desc.GetMessagesResponse{
		Message: converter.ToMessagesFromService(messages),
	}, nil
}
