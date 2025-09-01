package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	uploadPath := os.Getenv("UPLOAD_PATH")
	os.MkdirAll(uploadPath, os.ModePerm)

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	savePath := filepath.Join(uploadPath, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		return "", err
	}

	return savePath, nil
}
