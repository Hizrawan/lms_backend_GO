package models

import "time"

type Module struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CourseID  uint      `json:"course_id"`
	Title     string    `json:"title"`
	Position  int       `json:"position"`
	Lessons   []Lesson  `json:"lessons"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
