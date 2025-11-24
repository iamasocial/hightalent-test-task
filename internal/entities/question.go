package entities

import "time"

// Question represents a domain-level question entity
type Question struct {
	ID        uint
	Text      string
	CreatedAt time.Time
	Answers   []*Answer
}
