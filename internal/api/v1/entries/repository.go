package entries

import (
	"errors"

	"github.com/daddydemir/notarium/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]domain.Entry, error) {
	var entries []domain.Entry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *Repository) Create(entry domain.Entry) error {
	return r.db.Create(&entry).Error
}

func (r *Repository) GetByID(id string) (domain.Entry, error) {
	var entry domain.Entry
	if err := r.db.Preload("Topics").Preload("Topics.Notes").First(&entry, "id = ?", id).Error; err != nil {
		return entry, err
	}
	return entry, nil
}

func (r *Repository) Update(id string, entry domain.Entry) error {
	if err := r.db.Model(&domain.Entry{}).Where("id = ?", id).Updates(entry).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Entry{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByDate(date string) (domain.Entry, error) {
	var entry domain.Entry
	err := r.db.Preload("Topics").Preload("Topics.Notes").First(&entry, "date = ?", date).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entry, nil
	}
	return entry, err
}
