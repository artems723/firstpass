package service

import (
	"context"
	"firstpass/internal/server/model"
)

type NoteService interface {
	Create(ctx context.Context, note *model.Note) error
	GetByName(ctx context.Context, name string) (*model.Note, error)
	Update(ctx context.Context, note *model.Note) error
	Delete(ctx context.Context, name string) error
}

type NoteRepository interface {
	Create(ctx context.Context, note *model.Note) error
	GetByName(ctx context.Context, name string) (*model.Note, error)
	Update(ctx context.Context, note *model.Note) error
	Delete(ctx context.Context, name string) error
}

type noteService struct {
	note NoteRepository
}

func NewNoteService(note NoteRepository) *noteService {
	return &noteService{
		note: note,
	}
}
