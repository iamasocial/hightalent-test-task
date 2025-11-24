package service

import (
	"context"
	"fmt"
	"time"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/repository"
)

// AnswersService defines business logic operations for answers
type AnswersService interface {
	CreateAnswer(ctx context.Context, questionID uint, userID, text string) (*entities.Answer, error)
	GetAnswerByID(ctx context.Context, ID uint) (*entities.Answer, error)
	DeleteAnswer(ctx context.Context, ID uint) error
}

type answersService struct {
	questionRepo repository.QuestionsRepository
	answersRepo  repository.AnswersRepository
}

// NewAnswersService creates a new AnswersService instance
func NewAnswersService(qr repository.QuestionsRepository, ar repository.AnswersRepository) AnswersService {
	return &answersService{
		questionRepo: qr,
		answersRepo:  ar,
	}
}

// Create adds a new answer to a question after verifying the question exists
func (s *answersService) CreateAnswer(ctx context.Context, questionID uint, userID, text string) (*entities.Answer, error) {
	_, err := s.questionRepo.GetByID(ctx, questionID)
	if err != nil {
		return nil, fmt.Errorf("question not found")
	}

	a := &entities.Answer{
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
		CreatedAt:  time.Now(),
	}

	if err := s.answersRepo.Create(ctx, a); err != nil {
		return nil, err
	}

	return a, nil
}

// GetByID retieves an answer by its ID
func (s *answersService) GetAnswerByID(ctx context.Context, ID uint) (*entities.Answer, error) {
	return s.answersRepo.GetByID(ctx, ID)
}

// Delete removes an answer by its ID
func (s *answersService) DeleteAnswer(ctx context.Context, ID uint) error {
	return s.answersRepo.Delete(ctx, ID)
}
