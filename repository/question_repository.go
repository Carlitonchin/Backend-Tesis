package repository

import (
	"context"
	"fmt"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"gorm.io/gorm"
)

type questionRepository struct {
	DB *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) model.QuestionRepository {
	return &questionRepository{
		DB: db,
	}
}

func (s *questionRepository) CreateQuestion(
	ctx context.Context, question *model.Question) (
	*model.Question, error) {

	err := s.DB.Create(question).Error

	if err != nil {
		type_error := apperrors.Internal
		message := "Ocurrio un error mientras se insertaba la pregunta en la base de datos"

		e := apperrors.NewError(type_error, message)
		return nil, e
	}
	return question, nil
}

func (s *questionRepository) Clasify(ctx context.Context, question_id uint, area_id uint) error {
	err := s.DB.Model(&model.Question{}).Where("id = ?", question_id).Update("area_id", area_id).Error

	if err != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("No existe un area con id = '%v'", area_id)

		err = apperrors.NewError(type_error, message)
	}

	return err
}
