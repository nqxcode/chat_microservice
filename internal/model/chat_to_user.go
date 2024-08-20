package model

import (
	"time"
)

// ChatToUser relation model
type ChatToUser struct {
	ID        int64
	ChatID    int64
	UserID    int64
	CreatedAt time.Time
}
