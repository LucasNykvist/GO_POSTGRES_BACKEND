package main

import (
	"gopg_api/src/controllers"
	"gopg_api/src/database"
	"gopg_api/src/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplicaton()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.Snus{})
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplicaton() {
	router := gin.Default()
	router.POST("/snuses", controllers.CreateSnus)
	router.GET("/snuses", controllers.GetSnuses)
	router.Run(":8080")
}
