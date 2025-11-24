package dto

type CreateQuestionRequest struct {
	Text string `json:"text"`
}

type QuestionResponse struct {
	Id      uint             `json:"id"`
	Text    string           `json:"text"`
	Answers []AnswerResponse `json:"answers,omitempty"`
}
