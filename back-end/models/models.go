package models

import (
	"time"
)

// Model for use with GORM
type SavedURL struct {
	ID        int32  `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	URL       string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
