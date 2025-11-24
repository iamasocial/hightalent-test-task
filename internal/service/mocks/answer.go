package mocks

import (
	"context"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/service"
)

type MockAnswersService struct {
	CreateFunc  func(ctx context.Context, questionID uint, userID, text string) (*entities.Answer, error)
	GetByIDFunc func(ctx context.Context, ID uint) (*entities.Answer, error)
	DeleteFunc  func(ctx context.Context, ID uint) error
}

var _ service.AnswersService = (*MockAnswersService)(nil)

func (m *MockAnswersService) CreateAnswer(ctx context.Context, questionID uint, userID, text string) (*entities.Answer, error) {
	return m.CreateFunc(ctx, questionID, userID, text)
}

func (m *MockAnswersService) GetAnswerByID(ctx context.Context, ID uint) (*entities.Answer, error) {
	return m.GetByIDFunc(ctx, ID)
}

func (m *MockAnswersService) DeleteAnswer(ctx context.Context, ID uint) error {
	return m.DeleteFunc(ctx, ID)
}
