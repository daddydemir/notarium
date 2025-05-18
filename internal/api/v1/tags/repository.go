package tags

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

func (r *Repository) GetAll() ([]domain.Tag, error) {
	var tags []domain.Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *Repository) Create(tag domain.Tag) error {
	return r.db.Create(&tag).Error
}

func (r *Repository) GetByID(id string) (domain.Tag, error) {
	var tag domain.Tag
	if err := r.db.Preload("Entries").First(&tag, "id = ?", id).Error; err != nil {
		return tag, err
	}
	return tag, nil
}

func (r *Repository) Update(id string, tag domain.Tag) error {
	if err := r.db.Model(&domain.Tag{}).Where("id = ?", id).Updates(tag).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByEntryID(entryID string) ([]domain.Tag, error) {
	var tags []domain.Tag
	if err := r.db.Joins("JOIN entry_tags ON entry_tags.tag_id = tags.id").
		Where("entry_tags.entry_id = ?", entryID).
		Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *Repository) AddTagToEntry(entryID, tagID string) error {
	return r.db.Exec("INSERT INTO entry_tags (entry_id, tag_id) VALUES (?, ?)", entryID, tagID).Error
}

func (r *Repository) RemoveTagFromEntry(entryID, tagID string) error {
	return r.db.Exec("DELETE FROM entry_tags WHERE entry_id = ? AND tag_id = ?", entryID, tagID).Error
}
