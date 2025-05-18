package notes

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

func (r *Repository) GetAll() ([]domain.Note, error) {
	var notes []domain.Note
	if err := r.db.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *Repository) Create(note domain.Note) error {
	return r.db.Create(&note).Error
}

func (r *Repository) GetByID(id string) (domain.Note, error) {
	var note domain.Note
	if err := r.db.First(&note, "id = ?", id).Error; err != nil {
		return note, err
	}
	return note, nil
}

func (r *Repository) Update(id string, note domain.Note) error {
	if err := r.db.Model(&domain.Note{}).Where("id = ?", id).Updates(note).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Note{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByTopicID(topicID string) ([]domain.Note, error) {
	var notes []domain.Note
	if err := r.db.Where("topic_id = ?", topicID).Order("created_at ASC").Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}
