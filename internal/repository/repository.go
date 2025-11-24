package repository

import "gorm.io/gorm"

// Repository aggregates all domain repositories
type Repository struct {
	QuestionsRepository
	AnswersRepository
}

// NewRepository creates a Repository with all concrete sub-repositories initialized
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		QuestionsRepository: NewQuestionsRepository(db),
		AnswersRepository:   NewAnswersRepository(db),
	}
}
