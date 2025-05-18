package tags

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

func (s *Service) GetAll(ctx context.Context) ([]domain.Tag, error) {
	return s.repo.GetAll()
}

func (s *Service) Create(ctx context.Context, tag domain.Tag) error {
	if tag.Name == "" {
		return errors.New("tag name cannot be empty")
	}

	if len(tag.Color) > 0 && tag.Color[0] != '#' {
		return errors.New("tag color must start with '#'")
	}
	return s.repo.Create(tag)
}

func (s *Service) GetByID(ctx context.Context, id string) (domain.Tag, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(ctx context.Context, id string, tag domain.Tag) error {
	return s.repo.Update(id, tag)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}
