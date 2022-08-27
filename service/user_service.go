package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type UserService struct {
	UserRepository model.UserRepository
}

type USConfing struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfing) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

func (s *UserService) GetById(ctx context.Context, id uint) (*model.User, error) {
	return s.UserRepository.FindById(ctx, id)
}

func (s *UserService) SignUp(ctx context.Context, user *model.User) error {

	hashed_pass, err := hashPass(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashed_pass

	err = s.UserRepository.Create(ctx, user)

	return err
}
