package domain

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Note struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    TopicID   uuid.UUID `gorm:"type:uuid;not null"`
    Content   string    `gorm:"type:text;not null"`
    Version   int       `gorm:"not null;default:1"`
    CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
    UpdatedAt time.Time `gorm:"type:timestamptz"`

    Topic     Topic     `gorm:"foreignKey:TopicID;constraint:OnDelete:CASCADE"`
}

func (*Note) TableName() string {
    return "notes"
}

func (n *Note) BeforeUpdate(tx *gorm.DB) (err error) {
    n.UpdatedAt = time.Now()
    n.Version++
    return nil
}