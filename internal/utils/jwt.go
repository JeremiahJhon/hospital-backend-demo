package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("super-secret-key") // move to env later

type Claims struct {
	StaffID    uuid.UUID `json:"staff_id"`
	HospitalID uuid.UUID `json:"hospital_id"`
	jwt.RegisteredClaims
}

func GenerateToken(staffID, hospitalID uuid.UUID) (string, error) {
	claims := Claims{
		StaffID:    staffID,
		HospitalID: hospitalID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}