package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/Carlitonchin/Backend-Tesis/some_utils"
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

	status_send_id, err := some_utils.GetUintEnv("STATUS_SEND_CODE")
	if err != nil {
		type_error := apperrors.Internal
		message := "No se pudo leer el id del stado 'enviada'"

		e := apperrors.NewError(type_error, message)
		return nil, e
	}

	question.StatusId = status_send_id
	return s.repo.CreateQuestion(ctx, question)
}

func (s *questionService) Clasify(ctx context.Context, question_id uint, area_id uint) error {
	return s.repo.Clasify(ctx, question_id, area_id)
}
