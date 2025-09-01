package controllers

import (
	"net/http"

	"lms-backend/config"
	"lms-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateLesson(c *gin.Context) {
	moduleID := c.Param("id")
	var mod models.Module
	if err := config.DB.First(&mod, moduleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "module not found"})
		return
	}
	var input struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content"`
		VideoURL string `json:"video_url"`
		Position int    `json:"position"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lesson := models.Lesson{
		ModuleID: mod.ID, Title: input.Title, Content: input.Content, VideoURL: input.VideoURL, Position: input.Position,
	}
	config.DB.Create(&lesson)
	c.JSON(http.StatusOK, gin.H{"data": lesson})
}

// GetLesson - ambil lesson berdasarkan ID
func GetLesson(c *gin.Context) {
	id := c.Param("id")
	var lesson models.Lesson

	if err := config.DB.First(&lesson, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

// UpdateLesson - update lesson berdasarkan ID
func UpdateLesson(c *gin.Context) {
	id := c.Param("id")
	var lesson models.Lesson

	if err := config.DB.First(&lesson, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	var input models.Lesson
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&lesson).Updates(input)
	c.JSON(http.StatusOK, lesson)
}

// DeleteLesson - hapus lesson berdasarkan ID
func DeleteLesson(c *gin.Context) {
	id := c.Param("id")
	var lesson models.Lesson

	if err := config.DB.First(&lesson, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	config.DB.Delete(&lesson)
	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted"})
}

// ListLessons - ambil semua lesson
func ListLessons(c *gin.Context) {
	var lessons []models.Lesson
	config.DB.Find(&lessons)
	c.JSON(http.StatusOK, lessons)
}

func GetLessons(c *gin.Context) {
	moduleID := c.Param("id")
	var lessons []models.Lesson

	if err := config.DB.Where("module_id = ?", moduleID).Order("position asc").Find(&lessons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lessons)
}
