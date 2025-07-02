package ports

import (
	"note_notifications/internal/schemas"
)

type NoteRepository interface {
	GenerateUniqueID() (string, error)
	Create(newID string, note *schemas.NoteCreate) (string, error)
	Update(note *schemas.NoteUpdate) error
	Delete(id string) error
	List() (*[]schemas.NoteDTO, error)
	Get(id string) (*schemas.NoteResponse, error)
}

type NoteService interface {
	Create(note *schemas.NoteCreate) (string, error)
	Update(note *schemas.NoteUpdate) error
	Delete(id string) error
	List() (*[]schemas.NoteDTO, error)
	Get(id string) (*schemas.NoteResponse, error)
}