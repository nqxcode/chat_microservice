package chat

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

// Delete delete chat
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("Delete chat: %d", req.GetId())

	err := i.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}
