package model

import (
	"database/sql"
	"time"
)

// Chat repository model
type Chat struct {
	ID        int64 `db:"chat_id"`
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
