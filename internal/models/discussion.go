package models

import "time"

type DiscussionThread struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	UserID    uint              `json:"user_id"`
	CourseID  *uint             `json:"course_id"`
	Views     int               `json:"views"`
	Status    string            `json:"status"`
	Replies   []DiscussionReply `gorm:"foreignKey:ThreadID" json:"replies"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type DiscussionReply struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	ThreadID      uint      `json:"thread_id"`
	ParentReplyID *uint     `json:"parent_reply_id"`
	UserID        uint      `json:"user_id"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
