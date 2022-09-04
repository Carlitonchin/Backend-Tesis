package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type questionService struct {
	repo model.QuestionRepository
}

func NewQuestionService(question_repo model.QuestionRepository) model.QuestionService {
	return &questionService{
		repo: question_repo,
	}
}

func (s *questionService) AddQuestion(
	ctx context.Context, question *model.Question) (
	*model.Question, error) {

	return s.repo.CreateQuestion(ctx, question)
}
