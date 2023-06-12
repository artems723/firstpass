package service

import (
	"context"
	"firstpass/internal/server/model"
)

type CardService interface {
	Create(ctx context.Context, card *model.Card) error
	Update(ctx context.Context, card *model.Card) error
	Delete(ctx context.Context, card *model.Card) error
	GetAll(ctx context.Context, userID int) ([]*model.Card, error)
}

type CardRepository interface {
	Create(ctx context.Context, card *model.Card) error
	Update(ctx context.Context, card *model.Card) error
	Delete(ctx context.Context, card *model.Card) error
	GetAll(ctx context.Context, userID int) ([]*model.Card, error)
}

type cardService struct {
	card CardRepository
}

func NewCardService(card CardRepository) *cardService {
	return &cardService{
		card: card,
	}
}

func (s *cardService) Create(ctx context.Context, card *model.Card) error {
	return s.card.Create(ctx, card)
}

func (s *cardService) Update(ctx context.Context, card *model.Card) error {
	return s.card.Update(ctx, card)
}

func (s *cardService) Delete(ctx context.Context, card *model.Card) error {
	return s.card.Delete(ctx, card)
}

func (s *cardService) GetAll(ctx context.Context, userID int) ([]*model.Card, error) {
	return s.card.GetAll(ctx, userID)
}
