package repository

import (
	"github.com/iamasocial/hightalent-test-task/internal/db"
	"github.com/iamasocial/hightalent-test-task/internal/entities"
)

func toEntityQuestion(q db.Question) *entities.Question {
	answers := make([]*entities.Answer, 0, len(q.Answers))
	for _, a := range q.Answers {
		answers = append(answers, toEntityAnswer(a))
	}

	return &entities.Question{
		ID:        q.ID,
		Text:      q.Text,
		CreatedAt: q.CreatedAt,
		Answers:   answers,
	}
}

func toModelQuestion(q *entities.Question) db.Question {
	return db.Question{
		ID:        q.ID,
		Text:      q.Text,
		CreatedAt: q.CreatedAt,
	}
}

func toEntityAnswer(a db.Answer) *entities.Answer {
	return &entities.Answer{
		ID:         a.ID,
		QuestionID: a.QuestionID,
		UserID:     a.UserID,
		Text:       a.Text,
		CreatedAt:  a.CreatedAt,
	}
}

func toModelAnswer(a *entities.Answer) db.Answer {
	return db.Answer{
		ID:         a.ID,
		QuestionID: a.QuestionID,
		UserID:     a.UserID,
		Text:       a.Text,
		CreatedAt:  a.CreatedAt,
	}
}
