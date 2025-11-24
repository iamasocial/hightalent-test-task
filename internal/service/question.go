package service

import (
	"context"
	"fmt"
	"time"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/repository"
)

// QuestionsService defines business logic operations for questions
type QuestionsService interface {
	Create(ctx context.Context, text string) (*entities.Question, error)
	GetAll(ctx context.Context) ([]*entities.Question, error)
	GetByID(ctx context.Context, id uint) (*entities.Question, error)
	Delete(ctx context.Context, id uint) error
}

type questionsService struct {
	repo repository.QuestionsRepository
}

// NewQuestionsService creates a new QuestionsService instance
func NewQuestionsService(repo repository.QuestionsRepository) QuestionsService {
	return &questionsService{repo: repo}
}

// Create adds a new question after validating the text
func (s *questionsService) Create(ctx context.Context, text string) (*entities.Question, error) {
	if text == "" {
		return nil, fmt.Errorf("question text cannot be empty")
	}

	q := &entities.Question{
		Text:      text,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, q); err != nil {
		return nil, err
	}

	return q, nil
}

// GetAll returns all questions
func (s *questionsService) GetAll(ctx context.Context) ([]*entities.Question, error) {
	return s.repo.GetAll(ctx)
}

// GetById retrives a question by its ID
func (s *questionsService) GetByID(ctx context.Context, id uint) (*entities.Question, error) {
	return s.repo.GetByID(ctx, id)
}

// Delete removes a question by its ID
func (s *questionsService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
