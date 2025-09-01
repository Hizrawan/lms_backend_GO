package models

import "time"

type Question struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	QuizID       uint      `json:"quiz_id"`
	QuestionText string    `json:"question_text"`
	QuestionType string    `json:"question_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
