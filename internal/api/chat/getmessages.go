package chat

import (
	"context"
	"log"

	"github.com/nqxcode/chat_microservice/internal/converter"

	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetMessages get messages
func (i *Implementation) GetMessages(ctx context.Context, req *desc.GetMessagesRequest) (*desc.GetMessagesResponse, error) {
	log.Printf("Get messages for chat: %d", req.GetChatId())

	messages, err := i.chatService.GetMessages(ctx, req.GetChatId(), converter.ToLimitFromDesc(req.GetLimit()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.GetMessagesResponse{
		Message: converter.ToMessagesFromService(messages),
	}, nil
}
