package dto

import (
	"github.com/iamasocial/hightalent-test-task/internal/entities"
)

func QuestionToResponse(q *entities.Question) QuestionResponse {
	answers := make([]AnswerResponse, 0, len(q.Answers))
	for _, a := range q.Answers {
		answers = append(answers, AnswerToResponse(a))
	}

	return QuestionResponse{
		Id:      q.Id,
		Text:    q.Text,
		Answers: answers,
	}
}

func AnswerToResponse(a *entities.Answer) AnswerResponse {
	return AnswerResponse{
		Id:         a.Id,
		QuestionId: a.QuestionId,
		UserId:     a.UserId,
		Text:       a.Text,
	}
}
