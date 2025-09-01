package models

import "time"

type DiscussionReport struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `json:"user_id"`
	ReportableID   uint      `json:"reportable_id"`
	ReportableType string    `json:"reportable_type"`
	Reason         string    `json:"reason"`
	CreatedAt      time.Time `json:"created_at"`
}
