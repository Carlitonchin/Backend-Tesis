package repository

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB             *gorm.DB
	roleRepository *roleRepository
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	rp := &roleRepository{
		DB: db,
	}

	return &userRepository{
		DB:             db,
		roleRepository: rp,
	}
}

func (s *userRepository) Create(ctx context.Context, user *model.User) error {
	result := s.DB.Create(user)

	return result.Error
}

func (s *userRepository) FindById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	result := s.DB.First(&user, id)

	return &user, result.Error
}

func (s *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	tx := s.DB.Where("email = ?", email).First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	role, err := s.roleRepository.GetRoleById(ctx, user.RoleID)

	if err != nil {
		return nil, err
	}

	user.Role = &model.Role{
		Name: role.Name,
	}

	return &user, nil
}
