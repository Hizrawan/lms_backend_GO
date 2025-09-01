package controllers

import (
	"lms-backend/internal/config"
	"lms-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetThreads(c *gin.Context) {
	var threads []models.DiscussionThread
	config.DB.Preload("Replies").Find(&threads)
	c.JSON(http.StatusOK, gin.H{"data": threads})
}

func GetThreadByID(c *gin.Context) {
	id := c.Param("id")
	var thread models.DiscussionThread
	if err := config.DB.Preload("Replies").First(&thread, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "thread not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": thread})
}

func CreateThread(c *gin.Context) {
	var input struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content"`
		UserID   uint   `json:"user_id" binding:"required"`
		CourseID *uint  `json:"course_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	thread := models.DiscussionThread{
		Title:    input.Title,
		Content:  input.Content,
		UserID:   input.UserID,
		CourseID: input.CourseID,
	}
	config.DB.Create(&thread)
	c.JSON(http.StatusOK, gin.H{"data": thread})
}

func CreateReply(c *gin.Context) {
	threadID := c.Param("id")
	var input struct {
		ParentReplyID *uint  `json:"parent_reply_id"`
		Content       string `json:"content" binding:"required"`
		UserID        uint   `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reply := models.DiscussionReply{
		ThreadID:      parseUint(threadID),
		ParentReplyID: input.ParentReplyID,
		UserID:        input.UserID,
		Content:       input.Content,
	}
	config.DB.Create(&reply)
	c.JSON(http.StatusOK, gin.H{"data": reply})
}

func GetReplies(c *gin.Context) {
	threadID := c.Param("id")
	var replies []models.DiscussionReply
	config.DB.Where("thread_id = ?", threadID).Find(&replies)
	c.JSON(http.StatusOK, gin.H{"data": replies})
}
