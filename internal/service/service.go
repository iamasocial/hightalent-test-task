package service

import "github.com/iamasocial/hightalent-test-task/internal/repository"

type Service struct {
	QuestionsService
	AnswersService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		QuestionsService: NewQuestionsService(repos.QuestionsRepository),
		AnswersService:   NewAnswersService(repos.QuestionsRepository, repos.AnswersRepository),
	}
}
