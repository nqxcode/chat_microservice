package model

import "time"

type ChatToUser struct {
	ID        int64
	ChatID    int64
	UserID    int64
	CreatedAt time.Time
}
