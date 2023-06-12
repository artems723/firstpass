package model

type Note struct {
	ID       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	Text     string `json:"text" db:"text"`
	Metadata string `json:"comment" db:"comment"`
}
