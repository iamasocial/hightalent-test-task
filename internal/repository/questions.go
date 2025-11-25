package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/iamasocial/hightalent-test-task/internal/db"
	"github.com/iamasocial/hightalent-test-task/internal/entities"
)

// QuestionsRepository defines operations for managing questions in the storage layer
type QuestionsRepository interface {
	GetAll(ctx context.Context) ([]*entities.Question, error)
	Create(ctx context.Context, q *entities.Question) error
	GetByID(ctx context.Context, ID uint) (*entities.Question, error)
	Delete(ctx context.Context, ID uint) error
}

type questionRepo struct {
	db *gorm.DB
}

// NewQuestionsRepository creates a new instance of QuestionsRepository
func NewQuestionsRepository(db *gorm.DB) QuestionsRepository {
	return &questionRepo{db: db}
}

// GetAll returns all answers
func (r *questionRepo) GetAll(ctx context.Context) ([]*entities.Question, error) {
	var models []db.Question
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	questions := make([]*entities.Question, 0, len(models))
	for _, m := range models {
		questions = append(questions, toEntityQuestion(m))
	}

	return questions, nil
}

// Create inserts a new question record into the database
func (r *questionRepo) Create(ctx context.Context, q *entities.Question) error {
	model := toModelQuestion(q)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	q.ID = model.ID
	return nil
}

// GetByID returns question by ID or ErrNotFound if it doesn't exist
func (r *questionRepo) GetByID(ctx context.Context, ID uint) (*entities.Question, error) {
	var model db.Question

	err := r.db.WithContext(ctx).Preload("Answers").First(&model, ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return toEntityQuestion(model), err
}

// Delete removes question by ID or returns ErrNotFound if it doesn't exist
func (r *questionRepo) Delete(ctx context.Context, ID uint) error {
	err := r.db.WithContext(ctx).Delete(&db.Question{}, ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}

	return err
}
