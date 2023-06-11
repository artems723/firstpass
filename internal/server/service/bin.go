package service

import (
	"context"
	"firstpass/internal/server/model"
)

type BinService interface {
	Create(ctx context.Context, bin *model.Bin) error
	GetByName(ctx context.Context, name string) (*model.Bin, error)
	Update(ctx context.Context, bin *model.Bin) error
	Delete(ctx context.Context, name string) error
}

type BinRepository interface {
	Create(ctx context.Context, bin *model.Bin) error
	GetByName(ctx context.Context, name string) (*model.Bin, error)
	Update(ctx context.Context, bin *model.Bin) error
	Delete(ctx context.Context, name string) error
}

type binService struct {
	bin BinRepository
}

func NewBinService(bin BinRepository) *binService {
	return &binService{
		bin: bin,
	}
}
