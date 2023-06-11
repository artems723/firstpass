package model

import (
	"database/sql"
	"time"
)

type Card struct {
	ID        string       `json:"id" db:"id"`
	UserID    string       `json:"user_id" db:"user_id"`
	Name      string       `json:"name" db:"name"`
	Number    string       `json:"number" db:"number"`
	Holder    string       `json:"holder" db:"holder"`
	Expire    string       `json:"expire" db:"expire"`
	CVC       string       `json:"cvc" db:"cvc"`
	Comment   string       `json:"comment" db:"comment"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}
