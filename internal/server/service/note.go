package service

import (
	"context"
	"firstpass/internal/server/model"
)

type NoteService interface {
	Create(ctx context.Context, note *model.Note) error
	Update(ctx context.Context, note *model.Note) error
	Delete(ctx context.Context, note *model.Note) error
	GetAll(ctx context.Context, userID int) ([]*model.Note, error)
}

type NoteRepository interface {
	Create(ctx context.Context, note *model.Note) error
	Update(ctx context.Context, note *model.Note) error
	Delete(ctx context.Context, note *model.Note) error
	GetAll(ctx context.Context, userID int) ([]*model.Note, error)
}

type noteService struct {
	note NoteRepository
}

func NewNoteService(note NoteRepository) *noteService {
	return &noteService{
		note: note,
	}
}

func (s *noteService) Create(ctx context.Context, note *model.Note) error {
	return s.note.Create(ctx, note)
}

func (s *noteService) Update(ctx context.Context, note *model.Note) error {
	return s.note.Update(ctx, note)
}

func (s *noteService) Delete(ctx context.Context, note *model.Note) error {
	return s.note.Delete(ctx, note)
}

func (s *noteService) GetAll(ctx context.Context, userID int) ([]*model.Note, error) {
	return s.note.GetAll(ctx, userID)
}
