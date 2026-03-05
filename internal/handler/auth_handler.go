package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"hospital-backend-demo/internal/dto"
	"hospital-backend-demo/internal/service"
)

type AuthHandler struct {
	Service service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) CreateStaff(c *gin.Context) {
	var req dto.StaffCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.CreateStaff(c, req.Username, req.Password, req.HospitalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "staff created"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.StaffLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Service.Login(c, req.Username, req.Password, req.HospitalID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
