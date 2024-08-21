package service

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/pagination"
)

// ChatService interface
type ChatService interface {
	Create(ctx context.Context, chat *model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*model.Chat, error)
	SendMessage(ctx context.Context, message *model.Message) (int64, error)
	GetMessages(ctx context.Context, id int64, limit *pagination.Limit) ([]model.Message, error)
}

// LogService interface
type LogService interface {
	Create(ctx context.Context, message *model.Log) error
}
