package repository

import (
	"context"

	"github.com/nqxcode/chat_microservice/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, info *model.ChatInfo) (int64, error)
}
