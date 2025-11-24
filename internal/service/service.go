package service

import "github.com/iamasocial/hightalent-test-task/internal/repository"

// Service aggregates all domain services
type Service struct {
	QuestionsService
	AnswersService
}

// NewService creates a new Service with all sub-services initialized
func NewService(repos *repository.Repository) *Service {
	return &Service{
		QuestionsService: NewQuestionsService(repos.QuestionsRepository),
		AnswersService:   NewAnswersService(repos.QuestionsRepository, repos.AnswersRepository),
	}
}
