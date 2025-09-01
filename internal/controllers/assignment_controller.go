package controllers

import (
	"lms-backend/config"
	"lms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAssignment(c *gin.Context) {
	var input models.Assignment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func SubmitAssignment(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	submission := models.AssignmentSubmission{
		AssignmentID: parseUint(id),
		UserID:       input.UserID,
	}
	if err := config.DB.Create(&submission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": submission})
}

func GetAssignmentSubmissions(c *gin.Context) {
	id := c.Param("id")
	var subs []models.AssignmentSubmission
	config.DB.Where("assignment_id = ?", id).Find(&subs)
	c.JSON(http.StatusOK, gin.H{"data": subs})
}
func GetAssignments(c *gin.Context) {
	courseID := c.Query("course_id") // optional filter by course_id

	var assignments []models.Assignment

	if courseID != "" {
		if err := config.DB.Where("course_id = ?", courseID).Find(&assignments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := config.DB.Find(&assignments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": assignments})
}
