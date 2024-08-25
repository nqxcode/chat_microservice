package chat

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"google.golang.org/grpc/status"
)

// Delete delete chat
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("Delete chat: %d", req.GetId())

	err := i.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
