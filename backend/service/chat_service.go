package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type chatService struct {
	ChatRepository model.ChatRepository
}

func NewChatService(chat_repo model.ChatRepository) model.ChatService {
	return &chatService{
		ChatRepository: chat_repo,
	}
}

func (s *chatService) SendMessage(ctx context.Context, question_id uint, user_id uint, text string) (*model.MessageChat, error) {
	return s.ChatRepository.SendMessage(ctx, question_id, user_id, text)
}

func (s *chatService) GetMessages(ctx context.Context, question_id uint, user_id uint) ([]model.MessageChat, error) {
	messages, err := s.ChatRepository.GetMessages(ctx, question_id)

	if err != nil {
		return nil, err
	}

	if len(messages) > 0 {
		err = s.ChatRepository.ReadMessages(ctx, question_id, user_id)
	}

	return messages, err
}
