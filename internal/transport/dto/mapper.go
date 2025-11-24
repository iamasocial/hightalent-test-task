package dto

import (
	"github.com/iamasocial/hightalent-test-task/internal/entities"
)

// QuestionToResponse converts a Question entity to a QuestionResponse DTO
func QuestionToResponse(q *entities.Question) QuestionResponse {
	answers := make([]AnswerResponse, 0, len(q.Answers))
	for _, a := range q.Answers {
		answers = append(answers, AnswerToResponse(a))
	}

	return QuestionResponse{
		ID:       q.ID,
		Text:     q.Text,
		CreateAt: q.CreatedAt,
		Answers:  answers,
	}
}

// AnswerToResponse converts an Answer entity to an AnswerResponse DTO
func AnswerToResponse(a *entities.Answer) AnswerResponse {
	return AnswerResponse{
		ID:         a.ID,
		QuestionID: a.QuestionID,
		UserID:     a.UserID,
		Text:       a.Text,
		CreatedAt:  a.CreatedAt,
	}
}
