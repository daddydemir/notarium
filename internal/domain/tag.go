package domain

import (
    "fmt"
    "strings"
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Tag struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Name      string    `gorm:"type:varchar(50);not null;uniqueIndex"`
    Color     string    `gorm:"type:varchar(7);default:'#64748B'"`
    CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`

    Entries []Entry `gorm:"many2many:entry_tags;"`
}

func (*Tag) TableName() string {
    return "tags"
}

func (t *Tag) BeforeSave(tx *gorm.DB) error {
    if !strings.HasPrefix(t.Color, "#") || len(t.Color) != 7 {
        return fmt.Errorf("invalid color format")
    }
    return nil
}