package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates and configures an HTTP router with all application routes
func NewRouter(h *Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/questions/", h.CreateQuestion).Methods("POST")
	r.HandleFunc("/questions/", h.GetAllQuestions).Methods("GET")
	r.HandleFunc("/questions/{id:[0-9]+}", h.GetQuestionByID).Methods("GET")
	r.HandleFunc("/questions/{id:[0-9]+}", h.DeleteQuestion).Methods("DELETE")

	r.HandleFunc("/questions/{id:[0-9]+}/answers/", h.CreateAnswer).Methods("POST")
	r.HandleFunc("/answers/{id:[0-9]+}", h.GetAnswerByID).Methods("GET")
	r.HandleFunc("/answers/{id:[0-9]+}", h.DeleteAnswer).Methods("DELETE")

	return r
}
