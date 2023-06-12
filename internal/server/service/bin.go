package service

import (
	"context"
	"firstpass/internal/server/model"
)

type BinService interface {
	Create(ctx context.Context, bin *model.Bin) error
	Update(ctx context.Context, bin *model.Bin) error
	Delete(ctx context.Context, bin *model.Bin) error
	GetAll(ctx context.Context, userID int) ([]*model.Bin, error)
}

type BinRepository interface {
	Create(ctx context.Context, bin *model.Bin) error
	Update(ctx context.Context, bin *model.Bin) error
	Delete(ctx context.Context, bin *model.Bin) error
	GetAll(ctx context.Context, userID int) ([]*model.Bin, error)
}

type binService struct {
	bin BinRepository
}

func NewBinService(bin BinRepository) *binService {
	return &binService{
		bin: bin,
	}
}

func (s *binService) Create(ctx context.Context, bin *model.Bin) error {
	return s.bin.Create(ctx, bin)
}

func (s *binService) Update(ctx context.Context, bin *model.Bin) error {
	return s.bin.Update(ctx, bin)
}

func (s *binService) Delete(ctx context.Context, bin *model.Bin) error {
	return s.bin.Delete(ctx, bin)
}

func (s *binService) GetAll(ctx context.Context, userID int) ([]*model.Bin, error) {
	return s.bin.GetAll(ctx, userID)
}
