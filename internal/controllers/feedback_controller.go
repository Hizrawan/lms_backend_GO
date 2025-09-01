package controllers

import (
	"lms-backend/config"
	"lms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourseFeedback(c *gin.Context) {
	courseID := c.Param("course_id")
	var fb []models.Feedback
	config.DB.Where("course_id = ?", courseID).Find(&fb)
	c.JSON(http.StatusOK, gin.H{"data": fb})
}

func PostFeedback(c *gin.Context) {
	courseID := c.Param("course_id")
	var input struct {
		UserID  uint    `json:"user_id" binding:"required"`
		Rating  int     `json:"rating" binding:"required,min=1,max=5"`
		Comment *string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	feedback := models.Feedback{
		UserID:   input.UserID,
		CourseID: parseUint(courseID),
		Rating:   input.Rating,
		Comment:  input.Comment,
	}
	config.DB.Create(&feedback)
	c.JSON(http.StatusOK, gin.H{"data": feedback})
}
