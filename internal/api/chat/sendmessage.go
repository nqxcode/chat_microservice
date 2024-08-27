package chat

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nqxcode/chat_microservice/internal/converter"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

// SendMessage send message
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("Send message: %#v", req.GetInfo())

	_, err := i.chatService.SendMessage(ctx, converter.ToMessageFromDesc(req.GetInfo()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
