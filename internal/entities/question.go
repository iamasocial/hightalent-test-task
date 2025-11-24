package entities

import "time"

type Question struct {
	Id        uint
	Text      string
	CreatedAt time.Time
	Answers   []*Answer
}
