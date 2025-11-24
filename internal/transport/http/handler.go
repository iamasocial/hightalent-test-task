package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamasocial/hightalent-test-task/internal/service"
	"github.com/iamasocial/hightalent-test-task/internal/transport/dto"
)

// Handler handles HTTP requests for questions and answers
type Handler struct {
	service.Service
}

// NewHandler creates a new HTTP handler with injected services
func NewHandler(svc *service.Service) *Handler {
	return &Handler{Service: *svc}
}

// GetAllQuestions handles GET /questions and returns all questions
func (h *Handler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.Service.GetAllQuestions(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []dto.QuestionResponse
	for _, q := range questions {
		resp = append(resp, dto.QuestionToResponse(q))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateQuestion handles POST /questions and creates a new questions
func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q, err := h.Service.CreateQuestion(r.Context(), req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.QuestionToResponse(q)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetQuestionByID handles GET /questions/{id} and return a question by ID
func (h *Handler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q, err := h.Service.GetQuestionByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := dto.QuestionToResponse(q)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteQuestion handles DELETE /questions/{id} and removes a question by ID
func (h *Handler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteQuestion(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateAnswer handles POST /questions/{id}/answers and creates an answer
func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	questionIDStr := mux.Vars(r)["id"]
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req dto.CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a, err := h.Service.CreateAnswer(r.Context(), uint(questionID), req.UserID, req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.AnswerToResponse(a)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetAnswerByID handles GET /answers/{id} and returns an answer by ID
func (h *Handler) GetAnswerByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a, err := h.Service.GetAnswerByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := dto.AnswerToResponse(a)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteAnswer handles DELETE /answers/{id} and removes an answer
func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteAnswer(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
