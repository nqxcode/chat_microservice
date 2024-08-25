package chat

import (
	"context"
	"log"

	"github.com/nqxcode/chat_microservice/internal/converter"

	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Get user by id
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get chat: %d", req.GetId())

	chat, err := i.chatService.Get(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.GetResponse{
		Chat: converter.ToChatFromService(chat),
	}, nil
}
