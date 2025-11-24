package service

import (
	"context"
	"fmt"
	"time"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/repository"
)

type AnswersService interface {
	Create(ctx context.Context, questionId uint, userId, text string) (*entities.Answer, error)
	GetByID(ctx context.Context, id uint) (*entities.Answer, error)
	Delete(ctx context.Context, id uint) error
}

type answersService struct {
	questionRepo repository.QuestionsRepository
	answersRepo  repository.AnswersRepository
}

func NewAnswersService(qr repository.QuestionsRepository, ar repository.AnswersRepository) AnswersService {
	return &answersService{
		questionRepo: qr,
		answersRepo:  ar,
	}
}

func (s *answersService) Create(ctx context.Context, questionId uint, userId, text string) (*entities.Answer, error) {
	_, err := s.questionRepo.GetByID(ctx, questionId)
	if err != nil {
		return nil, fmt.Errorf("question not found")
	}

	a := &entities.Answer{
		QuestionId: questionId,
		UserId:     userId,
		Text:       text,
		CreatedAt:  time.Now(),
	}

	if err := s.answersRepo.Create(ctx, a); err != nil {
		return nil, err
	}

	return a, nil
}

func (s *answersService) GetByID(ctx context.Context, id uint) (*entities.Answer, error) {
	return s.answersRepo.GetByID(ctx, id)
}

func (s *answersService) Delete(ctx context.Context, id uint) error {
	return s.answersRepo.Delete(ctx, id)
}
