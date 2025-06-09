package entries

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

func (s *Service) GetAll(ctx context.Context) ([]domain.Entry, error) {
	return s.repo.GetAll()
}

func (s *Service) Create(ctx context.Context, entry domain.Entry) error {
	if entry.Title == "" {
		return errors.New("title cannot be empty")
	}
	return s.repo.Create(entry)
}

func (s *Service) GetByID(ctx context.Context, id string) (domain.Entry, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(ctx context.Context, id string, entry domain.Entry) error {
	return s.repo.Update(id, entry)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}

func (s *Service) GetByDate(ctx context.Context, date string) (domain.Entry, error) {
	return s.repo.GetByDate(date)
}
