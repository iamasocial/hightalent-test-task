package dto

import "time"

// CreateQuestionRequest represents the request payload to create a question
type CreateQuestionRequest struct {
	Text string `json:"text"`
}

// QuestionResponse represents a question with optional answers in the response
type QuestionResponse struct {
	ID       uint             `json:"id"`
	Text     string           `json:"text"`
	CreateAt time.Time        `json:"created_at"`
	Answers  []AnswerResponse `json:"answers,omitempty"`
}
