package repository

import (
	"context"
	"firstpass/internal/server/model"
	"github.com/jmoiron/sqlx"
)

type NoteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) *NoteRepository {
	return &NoteRepository{db}
}

func (n *NoteRepository) Create(ctx context.Context, note *model.Note) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO notes (user_id, text, metadata) VALUES (:user_id, :text, :metadata)", note)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (n *NoteRepository) GetAll(ctx context.Context, userID int) ([]*model.Note, error) {
	var notes []*model.Note
	err := n.db.Select(&notes, "SELECT text, metadata FROM notes WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *NoteRepository) Update(ctx context.Context, note *model.Note) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("UPDATE notes SET text = :text, metadata = :metadata WHERE id = :id", note)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (n *NoteRepository) Delete(ctx context.Context, note *model.Note) error {
	tx := n.db.MustBegin()
	_, err := tx.NamedExec("DELETE FROM notes WHERE user_id = :user_id AND id = :id", note)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
