package repository

import (
	"context"
	"firstpass/internal/server/model"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (n *AccountRepository) Create(ctx context.Context, account *model.Account) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO accounts (user_id, login, password, metadata) VALUES (:user_id, :login, :password, :metadata)", account)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (n *AccountRepository) GetAll(ctx context.Context, userID int) ([]*model.Account, error) {
	var notes []*model.Account
	err := n.db.Select(&notes, "SELECT login, password, metadata FROM accounts WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *AccountRepository) Update(ctx context.Context, account *model.Account) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("UPDATE accounts SET login = :login, password = :password, metadata = :metadata WHERE id = :id", account)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (n *AccountRepository) Delete(ctx context.Context, account *model.Account) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("DELETE FROM accounts WHERE user_id = :user_id AND id = :id", account)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
