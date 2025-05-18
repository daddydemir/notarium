package entries

import (
	"errors"

	"github.com/daddydemir/notarium/internal/domain"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAllEntries() ([]domain.Entry, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateEntry(entry domain.Entry) error {
	if entry.Title == "" {
		return errors.New("title cannot be empty")
	}
	return s.repo.Create(entry)
}
