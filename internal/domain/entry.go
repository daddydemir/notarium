package domain

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Entry struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Date      time.Time `gorm:"type:date;uniqueIndex;not null"`
    Title     string    `gorm:"type:varchar(255)"`
    CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
    UpdatedAt time.Time `gorm:"type:timestamptz"`
}

func (*Entry) TableName() string {
    return "entries"
}

func (e *Entry) BeforeUpdate(tx *gorm.DB) (err error) {
    tx.Statement.SetColumn("UpdatedAt", time.Now())
    return
}
