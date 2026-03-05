package config

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"hospital-backend-demo/internal/models"
)

func SeedDatabase(db *gorm.DB) error {
	var count int64

	// Prevent duplicate seeding
	db.Model(&models.Hospital{}).Count(&count)
	if count > 0 {
		log.Println("Seed already executed. Skipping...")
		return nil
	}

	// Create Hospitals
	hospital1 := models.Hospital{
		ID:   uuid.New(),
		Name: "ABC Hospital",
	}

	hospital2 := models.Hospital{
		ID:   uuid.New(),
		Name: "XYZ Hospital",
	}

	if err := db.Create(&hospital1).Error; err != nil {
		return err
	}

	if err := db.Create(&hospital2).Error; err != nil {
		return err
	}

	now := time.Now()

	// Create Patients
	patients := []models.Patient{
		{
			ID:          uuid.New(),
			HospitalID:  hospital1.ID,
			NationalID:  "1103701234567",
			FirstNameTH: "สมชาย",
			LastNameTH:  "ใจดี",
			FirstNameEN: "Somchai",
			LastNameEN:  "Jaidee",
			DateOfBirth: time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			PatientHN:   "HN0001",
			PhoneNumber: "0812345678",
			Email:       "somchai@example.com",
			Gender:      "M",
			CreatedAt:   now,
		},
		{
			ID:          uuid.New(),
			HospitalID:  hospital1.ID,
			PassportID:  "AA1234567",
			FirstNameTH: "สุดารัตน์",
			LastNameTH:  "สุขใจ",
			FirstNameEN: "Sudarat",
			LastNameEN:  "Sukjai",
			DateOfBirth: time.Date(1985, 8, 20, 0, 0, 0, 0, time.UTC),
			PatientHN:   "HN0002",
			PhoneNumber: "0898765432",
			Email:       "sudarat@example.com",
			Gender:      "F",
			CreatedAt:   now,
		},
		{
			ID:          uuid.New(),
			HospitalID:  hospital2.ID,
			NationalID:  "1209907654321",
			FirstNameTH: "อนันต์",
			LastNameTH:  "ทองดี",
			FirstNameEN: "Anan",
			LastNameEN:  "Thongdee",
			DateOfBirth: time.Date(2000, 1, 15, 0, 0, 0, 0, time.UTC),
			PatientHN:   "HN1001",
			PhoneNumber: "0823456789",
			Email:       "anan@example.com",
			Gender:      "M",
			CreatedAt:   now,
		},
	}

	if err := db.Create(&patients).Error; err != nil {
		return err
	}

	log.Println("Database seeded successfully ✅")
	return nil
}
