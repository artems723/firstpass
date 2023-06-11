package service

import (
	"context"
	"firstpass/internal/server/model"
)

type AccountService interface {
	Create(ctx context.Context, account *model.Account) error
	GetByName(ctx context.Context, name string) (*model.Account, error)
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, name string) error
}

type AccountRepository interface {
	Create(ctx context.Context, account *model.Account) error
	GetByName(ctx context.Context, name string) (*model.Account, error)
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, name string) error
}

type accountService struct {
	account AccountRepository
}

func NewAccountService(account AccountRepository) *accountService {
	return &accountService{
		account: account,
	}
}
