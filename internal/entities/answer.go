package entities

import "time"

// Answer represents a domain-level answer entity
type Answer struct {
	ID         uint
	QuestionID uint
	UserID     string
	Text       string
	CreatedAt  time.Time
}
