package service

import (
	"context"

	"github.com/google/uuid"
)

type AuthServiceInterface interface {
	CreateStaff(ctx context.Context, username, password string, hospitalId uuid.UUID) error
	Login(ctx context.Context, username, password string, hospitalId uuid.UUID) (string, error)
}
