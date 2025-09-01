package models

import "time"

type Lesson struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ModuleID  uint      `json:"module_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	VideoURL  string    `json:"video_url"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
