package chat

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/converter"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	_, err := i.chatService.SendMessage(ctx, converter.ToMessageFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
