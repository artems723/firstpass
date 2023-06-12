package model

type Card struct {
	ID       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	Number   string `json:"number" db:"number"`
	Holder   string `json:"holder" db:"holder"`
	Expire   string `json:"expire" db:"expire"`
	CVC      string `json:"cvc" db:"cvc"`
	Metadata string `json:"comment" db:"comment"`
}
