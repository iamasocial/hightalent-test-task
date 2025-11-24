package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/service"
	"github.com/iamasocial/hightalent-test-task/internal/service/mocks"
	"github.com/iamasocial/hightalent-test-task/internal/transport/dto"
	tr "github.com/iamasocial/hightalent-test-task/internal/transport/http"
)

func TestGetAllQuestions(t *testing.T) {
	mockQuestions := &mocks.MockQuestionsService{
		GetAllFunc: func(ctx context.Context) ([]*entities.Question, error) {
			return []*entities.Question{
				{ID: 1, Text: "Hello", CreatedAt: time.Now()},
			}, nil
		},
	}

	mockAnswers := &mocks.MockAnswersService{}

	svc := &service.Service{
		QuestionsService: mockQuestions,
		AnswersService:   mockAnswers,
	}

	h := tr.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/questions/", nil)
	w := httptest.NewRecorder()

	h.GetAllQuestions(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var body []dto.QuestionResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}

	if len(body) != 1 {
		t.Fatalf("expected 1 question, got %d", len(body))
	}

	if body[0].Text != "Hello" {
		t.Fatalf("unexpected body: %+v", body)
	}
}

func TestCreateQuestion(t *testing.T) {
	mockQuestions := &mocks.MockQuestionsService{
		CreateFunc: func(ctx context.Context, text string) (*entities.Question, error) {
			return &entities.Question{ID: 2, Text: text, CreatedAt: time.Now()}, nil
		},
	}

	mockAnswers := &mocks.MockAnswersService{}

	svc := &service.Service{
		QuestionsService: mockQuestions,
		AnswersService:   mockAnswers,
	}

	h := tr.NewHandler(svc)

	body := []byte(`{"text":"new question"}`)
	req := httptest.NewRequest(http.MethodPost, "/questions/", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.CreateQuestion(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var qr dto.QuestionResponse
	if err := json.NewDecoder(resp.Body).Decode(&qr); err != nil {
		t.Fatal(err)
	}

	if qr.Text != "new question" {
		t.Fatalf("unexpected response: %+v", qr)
	}
}

func TestGetQuestionByID(t *testing.T) {
	mockQuestions := &mocks.MockQuestionsService{
		GetByIDFunc: func(ctx context.Context, ID uint) (*entities.Question, error) {
			return &entities.Question{ID: ID, Text: "found"}, nil
		},
	}

	mockAnswers := &mocks.MockAnswersService{}

	svc := &service.Service{
		QuestionsService: mockQuestions,
		AnswersService:   mockAnswers,
	}

	h := tr.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/questions/5", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/questions/{id}", h.GetQuestionByID).Methods("GET")
	router.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var qr dto.QuestionResponse
	if err := json.NewDecoder(resp.Body).Decode(&qr); err != nil {
		t.Fatal(err)
	}

	if qr.ID != 5 {
		t.Fatalf("expected ID 5, got %d", qr.ID)
	}
}

func TestDeleteQuestions(t *testing.T) {
	mockQuestions := &mocks.MockQuestionsService{
		DeleteFunc: func(ctx context.Context, ID uint) error {
			return nil
		},
	}

	mockAnswers := &mocks.MockAnswersService{}

	svc := &service.Service{
		QuestionsService: mockQuestions,
		AnswersService:   mockAnswers,
	}

	h := tr.NewHandler(svc)

	req := httptest.NewRequest(http.MethodDelete, "/questions/8", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/questions/{id}", h.DeleteQuestion).Methods("DELETE")
	router.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", resp.StatusCode)
	}
}

func TestCreateAnswer(t *testing.T) {
	mockAnswers := &mocks.MockAnswersService{
		CreateFunc: func(ctx context.Context, questionID uint, userID, text string) (*entities.Answer, error) {
			return &entities.Answer{
				ID:         1,
				QuestionID: questionID,
				UserID:     userID,
				Text:       text,
				CreatedAt:  time.Now(),
			}, nil
		},
	}

	mockQuestions := &mocks.MockQuestionsService{}

	svc := &service.Service{
		QuestionsService: mockQuestions,
		AnswersService:   mockAnswers,
	}

	h := tr.NewHandler(svc)

	body := []byte(`{"user_id":"1234-uuid","text":"mock"}`)
	req := httptest.NewRequest(http.MethodPost, "/questions/3/answers/", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/questions/{id}/answers/", h.CreateAnswer).Methods("POST")
	router.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var ar dto.AnswerResponse
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		t.Fatal(err)
	}

	if ar.ID != 1 {
		t.Fatalf("expected ID 1, got %d", ar.ID)
	}

	if ar.QuestionID != 3 {
		t.Fatalf("expected questionID 3, got %d", ar.QuestionID)
	}

	if ar.Text != "mock" {
		t.Fatalf("expected text 'mock', got %q", ar.Text)
	}

	if ar.UserID != "1234-uuid" {
		t.Fatalf("expected userID '1234-uuid', got %q", ar.UserID)
	}
}

func TestGetAnswerByID(t *testing.T) {
	mockAnswers := &mocks.MockAnswersService{
		GetByIDFunc: func(ctx context.Context, ID uint) (*entities.Answer, error) {
			return &entities.Answer{
				ID:         1,
				QuestionID: 4,
				UserID:     "1234-uuid",
				Text:       "mux",
				CreatedAt:  time.Now(),
			}, nil
		},
	}

	mockQuestions := mocks.MockQuestionsService{}

	svc := &service.Service{
		QuestionsService: &mockQuestions,
		AnswersService:   mockAnswers,
	}

	h := tr.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/answers/1", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/answers/{id}", h.GetAnswerByID).Methods("GET")
	router.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var ar dto.AnswerResponse
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		t.Fatal(err)
	}

	if ar.ID != 1 {
		t.Fatalf("expected ID 1, got %d", ar.ID)
	}
}

func TestDeleteAnswer(t *testing.T) {
	mockAnswers := &mocks.MockAnswersService{
		DeleteFunc: func(ctx context.Context, ID uint) error {
			return nil
		},
	}

	mockQuestions := &mocks.MockQuestionsService{}

	svc := &service.Service{
		QuestionsService: mockQuestions,
		AnswersService:   mockAnswers,
	}

	h := tr.NewHandler(svc)

	req := httptest.NewRequest(http.MethodDelete, "/answers/1", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/answers/{id}", h.DeleteAnswer).Methods("DELETE")
	router.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", resp.StatusCode)
	}
}
