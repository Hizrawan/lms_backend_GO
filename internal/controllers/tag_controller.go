package controllers

import (
	"lms-backend/config"
	"lms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	var tags []models.Tag
	config.DB.Find(&tags)
	c.JSON(http.StatusOK, gin.H{"data": tags})
}

func CreateTag(c *gin.Context) {
	var t models.Tag
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
