package repository

import (
	"context"
	"errors"

	"github.com/iamasocial/hightalent-test-task/internal/db"
	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"gorm.io/gorm"
)

// AnswersRepository defines operations for managing answers in the storage layer
type AnswersRepository interface {
	Create(ctx context.Context, a *entities.Answer) error
	GetByID(ctx context.Context, ID uint) (*entities.Answer, error)
	Delete(ctx context.Context, ID uint) error
}

type answersRepo struct {
	db *gorm.DB
}

// NewAnswersRepository creates a new instance of AnswersRepository
func NewAnswersRepository(db *gorm.DB) AnswersRepository {
	return &answersRepo{db: db}
}

// Create inserts a new answer record into the database
func (r *answersRepo) Create(ctx context.Context, a *entities.Answer) error {
	model := toModelAnswer(a)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}

	a.ID = model.ID
	return nil
}

// GetByID returns an answer by its ID or ErrNotFound if it doesn't exists
func (r *answersRepo) GetByID(ctx context.Context, ID uint) (*entities.Answer, error) {
	var model db.Answer
	err := r.db.WithContext(ctx).First(&model, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return toEntityAnswer(model), nil
}

// Delete removes an answer by its ID or ErrNotFound if it doesn't exists
func (r *answersRepo) Delete(ctx context.Context, ID uint) error {
	err := r.db.WithContext(ctx).Delete(&db.Answer{}, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}

	return err
}
