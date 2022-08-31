package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type userService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &userService{
		UserRepository: c.UserRepository,
	}
}

func (s *userService) GetById(ctx context.Context, id uint) (*model.User, error) {
	return s.UserRepository.FindById(ctx, id)
}

func (s *userService) SignUp(ctx context.Context, user *model.User) error {

	hashed_pass, err := hashPass(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashed_pass

	err = s.UserRepository.Create(ctx, user)

	return err
}

func (s *userService) SignIn(ctx context.Context, user *model.User) error {
	panic("Not implemented")
}
