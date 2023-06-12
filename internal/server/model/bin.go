package model

type Bin struct {
	ID       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	Body     []byte `json:"body" db:"body"`
	Metadata string `json:"comment" db:"comment"`
}
