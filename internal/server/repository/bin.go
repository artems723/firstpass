package repository

import "github.com/jmoiron/sqlx"

type BinRepository struct {
	db *sqlx.DB
}

func NewBinRepository(db *sqlx.DB) *BinRepository {
	return &BinRepository{db}
}
