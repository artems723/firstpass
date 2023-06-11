package model

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        string       `json:"id" db:"id"`
	UserID    string       `json:"user_id" db:"user_id"`
	Name      string       `json:"name" db:"name"`
	Login     string       `json:"login" db:"login"`
	Password  string       `json:"password" db:"password"`
	Comment   string       `json:"comment" db:"comment"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}
