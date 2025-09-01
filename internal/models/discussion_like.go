package models

import "time"

type DiscussionLike struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"user_id"`
	LikeableID   uint      `json:"likeable_id"`
	LikeableType string    `json:"likeable_type"`
	CreatedAt    time.Time `json:"created_at"`
}
