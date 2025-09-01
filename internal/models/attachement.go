package models

import "time"

type Attachment struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	FileName       string    `json:"file_name"`
	FilePath       string    `json:"file_path"`
	FileType       string    `json:"file_type"`
	FileSize       int64     `json:"file_size"`
	UserID         *uint     `json:"user_id"`
	AttachableID   uint      `json:"attachable_id"`
	AttachableType string    `json:"attachable_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
