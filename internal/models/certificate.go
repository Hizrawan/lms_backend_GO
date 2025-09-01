package models

import "time"

type Certificate struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	CourseID       uint      `json:"course_id"`
	UserID         uint      `json:"user_id"`
	IssuedAt       time.Time `json:"issued_at"`
	CertificateURL *string   `json:"certificate_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
