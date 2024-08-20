package model

import "time"

// ChatToUser repository relation model
type ChatToUser struct {
	ID        int64 `db:"chat_to_user_id"`
	ChatID    int64
	UserID    int64
	CreatedAt time.Time
}
