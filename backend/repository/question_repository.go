package repository

import (
	"context"
	"fmt"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/Carlitonchin/Backend-Tesis/some_utils"
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
	status_clasified_id, eos := some_utils.GetUintEnv("STATUS_CLASIFIED1_CODE")

	if eos != nil {
		return eos
	}
	err := s.DB.Model(&model.Question{}).Where("id = ?", question_id).Updates(
		map[string]interface{}{"area_id": area_id, "status_id": status_clasified_id}).Error

	if err != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("No existe un area con id = '%v'", area_id)

		err = apperrors.NewError(type_error, message)
	}

	return err
}

func (s *questionRepository) GetById(ctx context.Context, question_id uint) (*model.Question, error) {
	var question model.Question
	err := s.DB.First(&question, question_id).Error

	if err != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("No exite pregunta con id = '%v'", question_id)
		err = apperrors.NewError(type_error, message)
	}

	return &question, err
}

func (s *questionRepository) TakeQuestion(ctx context.Context, user_id uint, question_id uint) error {
	err := s.DB.Model(&model.Question{}).Where("id = ?", question_id).Update(
		"user_responsible", user_id).Error

	if err != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("El usuario con id = '%v' no existe", user_id)
		err = apperrors.NewError(type_error, message)
	}

	return err
}

func (s *questionRepository) ResponseQuestion(ctx context.Context, question_id uint, response string) error {
	status_finish_code, err := some_utils.GetUintEnv("STATUS_FINISH_CODE")

	if err != nil {
		return err
	}

	err = s.DB.Model(&model.Question{}).Where("id = ?", question_id).Updates(
		map[string]interface{}{"response": response, "status_id": status_finish_code}).Error

	if err != nil {
		type_error := apperrors.Internal
		message := fmt.Sprintf(
			"Ocurrio un error inesperado mientras se respondia la pregunta con id = '%v'", question_id)

		err = apperrors.NewError(type_error, message)
	}

	return err

}

func (s *questionRepository) UpLevel(ctx context.Context, question_id uint) error {
	clasify2_code, err := some_utils.GetUintEnv("STATUS_CLASIFIED2_CODE")
	if err != nil {
		return err
	}

	err = s.DB.Model(&model.Question{}).Where("id = ?", question_id).Updates(
		map[string]interface{}{"user_responsible": nil, "status_id": clasify2_code}).Error

	if err != nil {
		type_error := apperrors.Internal
		message := fmt.Sprintf(
			"Ocurrio un error inesperado mientras se subia de nivel la pregunta con id = '%v'", question_id)

		err = apperrors.NewError(type_error, message)
	}

	return err
}

func (s *questionRepository) UpToAdmin(ctx context.Context, question_id uint) error {
	status_admin_code, err := some_utils.GetUintEnv("STATUS_ADMIN_CODE")
	if err != nil {
		return err
	}

	err = s.DB.Model(&model.Question{}).Where("id = ?", question_id).Updates(
		map[string]interface{}{"user_responsible": nil, "status_id": status_admin_code}).Error

	if err != nil {
		type_error := apperrors.Internal
		message := fmt.Sprintf(
			"Ocurrio un error inesperado mientras se subia de nivel la pregunta con id = '%v'", question_id)

		err = apperrors.NewError(type_error, message)
	}

	return err
}

func (s *questionRepository) GetMyQuestions(ctx context.Context, user_id uint) ([]model.Question, error) {
	var questions []model.Question
	err := s.DB.Where("user_refer = ?", user_id).Find(&questions).Error

	if err != nil {
		type_error := apperrors.Internal
		message := fmt.Sprintf("Ocurrió un error inesperado en la base de datos mientras se buscaban las preguntas del usuario con id = '%v'", user_id)
		err = apperrors.NewError(type_error, message)
	}

	return questions, err
}

func (s *questionRepository) GetUnClasifiedQuestions(ctx context.Context) ([]model.Question, error) {
	var questions []model.Question
	status_send, err := some_utils.GetUintEnv("STATUS_SEND_CODE")
	if err != nil {
		return nil, err
	}

	err = s.DB.Where("status_id = ?", status_send).Find(&questions).Error

	if err != nil {
		type_error := apperrors.Internal
		message := fmt.Sprintf("Ocurrió un error inesperado en la base de datos mientras se buscaban las preguntas sin clasificar")
		err = apperrors.NewError(type_error, message)
	}

	return questions, err
}

func (s *questionRepository) GetQuestionsByStatus(ctx context.Context, status uint, area_id *uint) ([]model.Question, error) {
	var questions []model.Question
	var err error
	if area_id == nil {
		err = s.DB.Where("status_id = ? AND user_responsible is null", status).Find(&questions).Error
	} else {
		err = s.DB.Where("status_id = ? AND area_id = ? AND user_responsible is null", status, *area_id).Find(&questions).Error
	}

	if err != nil {
		type_error := apperrors.Internal
		message := fmt.Sprintf("Ocurrio un error inesperado mientras se recuperaban las preguntas con status_id = %v", status)
		err = apperrors.NewError(type_error, message)
	}

	return questions, err
}

func (s *questionRepository) GetTakenQuestions(ctx context.Context, user_id uint) ([]model.Question, error) {
	var questions []model.Question
	err := s.DB.Where("user_responsible = ? AND response = ?", user_id, "").Find(&questions).Error

	if err != nil {
		type_error := apperrors.Internal
		message := "Ocurrio un error inesperado mientras se recuperaban algunas preguntas"
		err = apperrors.NewError(type_error, message)
	}

	return questions, err
}

func (s *questionRepository) GetTakenQuestionById(ctx context.Context, question_id uint, user_id uint) (*model.Question, error) {
	var question model.Question
	err := s.DB.First(&question, "id = ? AND user_responsible = ?", question_id, user_id).Error

	if err != nil {
		type_error := apperrors.NotFound
		message := fmt.Sprintf("La pregunta con id = '%v' o no existe o tiene otro responsable", question_id)
		err = apperrors.NewError(type_error, message)
	}

	return &question, err
}

func (s *questionRepository) GetQuestionById(ctx context.Context, question_id uint) (*model.Question, error) {
	var question model.Question

	err := s.DB.First(&question, question_id).Error

	if err != nil {
		type_error := apperrors.NotFound
		message := fmt.Sprintf("No existe ninguna pregunta con id = '%d'", question_id)
		err = apperrors.NewError(type_error, message)
	}

	return &question, err
}
