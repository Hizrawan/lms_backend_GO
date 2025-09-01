package controllers

import (
	"lms-backend/config"
	"lms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateQuiz(c *gin.Context) {
	var q models.Quiz
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&q).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": q})
}

func SubmitAnswer(c *gin.Context) {
	var a models.Answer
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&a).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func GetQuizResults(c *gin.Context) {
	quizID := c.Param("id")
	var answers []models.Answer
	config.DB.Joins("JOIN questions q ON q.id = answers.question_id").Where("q.quiz_id = ?", quizID).Find(&answers)
	c.JSON(http.StatusOK, gin.H{"data": answers})
}
func GetQuizzes(c *gin.Context) {
	courseID := c.Query("course_id") // optional filter by course_id

	var quizzes []models.Quiz

	if courseID != "" {
		if err := config.DB.Where("course_id = ?", courseID).Find(&quizzes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := config.DB.Find(&quizzes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": quizzes})
}
