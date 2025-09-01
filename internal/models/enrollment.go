package models

import "time"

type Enrollment struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	CourseID   uint      `json:"course_id"`
	EnrolledAt time.Time `json:"enrolled_at"`
}
