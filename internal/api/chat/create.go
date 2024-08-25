package chat

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"

	"github.com/nqxcode/chat_microservice/internal/converter"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"google.golang.org/grpc/status"
)

// Create create chat
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create chat: %#v", req.GetInfo())

	id, err := i.chatService.Create(ctx, converter.ToChatInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
