package topics

import (
	"github.com/daddydemir/notarium/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]domain.Topic, error) {
	var topics []domain.Topic
	if err := r.db.Preload("Notes").Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *Repository) Create(topic domain.Topic) error {
	return r.db.Create(&topic).Error
}

func (r *Repository) GetByID(id string) (domain.Topic, error) {
	var topic domain.Topic
	if err := r.db.Preload("Notes").First(&topic, "id = ?", id).Error; err != nil {
		return topic, err
	}
	return topic, nil
}

func (r *Repository) Update(id string, topic domain.Topic) error {
	if err := r.db.Model(&domain.Topic{}).Where("id = ?", id).Updates(topic).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Topic{}).Error; err != nil {
		return err
	}
	return nil
}
