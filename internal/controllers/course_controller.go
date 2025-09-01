package controllers

import (
	"lms-backend/config"
	"lms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	var courses []models.Course
	q := config.DB.Preload("Modules.Lessons").Preload("Tags")
	// basic filters
	if qStr := c.Query("q"); qStr != "" {
		q = q.Where("title ILIKE ?", "%"+qStr+"%")
	}
	q.Find(&courses)
	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := config.DB.Preload("Modules.Lessons").Preload("Tags").First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": course})
}

func CreateCourse(c *gin.Context) {
	var input struct {
		Title        string `json:"title" binding:"required"`
		Description  string `json:"description"`
		InstructorID uint   `json:"instructor_id" binding:"required"`
		CategoryID   *uint  `json:"category_id"`
		TagIDs       []uint `json:"tag_ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := models.Course{
		Title:        input.Title,
		Description:  input.Description,
		InstructorID: input.InstructorID,
		CategoryID:   input.CategoryID,
	}
	if err := config.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// attach tags if provided
	if len(input.TagIDs) > 0 {
		var tags []models.Tag
		config.DB.Find(&tags, input.TagIDs)
		config.DB.Model(&course).Association("Tags").Replace(&tags)
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := config.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
	}
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&course).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": course})
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Course{}, id)
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
