package domain

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Topic struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    EntryID     uuid.UUID `gorm:"type:uuid;not null"`
    Name        string    `gorm:"type:varchar(100);not null"`
    IsStarred   bool      `gorm:"default:false;not null"`
    IsCompleted bool      `gorm:"default:false;not null"`
    Order       int       `gorm:"default:0"`
    CreatedAt   time.Time `gorm:"type:timestamptz;not null;default:now()"`
    UpdatedAt   time.Time `gorm:"type:timestamptz"`

    Entry       Entry     `gorm:"foreignKey:EntryID;constraint:OnDelete:CASCADE"`
}

func (*Topic) TableName() string {
    return "topics"
}

func (t *Topic) BeforeUpdate(tx *gorm.DB) (err error) {
    tx.Statement.SetColumn("UpdatedAt", time.Now())
    return
}