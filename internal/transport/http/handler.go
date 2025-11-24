package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamasocial/hightalent-test-task/internal/service"
	"github.com/iamasocial/hightalent-test-task/internal/transport/dto"
)

type Handler struct {
	service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{Service: *svc}
}

// GET /questions
func (h *Handler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.QuestionsService.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []dto.QuestionResponse
	for _, q := range questions {
		resp = append(resp, dto.QuestionToResponse(q))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /questions
func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q, err := h.QuestionsService.Create(r.Context(), req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.QuestionToResponse(q)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /questions/{id}
func (h *Handler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	q, err := h.QuestionsService.GetbyID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := dto.QuestionToResponse(q)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /questions/{id}
func (h *Handler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	err := h.QuestionsService.Delete(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// POST /questions/{id}/answers
func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	questionIdStr := mux.Vars(r)["id"]
	questionId, _ := strconv.Atoi(questionIdStr)

	var req dto.CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a, err := h.AnswersService.Create(r.Context(), uint(questionId), req.UserId, req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.AnswerToResponse(a)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /answers/{id}
func (h *Handler) GetAnswerById(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	a, err := h.AnswersService.GetByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := dto.AnswerToResponse(a)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /answers/{id}
func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	err := h.AnswersService.Delete(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
