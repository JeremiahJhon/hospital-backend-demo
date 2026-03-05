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
)

type MockAuthService struct {
	CreateStaffFn func(ctx context.Context, username, password string, hospitalId uuid.UUID) error
	LoginFn       func(ctx context.Context, username, password string, hospitalId uuid.UUID) (string, error)
}

func (m *MockAuthService) CreateStaff(ctx context.Context, username, password string, hospitalId uuid.UUID) error {
	return m.CreateStaffFn(ctx, username, password, hospitalId)
}

func (m *MockAuthService) Login(ctx context.Context, username, password string, hospitalId uuid.UUID) (string, error) {
	return m.LoginFn(ctx, username, password, hospitalId)
}

func setupRouterAuth(h *handler.AuthHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/staff", h.CreateStaff)
	r.POST("/login", h.Login)
	return r
}

func TestCreateStaff_Success(t *testing.T) {
	mock := &MockAuthService{
		CreateStaffFn: func(ctx context.Context, username, password string, hospitalId uuid.UUID) error {
			return nil
		},
	}

	h := handler.NewAuthHandler(mock)
	router := setupRouterAuth(h)

	body, _ := json.Marshal(dto.StaffCreateRequest{
		Username:   "john",
		Password:   "pass",
		HospitalID: uuid.MustParse("6c30824f-259c-4ee8-88c4-9e51258f9b22"),
	})

	req := httptest.NewRequest(http.MethodPost, "/staff", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.Code)
	}
}

func TestCreateStaff_BindError(t *testing.T) {
	mock := &MockAuthService{}
	h := handler.NewAuthHandler(mock)
	router := setupRouterAuth(h)

	req := httptest.NewRequest(http.MethodPost, "/staff", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.Code)
	}
}

func TestLogin_Success(t *testing.T) {
	mock := &MockAuthService{
		LoginFn: func(ctx context.Context, username, password string, hospitalId uuid.UUID) (string, error) {
			return "mock-token", nil
		},
	}

	h := handler.NewAuthHandler(mock)
	router := setupRouterAuth(h)

	body, _ := json.Marshal(dto.StaffLoginRequest{
		Username: "john",
		Password: "pass",
	})

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	mock := &MockAuthService{
		LoginFn: func(ctx context.Context, username, password string, hospitalId uuid.UUID) (string, error) {
			return "", errors.New("invalid credentials")
		},
	}

	h := handler.NewAuthHandler(mock)
	router := setupRouterAuth(h)

	body, _ := json.Marshal(dto.StaffLoginRequest{
		Username: "john",
		Password: "wrong",
	})

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", resp.Code)
	}
}
