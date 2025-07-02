package services

import "note_notifications/internal/schemas"

func (s *NoteService) Create(note *schemas.NoteCreate) (string, error) {
	newID, err := s.NoteRepository.GenerateUniqueID()
	if err != nil {
		return "", err
	}
	return s.NoteRepository.Create(newID, note)
}

func Update(note *schemas.NoteUpdate) error {
	return nil
}

func Delete(id string) error {
	return nil
}

func List() (*[]schemas.NoteDTO, error) {
	return nil, nil
}

func Get(id string) (*schemas.NoteResponse, error) {
	return nil, nil
}