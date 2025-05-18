package domain

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Reminder struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TopicID     uuid.UUID `gorm:"type:uuid;not null"`
	RemindAt    time.Time `gorm:"type:timestamptz;not null"`
	IsCompleted bool      `gorm:"default:false;not null"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null;default:now()"`

	Topic Topic `gorm:"foreignKey:TopicID;constraint:OnDelete:CASCADE"`
}

func (*Reminder) TableName() string {
	return "reminders"
}
