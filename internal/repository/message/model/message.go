package model

import (
	"database/sql"
	"time"
)

type Message struct {
	ID        int64
	ChatID    int64
	From      string
	Message   string
	SentAt    sql.NullTime
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
