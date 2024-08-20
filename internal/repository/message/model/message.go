package model

import (
	"database/sql"
	"time"
)

// Message repository model
type Message struct {
	ID        int64 `db:"message_id"`
	ChatID    int64
	From      string
	Message   string
	SentAt    sql.NullTime
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
