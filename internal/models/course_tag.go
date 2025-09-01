package models

type CourseTag struct {
	CourseID uint `gorm:"primaryKey"`
	TagID    uint `gorm:"primaryKey"`
}
