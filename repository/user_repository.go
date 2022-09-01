package repository

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {

	return &userRepository{
		DB: db,
	}
}

func (s *userRepository) Create(ctx context.Context, user *model.User) error {
	result := s.DB.Create(user)

	return result.Error
}

func (s *userRepository) FindById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	err := s.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	role, err := s.GetRoleById(ctx, user.RoleID)

	if err != nil {
		return nil, err
	}

	user.Role = &model.Role{
		Name: role.Name,
	}

	return &user, nil
}

func (s *userRepository) GetRoleById(ctx context.Context, id uint) (*model.Role, error) {
	var role *model.Role

	err := s.DB.First(&role, id).Error

	if err != nil {
		return nil, err
	}

	return role, err
}

func (s *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	tx := s.DB.Where("email = ?", email).First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	role, err := s.GetRoleById(ctx, user.RoleID)

	if err != nil {
		return nil, err
	}

	user.Role = &model.Role{
		Name: role.Name,
	}

	return &user, nil
}
