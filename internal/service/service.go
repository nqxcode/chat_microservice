package service

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
}
