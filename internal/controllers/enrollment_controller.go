package controllers

import (
	"fmt"
	"lms-backend/config"
	"lms-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func EnrollCourse(c *gin.Context) {
	courseID := c.Param("id")
	var input struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enrollment := models.Enrollment{
		UserID:     input.UserID,
		CourseID:   parseUint(courseID),
		EnrolledAt: time.Now(),
	}
	if err := config.DB.Where("user_id = ? AND course_id = ?", enrollment.UserID, enrollment.CourseID).First(&models.Enrollment{}).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "already enrolled"})
		return
	}
	config.DB.Create(&enrollment)
	c.JSON(http.StatusOK, gin.H{"data": enrollment})
}

func GetUserEnrollments(c *gin.Context) {
	userID := c.Param("id")
	var enrollments []models.Enrollment
	config.DB.Where("user_id = ?", userID).Find(&enrollments)
	c.JSON(http.StatusOK, gin.H{"data": enrollments})
}

// helper
func parseUint(s string) uint {
	var v uint = 0
	fmt.Sscan(s, &v)
	return v
}
