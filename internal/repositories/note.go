package repositories

import (
	"fmt"
	"math/rand"
	"note_notifications/internal/models"
	"note_notifications/internal/schemas"
)

func (r *Repository) Create(newID string, note *schemas.NoteCreate) (string, error) {
	if err := r.DB.Create(&models.Note{
		ID:          newID,
		Title:       note.Title,
		Description: note.Description,
		Url:         note.Url,
		Date:        note.Date.ToTime(),
		Time:        note.Time.ToTime(),
	}).Error; err != nil {
		return "", err
	}

	fmt.Printf("Nota creada con ID: %s\n", newID)
	return newID, nil
}

func (r *Repository) Update(note *schemas.NoteUpdate) error {
	return nil
}

func (r *Repository) Delete(id string) error {
	return nil
}

func (r *Repository) List() (*[]schemas.NoteDTO, error) {
	return nil, nil
}

func (r *Repository) Get(id string) (*schemas.NoteResponse, error) {
	return nil, nil
}

func (r *Repository) GenerateUniqueID() (string, error) {
	maxAttempts := 10
	for i := 0; i < maxAttempts; i++ {
		id := fmt.Sprintf("%06d", rand.Intn(1000000)) // 000000 a 999999

		var count int64
		if err := r.DB.Model(&models.Note{}).Where("id = ?", id).Count(&count).Error; err != nil {
			return "", err
		}

		if count == 0 {
			return id, nil
		}
	}
	return "", fmt.Errorf("no se pudo generar un ID único tras %d intentos", maxAttempts)
}