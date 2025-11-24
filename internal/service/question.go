package service

import (
	"context"
	"fmt"
	"time"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/repository"
)

type QuestionsService interface {
	Create(ctx context.Context, text string) (*entities.Question, error)
	GetAll(ctx context.Context) ([]*entities.Question, error)
	GetbyID(ctx context.Context, id uint) (*entities.Question, error)
	Delete(ctx context.Context, id uint) error
}

type questionsService struct {
	repo repository.QuestionsRepository
}

func NewQuestionsService(repo repository.QuestionsRepository) QuestionsService {
	return &questionsService{repo: repo}
}

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

func (s *questionsService) GetAll(ctx context.Context) ([]*entities.Question, error) {
	return s.repo.GetAll(ctx)
}

func (s *questionsService) GetbyID(ctx context.Context, id uint) (*entities.Question, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *questionsService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
