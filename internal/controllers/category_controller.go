package controllers

import (
	"lms-backend/config"
	"lms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var cats []models.Category
	config.DB.Find(&cats)
	c.JSON(http.StatusOK, gin.H{"data": cats})
}

func CreateCategory(c *gin.Context) {
	var t models.Category
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&t).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": t})
}
