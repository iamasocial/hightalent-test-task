package db

import (
	"time"
)

type Question struct {
	Id        uint   `gorm:"primaryKey"`
	Text      string `gorm:"not null"`
	CreatedAt time.Time
	Answers   []Answer `gorm:"constraint:OnDelete:CASCADE"`
}

type Answer struct {
	Id         uint   `gorm:"primaryKey"`
	QuestionId uint   `gorm:"not null;index"`
	UserId     string `gorm:"not null"`
	Text       string `gorm:"not null"`
	CreatedAt  time.Time
}
