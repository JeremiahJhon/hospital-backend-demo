package repository

import (
	"context"
	"time"

	"hospital-backend-demo/internal/dto"
	"hospital-backend-demo/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PatientRepository struct {
	DB *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{DB: db}
}

func (r *PatientRepository) Search(
	ctx context.Context,
	hospitalID uuid.UUID,
	input dto.PatientSearchRequest,
) ([]models.Patient, error) {

	query := r.DB.WithContext(ctx).
		Where("hospital_id = ?", hospitalID)

	if input.NationalID != "" {
		query = query.Where("national_id = ?", input.NationalID)
	}

	if input.PassportID != "" {
		query = query.Where("passport_id = ?", input.PassportID)
	}

	if input.FirstName != "" {
		query = query.Where("first_name_en ILIKE ?", "%"+input.FirstName+"%")
	}

	if input.LastName != "" {
		query = query.Where("last_name_en ILIKE ?", "%"+input.LastName+"%")
	}

	if input.PhoneNumber != "" {
		query = query.Where("phone_number = ?", input.PhoneNumber)
	}

	if input.Email != "" {
		query = query.Where("email = ?", input.Email)
	}

	if input.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", input.DateOfBirth)
		if err == nil {
			query = query.Where("date_of_birth = ?", dob)
		}
	}

	var patients []models.Patient
	err := query.Find(&patients).Error

	return patients, err
}

func (r *PatientRepository) Upsert(
	ctx context.Context,
	hospitalID uuid.UUID,
	patient models.Patient,
) error {

	var existing models.Patient

	query := r.DB.WithContext(ctx).
		Where("hospital_id = ?", hospitalID)

	if patient.NationalID != "" {
		query = query.Where("national_id = ?", patient.NationalID)
	} else if patient.PassportID != "" {
		query = query.Where("passport_id = ?", patient.PassportID)
	}

	err := query.First(&existing).Error

	if err == nil {
		patient.ID = existing.ID
		return r.DB.WithContext(ctx).Save(&patient).Error
	}

	return r.DB.WithContext(ctx).Create(&patient).Error
}
