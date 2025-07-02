package services

import "note_notifications/internal/ports"

type NoteService struct {
	NoteRepository ports.NoteRepository
}