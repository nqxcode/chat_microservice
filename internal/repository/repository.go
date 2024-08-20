package repository

import (
	"context"

	"github.com/nqxcode/chat_microservice/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, model *model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*model.Chat, error)
}

type ChatToUserRepository interface {
	Create(ctx context.Context, model *model.ChatToUser) (int64, error)
	Get(ctx context.Context, chatID int64, limit Limit) ([]model.ChatToUser, error)
	DeleteByChatID(ctx context.Context, chatID int64) error
}

type MessageRepository interface {
	Create(ctx context.Context, model *model.Message) (int64, error)
	Get(ctx context.Context, chatID int64, limit Limit) ([]model.Message, error)
	DeleteByChatID(ctx context.Context, chatID int64) error
}

type LogRepository interface {
	Create(ctx context.Context, model *model.Log) error
}
