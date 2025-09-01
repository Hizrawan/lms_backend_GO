package controllers

import (
	"lms-backend/config"
	"lms-backend/models"
	"lms-backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UploadAttachment(c *gin.Context) {
	attachableType := c.PostForm("attachable_type")
	attachableIDStr := c.PostForm("attachable_id")
	userIDStr := c.PostForm("user_id")

	if attachableType == "" || attachableIDStr == "" || userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "attachable_type, attachable_id, user_id required"})
		return
	}

	attachableID := parseUint(attachableIDStr)
	userID := parseUint(userIDStr)

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["attachments"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no files"})
		return
	}

	var created []models.Attachment
	for _, file := range files {
		// gunakan utils.UploadFile untuk simpan file
		savePath, err := utils.UploadFile(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		att := models.Attachment{
			FileName:       file.Filename,
			FilePath:       savePath, // simpan path lokal
			FileType:       file.Header.Get("Content-Type"),
			FileSize:       file.Size,
			UserID:         &userID,
			AttachableID:   attachableID,
			AttachableType: attachableType,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := config.DB.Create(&att).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		created = append(created, att)
	}

	c.JSON(http.StatusOK, gin.H{"uploaded": created})
}

func saveAttachment(db *gorm.DB, a *models.Attachment) error {
	return db.Create(a).Error
}
