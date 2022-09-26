package repository

import (
	"context"
	"time"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"gorm.io/gorm"
)

type chatRepository struct {
	DB *gorm.DB
}

func NewChatRepository(db *gorm.DB) model.ChatRepository {
	return &chatRepository{
		DB: db,
	}
}

func (s *chatRepository) SendMessage(ctx context.Context, question_id uint, user_id uint, text string) error {
	message := &model.MessageChat{
		Question: question_id,
		User_Id:  user_id,
		Text:     text,
		Readed:   false,
	}

	err := s.DB.Create(message).Error

	if err != nil {
		type_error := apperrors.Internal
		message := "Ocurrio un error inesperado mientras se escribia un mensaje"
		err = apperrors.NewError(type_error, message)
	}

	return err
}
func (s *chatRepository) GetMessages(ctx context.Context, question_id uint) ([]model.MessageChat, error) {
	var messages []model.MessageChat

	rows, err := s.DB.Table("message_chats").
		Select("users.name, message_chats.text, message_chats.created_at").
		Joins("left join users on message_chats.user_id = users.id").
		Where("message_chats.question = ?", question_id).
		Rows()

	if err != nil {
		type_error := apperrors.Internal
		message := "Ocurrio un error inesperado mientras se recuperaban mensajes de chat"
		err = apperrors.NewError(type_error, message)
		return nil, err
	}

	messages = make([]model.MessageChat, 0)

	for rows.Next() {
		var message_text string
		var user_name string
		var creation_date time.Time

		rows.Scan(&user_name, &message_text, &creation_date)

		m := model.MessageChat{
			UserName: user_name,
			Text:     message_text,
			Model: gorm.Model{
				CreatedAt: creation_date,
			},
		}

		messages = append(messages, m)
	}

	return messages, nil
}

func (s *chatRepository) ReadMessages(ctx context.Context, question_id uint, user_id uint) error {

	err := s.DB.Model(&model.MessageChat{}).
		Where("question = ? and not user_id = ?", question_id, user_id).
		Update("readed", true).
		Error

	if err != nil {
		type_error := apperrors.Internal
		message := "Ocurrio un error inesperado mientras se modificaban preguntas"
		err = apperrors.NewError(type_error, message)
	}

	return err
}
