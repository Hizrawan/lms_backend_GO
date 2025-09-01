package models

import "time"

type Assignment struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	CourseID    uint       `json:"course_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
