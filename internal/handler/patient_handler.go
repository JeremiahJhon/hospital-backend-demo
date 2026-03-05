package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"hospital-backend-demo/internal/dto"
	"hospital-backend-demo/internal/service"
)

type PatientHandler struct {
	Service service.PatientServiceInterface
}

func NewPatientHandler(service service.PatientServiceInterface) *PatientHandler {
	return &PatientHandler{Service: service}
}

func (h *PatientHandler) Search(c *gin.Context) {

	var req dto.PatientSearchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospitalIDRaw, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "hospital context missing"})
		return
	}

	hospitalID := hospitalIDRaw.(uuid.UUID)

	patients, err := h.Service.Search(c, hospitalID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}
