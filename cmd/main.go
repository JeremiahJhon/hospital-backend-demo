package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"hospital-backend-demo/internal/client"
	"hospital-backend-demo/internal/config"
	"hospital-backend-demo/internal/models"
	"hospital-backend-demo/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	// Load app config
	appConfig := config.LoadConfig()

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err := db.AutoMigrate(
		&models.Hospital{},
		&models.Staff{},
		&models.Patient{},
	); err != nil {
		log.Fatal("Migration failed:", err)
	}

	if err := config.SeedDatabase(db); err != nil {
		log.Fatal("Seeding failed:", err)
	}

	// Create HIS client
	hisClient := client.NewHISClient(appConfig.HISBaseURL)

	router := gin.Default()

	// Pass both db and hisClient
	routes.RegisterRoutes(router, db, hisClient)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
