package models

import "time"

type Course struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	InstructorID uint      `json:"instructor_id"`
	Instructor   User      `gorm:"foreignKey:InstructorID"`
	CategoryID   *uint     `json:"category_id"`
	Modules      []Module  `json:"modules"`
	Tags         []Tag     `gorm:"many2many:course_tags" json:"tags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
