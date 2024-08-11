package model

import (
	"database/sql"
	"time"
)

type Chat struct {
	ID        int64
	Info      ChatInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type ChatInfo struct {
	Name    string
	UserIDs []int64
}
