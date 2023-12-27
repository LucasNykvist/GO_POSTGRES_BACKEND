package controllers

import (
	"gopg_api/src/database"
	"gopg_api/src/models"

	"github.com/gin-gonic/gin"
)

func CreateSnus(c *gin.Context) {
	var input models.Snus
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	snus := models.Snus{
		Brand:        input.Brand,
		Flavour: input.Flavour,
		Price:       input.Price,
	}

	err := database.Database.Create(&snus).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create snus"})
		return
	}

	c.JSON(200, gin.H{"message": "Snus created successfully"})
}

func GetSnuses(c *gin.Context) {
	var snuses []models.Snus

	err := database.Database.Find(&snuses).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch snuses"})
		return
	}

	c.JSON(200, snuses)
}