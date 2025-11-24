package dto

type CreateAnswerRequest struct {
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}

type AnswerResponse struct {
	Id         uint   `json:"id"`
	QuestionId uint   `json:"question_id"`
	UserId     string `json:"user_id"`
	Text       string `json:"text"`
}
