package mocks

import (
	"context"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/service"
)

type MockQuestionsService struct {
	GetAllFunc  func(ctx context.Context) ([]*entities.Question, error)
	CreateFunc  func(ctx context.Context, text string) (*entities.Question, error)
	GetByIDFunc func(ctx context.Context, ID uint) (*entities.Question, error)
	DeleteFunc  func(ctx context.Context, ID uint) error
}

var _ service.QuestionsService = (*MockQuestionsService)(nil)

func (m *MockQuestionsService) GetAllQuestions(ctx context.Context) ([]*entities.Question, error) {
	return m.GetAllFunc(ctx)
}

func (m *MockQuestionsService) CreateQuestion(ctx context.Context, text string) (*entities.Question, error) {
	return m.CreateFunc(ctx, text)
}

func (m *MockQuestionsService) GetQuestionByID(ctx context.Context, ID uint) (*entities.Question, error) {
	return m.GetByIDFunc(ctx, ID)
}

func (m *MockQuestionsService) DeleteQuestion(ctx context.Context, ID uint) error {
	return m.DeleteFunc(ctx, ID)
}
