package mocks

import (
	"context"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
)

type MockAnswersRepository struct {
	CreateFunc  func(ctx context.Context, a *entities.Answer) error
	GetByIDFunc func(ctx context.Context, ID uint) (*entities.Answer, error)
	DeleteFunc  func(ctx context.Context, ID uint) error
}

func (m *MockAnswersRepository) Create(ctx context.Context, a *entities.Answer) error {
	return m.CreateFunc(ctx, a)
}

func (m *MockAnswersRepository) GetByID(ctx context.Context, ID uint) (*entities.Answer, error) {
	return m.GetByIDFunc(ctx, ID)
}

func (m *MockAnswersRepository) Delete(ctx context.Context, ID uint) error {
	return m.DeleteFunc(ctx, ID)
}
