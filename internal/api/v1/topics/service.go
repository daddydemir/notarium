package topics

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

func (s *Service) GetAll(ctx context.Context) ([]domain.Topic, error) {
	return s.repo.GetAll()
}

func (s *Service) Create(ctx context.Context, topic domain.Topic) error {
	if topic.Name == "" {
		return errors.New("topic name cannot be empty")
	}
	if topic.EntryID == [16]byte{} {
		return errors.New("entry ID is required for a topic")
	}
	return s.repo.Create(topic)
}

func (s *Service) GetByID(ctx context.Context, id string) (domain.Topic, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(ctx context.Context, id string, topic domain.Topic) error {
	return s.repo.Update(id, topic)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}
