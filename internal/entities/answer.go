package entities

import "time"

type Answer struct {
	Id         uint
	QuestionId uint
	UserId     string
	Text       string
	CreatedAt  time.Time
}
