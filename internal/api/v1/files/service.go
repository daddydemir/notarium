package files

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

func (s *Service) GetAll(ctx context.Context) ([]domain.File, error) {
	return s.repo.GetAll()
}

func (s *Service) Create(ctx context.Context, file domain.File) error {
	if file.NoteID == [16]byte{} {
		return errors.New("note ID is required")
	}
	if file.URL == "" {
		return errors.New("file URL cannot be empty")
	}
	if file.FileName == "" {
		return errors.New("file name cannot be empty")
	}
	if file.MimeType == "" {
		return errors.New("MIME type is required")
	}
	return s.repo.Create(file)
}

func (s *Service) GetByID(ctx context.Context, id string) (domain.File, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(ctx context.Context, id string, file domain.File) error {
	return s.repo.Update(id, file)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}
