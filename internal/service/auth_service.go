package service

import (
	"context"
	"errors"

	"hospital-backend-demo/internal/models"
	"hospital-backend-demo/internal/repository"
	"hospital-backend-demo/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.StaffRepository
}

func NewAuthService(repo *repository.StaffRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) CreateStaff(ctx context.Context, username, password string, hospitalId uuid.UUID) error {
	// Check if hospital exists
	hospital, err := s.Repo.FindHospitalById(ctx, hospitalId)
	if err != nil {
		return errors.New("hospital not found")
	}

	// Check duplicate username
	existing, _ := s.Repo.FindByUsername(ctx, username, hospitalId)
	if existing != nil {
		return errors.New("username already exists")
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	staff := &models.Staff{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: string(hash),
		HospitalID:   hospital.ID,
	}

	return s.Repo.Create(ctx, staff)
}

func (s *AuthService) Login(ctx context.Context, username, password string, hospitalId uuid.UUID) (string, error) {
	staff, err := s.Repo.FindByUsername(ctx, username, hospitalId)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(staff.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(staff.ID, staff.HospitalID)
}
