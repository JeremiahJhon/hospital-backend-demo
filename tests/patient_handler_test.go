package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"hospital-backend-demo/internal/dto"
	"hospital-backend-demo/internal/handler"
	"hospital-backend-demo/internal/models"
)

type MockPatientService struct {
	SearchFn func(ctx context.Context, hospitalID uuid.UUID, req dto.PatientSearchRequest) ([]models.Patient, error)
}

func (m *MockPatientService) Search(
	ctx context.Context,
	hospitalID uuid.UUID,
	req dto.PatientSearchRequest,
) ([]models.Patient, error) {
	return m.SearchFn(ctx, hospitalID, req)
}

func setupRouterPatient(h *handler.PatientHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/patients/search", h.Search)
	return r
}

func TestPatientSearch_Success(t *testing.T) {
	mock := &MockPatientService{
		SearchFn: func(ctx context.Context, hospitalID uuid.UUID, req dto.PatientSearchRequest) ([]models.Patient, error) {
			return []models.Patient{
				{ID: uuid.New()},
			}, nil
		},
	}

	h := handler.NewPatientHandler(mock)
	router := setupRouterPatient(h)

	body, _ := json.Marshal(dto.PatientSearchRequest{})

	req := httptest.NewRequest(http.MethodPost, "/patients/search", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	// inject hospital_id middleware BEFORE request
	router.Use(func(c *gin.Context) {
		c.Set("hospital_id", uuid.New())
	})

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}
}

func TestPatientSearch_MissingHospitalContext(t *testing.T) {
	mock := &MockPatientService{}
	h := handler.NewPatientHandler(mock)
	router := setupRouterPatient(h)

	body, _ := json.Marshal(dto.PatientSearchRequest{})
	req := httptest.NewRequest(http.MethodPost, "/patients/search", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", resp.Code)
	}
}

func TestPatientSearch_ServiceError(t *testing.T) {
	mock := &MockPatientService{
		SearchFn: func(ctx context.Context, hospitalID uuid.UUID, req dto.PatientSearchRequest) ([]models.Patient, error) {
			return nil, errors.New("db error")
		},
	}

	h := handler.NewPatientHandler(mock)
	router := setupRouterPatient(h)

	body, _ := json.Marshal(dto.PatientSearchRequest{})
	req := httptest.NewRequest(http.MethodPost, "/patients/search", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.Use(func(c *gin.Context) {
		c.Set("hospital_id", uuid.New())
	})

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", resp.Code)
	}
}
