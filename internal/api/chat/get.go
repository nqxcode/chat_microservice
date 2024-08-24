package chat

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/converter"
	"log"

	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Get user by id
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get chat: %d", req.GetId())

	chat, err := i.chatService.Get(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Chat not found")
	}

	return &desc.GetResponse{
		Chat: converter.ToChatFromService(chat),
	}, nil
}
