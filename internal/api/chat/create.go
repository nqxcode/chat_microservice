package chat

import (
	"context"
	"log"

	"github.com/nqxcode/chat_microservice/internal/converter"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

// Create create chat
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create chat: %#v", req.GetInfo())

	id, err := i.chatService.Create(ctx, converter.ToChatInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
