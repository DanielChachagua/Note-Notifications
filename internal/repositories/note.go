package repositories

import (
	"errors"
	"fmt"
	"math/rand"
	"note_notifications/internal/models"
	"note_notifications/internal/schemas"
	"time"

	"gorm.io/gorm"
)

func (r *Repository) Create(newID string, note *schemas.NoteCreate) (string, error) {
	if err := r.DB.Create(&models.Note{
		ID:          newID,
		Title:       note.Title,
		Description: note.Description,
		Url:         note.Url,
		Date:        note.Date.ToTime(),
		Time:        note.Time.ToTime(),
		Warn:        note.Warn,
	}).Error; err != nil {
		return "", err
	}

	return newID, nil
}

func (r *Repository) Update(note *schemas.NoteUpdate) error {
	var noteToUpdate models.Note
	if err := r.DB.Where("id = ?", note.ID).First(&noteToUpdate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("no se encontró la nota")
		}
		return fmt.Errorf("error interno al buscar la nota")
	}

	noteToUpdate.Title = note.Title
	noteToUpdate.Description = note.Description
	noteToUpdate.Url = note.Url
	noteToUpdate.Date = note.Date.ToTime()
	noteToUpdate.Time = note.Time.ToTime()

	if note.Warn != nil {
		noteToUpdate.Warn = *note.Warn
	}

	if err := r.DB.Save(&noteToUpdate).Error; err != nil {
		return fmt.Errorf("error al guardar la nota actualizada: %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Note{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return  fmt.Errorf("no se encontró la nota")
		}
		return  fmt.Errorf("error interno al eliminar la nota")
	}
	return nil
}

func (r *Repository) List() (*[]schemas.NoteDTO, error) {
	var notes []models.Note
	if err := r.DB.Find(&notes).Error; err != nil {
		return nil, fmt.Errorf("error interno al buscar la nota")
	}

	listNote := make([]schemas.NoteDTO, len(notes))
	for i, note := range notes {
		listNote[i] = schemas.NoteDTO{
			ID: note.ID,
			Title: note.Title,
			Date: schemas.CustomDate(note.Date),
			Time: schemas.CustomTime(note.Time),
			Warn: note.Warn,
		}
	}

	return &listNote, nil
}

func (r *Repository) Get(id string) (*schemas.NoteResponse, error) {
	var note models.Note
	if err := r.DB.Model(&models.Note{}).
		Where("id = ?", id).
		First(&note).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return  nil, fmt.Errorf("no se encontró la nota")
		}
		return  nil, fmt.Errorf("error interno al buscar la nota")
	}

	noteResp := schemas.NoteResponse{
		ID: note.ID,
		Title: note.Title,
		Description: note.Description,
		Url: note.Url,
		Date: schemas.CustomDate(note.Date),
		Time: schemas.CustomTime(note.Time),
		Warn: note.Warn,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}


	return &noteResp, nil
}

func (r *Repository) GetAllWarn() (*[]schemas.NoteResponse, error) {
	var notes []models.Note

	today := time.Now().Truncate(24 * time.Hour)
	sevenDaysLater := today.AddDate(0, 0, 7)


	if err := r.DB.
		Where("warn = ? AND date BETWEEN ? AND ?", true, today, sevenDaysLater).
		Order("date ASC").
		Order("time ASC").
		Limit(10).
		Find(&notes).Error; err != nil {
		return nil, fmt.Errorf("error interno al buscar la nota")
	}

	listNote := make([]schemas.NoteResponse, len(notes))
	for i, note := range notes {
		listNote[i] = schemas.NoteResponse{
			ID: note.ID,
			Title: note.Title,
			Description: note.Description,
			Url: note.Url,
			Date: schemas.CustomDate(note.Date),
			Time: schemas.CustomTime(note.Time),
			Warn: note.Warn,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		}
	}

	return &listNote, nil
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