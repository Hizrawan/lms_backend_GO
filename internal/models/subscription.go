package models

import "time"

type DiscussionSubscription struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `json:"user_id"`
	ThreadID      uint      `json:"thread_id"`
	NotifyOnReply bool      `json:"notify_on_reply"`
	CreatedAt     time.Time `json:"created_at"`
}
