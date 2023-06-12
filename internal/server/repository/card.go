package repository

import (
	"context"
	"firstpass/internal/server/model"
	"github.com/jmoiron/sqlx"
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{db}
}

func (n *CardRepository) Create(ctx context.Context, card *model.Card) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO cards (user_id, number, holder, expire, cvc, metadata) VALUES (:user_id, :number, :holder, :expire, :cvc, :metadata)", card)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (n *CardRepository) GetAll(ctx context.Context, userID int) ([]*model.Card, error) {
	var notes []*model.Card
	err := n.db.Select(&notes, "SELECT number, holder, expire, cvc, metadata FROM cards WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *CardRepository) Update(ctx context.Context, card *model.Card) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("UPDATE cards SET number = :number, holder = :holder, expire = :expire, cvc = :cvc, metadata = :metadata WHERE id = :id", card)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (n *CardRepository) Delete(ctx context.Context, card *model.Card) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("DELETE FROM cards WHERE user_id = :user_id AND id = :id", card)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
