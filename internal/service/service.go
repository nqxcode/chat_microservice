package service

import (
	"context"

	"github.com/nqxcode/chat_microservice/internal/model"
)

// ChatService interface
type ChatService interface {
	Create(ctx context.Context, chat *model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message *model.Message) (int64, error)
}

// LogService interface
type LogService interface {
	Create(ctx context.Context, message *model.Log) error
}
