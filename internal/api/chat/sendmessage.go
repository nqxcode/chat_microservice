package chat

import (
	"context"
	"log"

	"github.com/nqxcode/chat_microservice/internal/converter"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"github.com/golang/protobuf/ptypes/empty"
)

// SendMessage send message
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("Send message: %#v", req.GetInfo())

	_, err := i.chatService.SendMessage(ctx, converter.ToMessageFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
