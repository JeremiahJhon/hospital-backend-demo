package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hospital struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name       string    `gorm:"size:150;not null"`
	ApiBaseURL string    `gorm:"type:text;not null"`
	CreatedAt  time.Time
}

func (h *Hospital) BeforeCreate(tx *gorm.DB) error {
	h.ID = uuid.New()
	return nil
}