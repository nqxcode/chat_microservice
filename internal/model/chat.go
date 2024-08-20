package model

import (
	"database/sql"
	"time"
)

// Chat model
type Chat struct {
	ID        int64
	Info      ChatInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// ChatInfo model
type ChatInfo struct {
	Name    string
	UserIDs []int64
}
