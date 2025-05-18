package notes

import (
	"context"
	"errors"

	"github.com/daddydemir/notarium/internal/domain"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Note, error) {
	return s.repo.GetAll()
}

func (s *Service) Create(ctx context.Context, note domain.Note) error {
	if note.Content == "" {
		return errors.New("note content cannot be empty")
	}
	if note.TopicID == [16]byte{} {
		return errors.New("topic ID is required")
	}
	if note.Version <= 0 {
		note.Version = 1
	}
	return s.repo.Create(note)
}

func (s *Service) GetByID(ctx context.Context, id string) (domain.Note, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(ctx context.Context, id string, note domain.Note) error {
	return s.repo.Update(id, note)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}
