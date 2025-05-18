package domain

import (
	"time"
	"fmt"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type File struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	NoteID    uuid.UUID `gorm:"type:uuid;not null"`
	URL       string    `gorm:"type:text;not null"`
	FileName  string    `gorm:"type:varchar(255);not null"`
	MimeType  string    `gorm:"type:varchar(100);not null"`
	Size      int64     `gorm:"type:bigint"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`

	// İlişki tanımı
	Note Note `gorm:"foreignKey:NoteID;constraint:OnDelete:CASCADE"`
}

func (*File) TableName() string {
	return "files"
}

func (f *File) HumanReadableSize() string {
	const unit = 1024
	if f.Size < unit {
		return fmt.Sprintf("%d B", f.Size)
	}
	div, exp := int64(unit), 0
	for n := f.Size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(f.Size)/float64(div), "KMGTPE"[exp])
}
