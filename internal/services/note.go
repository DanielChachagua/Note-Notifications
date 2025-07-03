package services

import "note_notifications/internal/schemas"

func (s *NoteService) Create(note *schemas.NoteCreate) (string, error) {
	newID, err := s.NoteRepository.GenerateUniqueID()
	if err != nil {
		return "", err
	}
	return s.NoteRepository.Create(newID, note)
}

func (s *NoteService) Update(note *schemas.NoteUpdate) error {
	return s.NoteRepository.Update(note)
}

func (s *NoteService) Delete(id string) error {
	return s.NoteRepository.Delete(id)
}

func (s *NoteService) List() (*[]schemas.NoteDTO, error) {
	return s.NoteRepository.List()
}

func (s *NoteService) Get(id string) (*schemas.NoteResponse, error) {
	return s.NoteRepository.Get(id)
}

func (s *NoteService) GetAllWarn() (*[]schemas.NoteResponse, error) {
	return s.NoteRepository.GetAllWarn()
}