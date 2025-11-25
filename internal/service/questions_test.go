package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"github.com/iamasocial/hightalent-test-task/internal/repository/mocks"
	"github.com/iamasocial/hightalent-test-task/internal/service"
)

func TestQuestionsServiceCreate(t *testing.T) {
	timestamp := time.Now()
	repo := &mocks.MockQuestionsRepository{
		CreateFunc: func(ctx context.Context, q *entities.Question) error {
			q.ID = 1
			q.CreatedAt = timestamp
			return nil
		},
	}

	svc := service.NewQuestionsService(repo)

	text := "what?"
	q, err := svc.CreateQuestion(context.Background(), text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if q.ID != 1 {
		t.Fatalf("expected ID 1, got %d", q.ID)
	}

	if q.Text != text {
		t.Fatalf("expected text %q, got %q", text, q.Text)
	}

	if q.CreatedAt != timestamp {
		t.Fatalf("expected CreatedAt %v, got %v", timestamp, q.CreatedAt)
	}
}

func TestQuestionsServiceGetAll(t *testing.T) {
	q1 := "q1"
	q2 := "q2"
	repo := &mocks.MockQuestionsRepository{
		GetAllFunc: func(ctx context.Context) ([]*entities.Question, error) {
			return []*entities.Question{
				{
					ID:   1,
					Text: q1,
				},

				{
					ID:   2,
					Text: q2,
				},
			}, nil
		},
	}

	svc := service.NewQuestionsService(repo)
	qs, err := svc.GetAllQuestions(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(qs) != 2 {
		t.Fatalf("expected len 2, got %d", len(qs))
	}

	if qs[0].ID != 1 {
		t.Fatalf("expected qs[0].ID = 1, got %d", qs[0].ID)
	}

	if qs[0].Text != q1 {
		t.Fatalf("expected qs[0].Text = %q, got %q", q1, qs[0].Text)
	}

	if qs[1].ID != 2 {
		t.Fatalf("expected qs[1].ID = 2, got %d", qs[1].ID)
	}

	if qs[1].Text != q2 {
		t.Fatalf("expected qs[1].Text = %q, got %q", q2, qs[1].Text)
	}
}

func TestQuestionsServiceGetByID(t *testing.T) {
	repo := &mocks.MockQuestionsRepository{
		GetByIDFunc: func(ctx context.Context, ID uint) (*entities.Question, error) {
			return &entities.Question{
				ID: 12,
			}, nil
		},
	}

	svc := service.NewQuestionsService(repo)

	q, err := svc.GetQuestionByID(context.Background(), 12)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if q.ID != 12 {
		t.Fatalf("expected ID 12, got %d", q.ID)
	}
}

func TestQuestionsServiceDelete(t *testing.T) {
	called := false
	repo := &mocks.MockQuestionsRepository{
		DeleteFunc: func(ctx context.Context, ID uint) error {
			called = true
			if ID != 12 {
				return fmt.Errorf("not found")
			}
			return nil
		},
	}

	svc := service.NewQuestionsService(repo)

	if err := svc.DeleteQuestion(context.Background(), 12); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !called {
		t.Fatalf("expected DeleteFunc to be called")
	}
}
