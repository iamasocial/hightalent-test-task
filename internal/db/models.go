package db

import (
	"time"
)

// Question represents a question entity stored in the database
type Question struct {
	ID        uint   `gorm:"primaryKey"`
	Text      string `gorm:"not null"`
	CreatedAt time.Time
	Answers   []Answer `gorm:"constraint:OnDelete:CASCADE"`
}

// Answer represents an answer to a specific question
type Answer struct {
	ID         uint   `gorm:"primaryKey"`
	QuestionID uint   `gorm:"not null;index"`
	UserID     string `gorm:"not null"`
	Text       string `gorm:"not null"`
	CreatedAt  time.Time
}
