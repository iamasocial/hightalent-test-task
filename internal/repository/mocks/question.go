package mocks

import (
	"context"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
)

type MockQuestionsRepository struct {
	GetAllFunc  func(ctx context.Context) ([]*entities.Question, error)
	CreateFunc  func(ctx context.Context, q *entities.Question) error
	GetByIDFunc func(ctx context.Context, ID uint) (*entities.Question, error)
	DeleteFunc  func(ctx context.Context, ID uint) error
}

func (m *MockQuestionsRepository) GetAll(ctx context.Context) ([]*entities.Question, error) {
	return m.GetAllFunc(ctx)
}

func (m *MockQuestionsRepository) Create(ctx context.Context, q *entities.Question) error {
	return m.CreateFunc(ctx, q)
}

func (m *MockQuestionsRepository) GetByID(ctx context.Context, ID uint) (*entities.Question, error) {
	return m.GetByIDFunc(ctx, ID)
}

func (m *MockQuestionsRepository) Delete(ctx context.Context, ID uint) error {
	return m.DeleteFunc(ctx, ID)
}
