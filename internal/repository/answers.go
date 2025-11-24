package repository

import (
	"context"
	"errors"

	"github.com/iamasocial/hightalent-test-task/internal/db"
	"github.com/iamasocial/hightalent-test-task/internal/entities"
	"gorm.io/gorm"
)

type AnswersRepository interface {
	Create(ctx context.Context, a *entities.Answer) error
	GetByID(ctx context.Context, id uint) (*entities.Answer, error)
	Delete(ctx context.Context, id uint) error
}

type answersRepo struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) AnswersRepository {
	return &answersRepo{db: db}
}

func (r *answersRepo) Create(ctx context.Context, a *entities.Answer) error {
	model := toModelAnswer(a)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}

	a.Id = model.Id
	return nil
}

func (r *answersRepo) GetByID(ctx context.Context, id uint) (*entities.Answer, error) {
	var model db.Answer
	err := r.db.WithContext(ctx).First(&model, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return toEntityAnswer(model), nil
}

func (r *answersRepo) Delete(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&db.Answer{}, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}

	return err
}
