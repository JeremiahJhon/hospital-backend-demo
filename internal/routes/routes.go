package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"hospital-backend-demo/internal/client"
	"hospital-backend-demo/internal/handler"
	"hospital-backend-demo/internal/middleware"
	"hospital-backend-demo/internal/repository"
	"hospital-backend-demo/internal/service"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, hisClient *client.HISClient) {

	// ===== Staff (Auth) =====
	authRepo := repository.NewStaffRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	router.POST("/staff/create", authHandler.CreateStaff)
	router.POST("/staff/login", authHandler.Login)

	// ===== Patient =====
	patientRepo := repository.NewPatientRepository(db)
	staffRepo := repository.NewStaffRepository(db)
	patientService := service.NewPatientService(patientRepo, staffRepo, nil) // inject HIS later
	patientHandler := handler.NewPatientHandler(patientService)

	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.POST("/patient/search", patientHandler.Search)
	}
}
