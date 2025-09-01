package models

import "time"

type AssignmentSubmission struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	AssignmentID uint      `json:"assignment_id"`
	UserID       uint      `json:"user_id"`
	SubmittedAt  time.Time `json:"submitted_at"`
	Grade        *float64  `json:"grade"`
	Feedback     *string   `json:"feedback"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
