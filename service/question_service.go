package service

import (
	"context"
	"fmt"

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
	question, err := s.repo.GetById(ctx, question_id)
	if err != nil {
		return err
	}

	status_sended_id, err := some_utils.GetUintEnv("STATUS_SEND_CODE")
	if err != nil {
		return err
	}

	if question.StatusId != status_sended_id {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("Solo se pueden clasificar de esa manera las preguntas con status_id = '%v'", status_sended_id)

		e := apperrors.NewError(type_error, message)
		return e
	}

	return s.repo.Clasify(ctx, question_id, area_id)
}
