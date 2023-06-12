package repository

import (
	"context"
	"firstpass/internal/server/model"
	"github.com/jmoiron/sqlx"
)

type BinRepository struct {
	db *sqlx.DB
}

func NewBinRepository(db *sqlx.DB) *BinRepository {
	return &BinRepository{db}
}

func (n *BinRepository) Create(ctx context.Context, bin *model.Bin) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO binary (user_id, body, metadata) VALUES (:user_id, :body, :metadata)", bin)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (n *BinRepository) GetAll(ctx context.Context, userID int) ([]*model.Bin, error) {
	var notes []*model.Bin
	err := n.db.Select(&notes, "SELECT body, metadata FROM binary WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *BinRepository) Update(ctx context.Context, bin *model.Bin) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("UPDATE binary SET body = :body, metadata = :metadata WHERE id = :id", bin)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (n *BinRepository) Delete(ctx context.Context, bin *model.Bin) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("DELETE FROM binary WHERE user_id = :user_id AND id = :id", bin)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
