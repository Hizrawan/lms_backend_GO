package controllers

import (
	"lms-backend/internal/config"
	"lms-backend/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func IssueCertificate(c *gin.Context) {
	var input struct {
		CourseID uint   `json:"course_id" binding:"required"`
		UserID   uint   `json:"user_id" binding:"required"`
		URL      string `json:"url"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cert := models.Certificate{
		CourseID: input.CourseID,
		UserID:   input.UserID,
		IssuedAt: time.Now(),
	}
	if input.URL != "" {
		cert.CertificateURL = &input.URL
	}
	config.DB.Create(&cert)
	c.JSON(http.StatusOK, gin.H{"data": cert})
}
func GetCertificates(c *gin.Context) {
	userID := c.Query("user_id")
	courseID := c.Query("course_id")

	var certs []models.Certificate
	db := config.DB

	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}

	if courseID != "" {
		db = db.Where("course_id = ?", courseID)
	}

	if err := db.Find(&certs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": certs})
}
