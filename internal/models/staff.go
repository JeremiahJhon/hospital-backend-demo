package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Staff struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username     string    `gorm:"size:100;unique;not null"`
	PasswordHash string    `gorm:"type:text;not null"`
	HospitalID   uuid.UUID `gorm:"type:uuid;not null"`
	Hospital     Hospital  `gorm:"foreignKey:HospitalID"`

	CreatedAt time.Time
}

func (s *Staff) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
