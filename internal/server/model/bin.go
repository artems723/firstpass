package model

import (
	"database/sql"
	"time"
)

type Bin struct {
	ID        string       `json:"id" db:"id"`
	UserID    string       `json:"user_id" db:"user_id"`
	Name      string       `json:"name" db:"name"`
	Body      []byte       `json:"body" db:"body"`
	Comment   string       `json:"comment" db:"comment"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}
