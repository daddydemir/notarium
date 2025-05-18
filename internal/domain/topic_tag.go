package domain

import (
    "time"

    "github.com/google/uuid"
    _ "gorm.io/gorm"
)

type TopicTag struct {
    TopicID    uuid.UUID `gorm:"type:uuid;primaryKey"`
    TagID      uuid.UUID `gorm:"type:uuid;primaryKey"`
    CreatedAt  time.Time `gorm:"type:timestamptz;not null;default:now()"`

    Topic      Topic     `gorm:"foreignKey:TopicID;constraint:OnDelete:CASCADE"`
    Tag        Tag       `gorm:"foreignKey:TagID;constraint:OnDelete:CASCADE"`
}

func (*TopicTag) TableName() string {
    return "topic_tags"
}