package repository

import (
	"context"

	"hospital-backend-demo/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StaffRepository struct {
	DB *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{DB: db}
}

func (r *StaffRepository) Create(ctx context.Context, staff *models.Staff) error {
	return r.DB.WithContext(ctx).Create(staff).Error
}

func (r *StaffRepository) FindByUsername(ctx context.Context, username string, hospitalId uuid.UUID) (*models.Staff, error) {
	var staff models.Staff
	err := r.DB.WithContext(ctx).
		Where(&models.Staff{Username: username, HospitalID: hospitalId}).
		First(&staff).Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (r *StaffRepository) FindHospitalById(ctx context.Context, id uuid.UUID) (*models.Hospital, error) {
	var hospital models.Hospital
	err := r.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&hospital).Error

	if err != nil {
		return nil, err
	}

	return &hospital, nil
}

func (r *StaffRepository) FindHospitalByName(ctx context.Context, name string) (*models.Hospital, error) {
	var hospital models.Hospital
	err := r.DB.WithContext(ctx).
		Where("name = ?", name).
		First(&hospital).Error

	if err != nil {
		return nil, err
	}

	return &hospital, nil
}
