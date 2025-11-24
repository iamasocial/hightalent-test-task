package dto

import "time"

// CreateAnswerRequest represents the request body for creating an answer
type CreateAnswerRequest struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

// AnswerResponse represents the response body for an answer
type AnswerResponse struct {
	ID         uint      `json:"id"`
	QuestionID uint      `json:"question_id"`
	UserID     string    `json:"user_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
