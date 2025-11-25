package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/repository/mocks"
	"github.com/iamasocial/hightalent-test-task/internal/service"
)

func TestAnswersServiceCreate(t *testing.T) {
	ansRepo := &mocks.MockAnswersRepository{
		CreateFunc: func(ctx context.Context, a *entities.Answer) error {
			a.ID = 12
			return nil
		},
	}

	qRepo := &mocks.MockQuestionsRepository{
		GetByIDFunc: func(ctx context.Context, ID uint) (*entities.Question, error) {
			return &entities.Question{
				ID: 1,
			}, nil
		},
	}

	svc := service.NewAnswersService(qRepo, ansRepo)

	a, err := svc.CreateAnswer(context.Background(), 1, "test", "text")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if a.ID != 12 {
		t.Fatalf("expected ID 12, got %d", a.ID)
	}
}

func TestAnswersServiceGetByID(t *testing.T) {
	ansRepo := &mocks.MockAnswersRepository{
		GetByIDFunc: func(ctx context.Context, ID uint) (*entities.Answer, error) {
			return &entities.Answer{
				ID: 1,
			}, nil
		},
	}

	qRepo := &mocks.MockQuestionsRepository{}

	svc := service.NewAnswersService(qRepo, ansRepo)

	a, err := svc.GetAnswerByID(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if a.ID != 1 {
		t.Fatalf("expected ID 1, got %d", a.ID)
	}
}

func TestAnswersServiceDelete(t *testing.T) {
	called := false
	ansRepo := &mocks.MockAnswersRepository{
		DeleteFunc: func(ctx context.Context, ID uint) error {
			called = true

			if ID != 12 {
				return fmt.Errorf("not found")
			}

			return nil
		},
	}

	qRepo := &mocks.MockQuestionsRepository{}

	svc := service.NewAnswersService(qRepo, ansRepo)

	if err := svc.DeleteAnswer(context.Background(), 12); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !called {
		t.Fatalf("expected DeleteFunc to be called")
	}
}
