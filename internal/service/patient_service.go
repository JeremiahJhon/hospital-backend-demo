package service

import (
	"context"
	"time"

	"hospital-backend-demo/internal/client"
	"hospital-backend-demo/internal/dto"
	"hospital-backend-demo/internal/models"
	"hospital-backend-demo/internal/repository"

	"github.com/google/uuid"
)

type PatientService struct {
	Repo      *repository.PatientRepository
	StaffRepo *repository.StaffRepository
	HISClient *client.HISClient
}

func NewPatientService(
	repo *repository.PatientRepository,
	staffrepo *repository.StaffRepository,
	his *client.HISClient,
) *PatientService {
	return &PatientService{
		Repo:      repo,
		StaffRepo: staffrepo,
		HISClient: his,
	}
}

func (s *PatientService) Search(
	ctx context.Context,
	hospitalID uuid.UUID,
	input dto.PatientSearchRequest,
) ([]models.Patient, error) {

	// Sync from HIS if National ID provided
	if input.NationalID != "" {
		if err := s.syncFromHIS(ctx, hospitalID, input.NationalID); err != nil {
			return nil, err
		}
	}

	// Sync from HIS if Passport ID provided
	if input.PassportID != "" {
		if err := s.syncFromHIS(ctx, hospitalID, input.PassportID); err != nil {
			return nil, err
		}
	}

	// Search from local DB
	return s.Repo.Search(ctx, hospitalID, input)
}

func (s *PatientService) syncFromHIS(
	ctx context.Context,
	hospitalID uuid.UUID,
	identifier string,
) error {

	if s.HISClient == nil {
		return nil // HIS integration optional
	}

	hospital, err := s.StaffRepo.FindHospitalById(ctx, hospitalID)

	his, err := s.HISClient.SearchPatient(identifier, hospital.ApiBaseURL)
	if err != nil || his == nil {
		return err
	}

	dob, _ := time.Parse("2006-01-02", his.DateOfBirth)

	patient := models.Patient{
		ID:           uuid.New(),
		HospitalID:   hospitalID,
		NationalID:   his.NationalID,
		PassportID:   his.PassportID,
		FirstNameTH:  his.FirstNameTH,
		MiddleNameTH: his.MiddleNameTH,
		LastNameTH:   his.LastNameTH,
		FirstNameEN:  his.FirstNameEN,
		MiddleNameEN: his.MiddleNameEN,
		LastNameEN:   his.LastNameEN,
		DateOfBirth:  dob,
		PatientHN:    his.PatientHN,
		PhoneNumber:  his.PhoneNumber,
		Email:        his.Email,
		Gender:       his.Gender,
	}

	return s.Repo.Upsert(ctx, hospitalID, patient)
}
