package entries

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
	if err := r.db.First(&entry, "id = ?", id).Error; err != nil {
		return entry, err
	}
	return entry, nil
}
