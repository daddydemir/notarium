package files

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

func (r *Repository) GetAll() ([]domain.File, error) {
	var files []domain.File
	if err := r.db.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (r *Repository) Create(file domain.File) error {
	return r.db.Create(&file).Error
}

func (r *Repository) GetByID(id string) (domain.File, error) {
	var file domain.File
	if err := r.db.First(&file, "id = ?", id).Error; err != nil {
		return file, err
	}
	return file, nil
}

func (r *Repository) Update(id string, file domain.File) error {
	if err := r.db.Model(&domain.File{}).Where("id = ?", id).Updates(file).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.File{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByNoteID(noteID string) ([]domain.File, error) {
	var files []domain.File
	if err := r.db.Where("note_id = ?", noteID).Order("created_at ASC").Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}
