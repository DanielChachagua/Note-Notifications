package dependencies

import (
	"note_notifications/internal/repositories"
	"note_notifications/internal/services"

	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB
	Services struct {
		Note *services.NoteService
	}
	Repositories struct {
		Note *repositories.Repository
	}
}

func NewContainer(db *gorm.DB) *Container {
	container := &Container{
		DB: db,
	}

	container.Repositories.Note = &repositories.Repository{DB: db}
	container.Services.Note = &services.NoteService{
		NoteRepository: container.Repositories.Note,
	}

	return container
}