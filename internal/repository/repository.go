package repository

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/platform_common/pagination"
)

// ChatRepository chat repository
type ChatRepository interface {
	Create(ctx context.Context, model *model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*model.Chat, error)
}

// ChatToUserRepository chat to user relation repository
type ChatToUserRepository interface {
	Create(ctx context.Context, model *model.ChatToUser) (int64, error)
	Get(ctx context.Context, chatID int64, limit pagination.Limit) ([]model.ChatToUser, error)
	DeleteByChatID(ctx context.Context, chatID int64) error
}

// MessageRepository message repository
type MessageRepository interface {
	Create(ctx context.Context, model *model.Message) (int64, error)
	Get(ctx context.Context, chatID int64, limit *pagination.Limit) ([]model.Message, error)
	DeleteByChatID(ctx context.Context, chatID int64) error
}

// LogRepository log repository
type LogRepository interface {
	Create(ctx context.Context, model *model.Log) error
}
