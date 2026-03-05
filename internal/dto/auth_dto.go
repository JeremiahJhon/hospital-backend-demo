package dto

import "github.com/google/uuid"

type StaffCreateRequest struct {
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	HospitalID uuid.UUID `json:"hospital" binding:"required"`
}

type StaffLoginRequest struct {
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	HospitalID uuid.UUID `json:"hospital" binding:"required"`
}
