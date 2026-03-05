package service

import (
	"context"

	"hospital-backend-demo/internal/dto"
	"hospital-backend-demo/internal/models"

	"github.com/google/uuid"
)

type PatientServiceInterface interface {
	Search(ctx context.Context, hospitalID uuid.UUID, req dto.PatientSearchRequest) ([]models.Patient, error)
}
