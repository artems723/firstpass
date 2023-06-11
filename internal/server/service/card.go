package service

import (
	"context"
	"firstpass/internal/server/model"
)

type CardService interface {
	Create(ctx context.Context, card *model.Card) error
	GetByName(ctx context.Context, name string) (*model.Card, error)
	Update(ctx context.Context, card *model.Card) error
	Delete(ctx context.Context, name string) error
}

type CardRepository interface {
	Create(ctx context.Context, card *model.Card) error
	GetByName(ctx context.Context, name string) (*model.Card, error)
	Update(ctx context.Context, card *model.Card) error
	Delete(ctx context.Context, name string) error
}

type cardService struct {
	card CardRepository
}

func NewCardService(card CardRepository) *cardService {
	return &cardService{
		card: card,
	}
}
