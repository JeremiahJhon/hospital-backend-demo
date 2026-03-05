package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	HospitalID uuid.UUID `gorm:"type:uuid;index;not null"`

	NationalID string `gorm:"size:20;index"`
	PassportID string `gorm:"size:20;index"`

	FirstNameTH  string
	MiddleNameTH string
	LastNameTH   string

	FirstNameEN  string
	MiddleNameEN string
	LastNameEN   string

	DateOfBirth time.Time
	PatientHN   string

	PhoneNumber string
	Email       string

	Gender string `gorm:"size:1"`

	CreatedAt time.Time

	Hospital Hospital `gorm:"foreignKey:HospitalID"`
}

func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
