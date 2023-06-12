package service

import (
	"context"
	"firstpass/internal/server/model"
)

type AccountService interface {
	Create(ctx context.Context, account *model.Account) error
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, account *model.Account) error
	GetAll(ctx context.Context, userID int) ([]*model.Account, error)
}

type AccountRepository interface {
	Create(ctx context.Context, account *model.Account) error
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, account *model.Account) error
	GetAll(ctx context.Context, userID int) ([]*model.Account, error)
}

type accountService struct {
	account AccountRepository
}

func NewAccountService(account AccountRepository) *accountService {
	return &accountService{
		account: account,
	}
}

func (s *accountService) Create(ctx context.Context, account *model.Account) error {
	return s.account.Create(ctx, account)
}

func (s *accountService) Update(ctx context.Context, account *model.Account) error {
	return s.account.Update(ctx, account)
}

func (s *accountService) Delete(ctx context.Context, account *model.Account) error {
	return s.account.Delete(ctx, account)
}

func (s *accountService) GetAll(ctx context.Context, userID int) ([]*model.Account, error) {
	return s.account.GetAll(ctx, userID)
}
