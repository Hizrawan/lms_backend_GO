package controllers

import (
	"lms-backend/config"
	"lms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateModule(c *gin.Context) {
	courseID := c.Param("id")
	var course models.Course
	if err := config.DB.First(&course, courseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
	}
	var input struct {
		Title    string `json:"title" binding:"required"`
		Position int    `json:"position"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mod := models.Module{CourseID: course.ID, Title: input.Title, Position: input.Position}
	config.DB.Create(&mod)
	c.JSON(http.StatusOK, gin.H{"data": mod})
}
