package repository

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) model.RoleRepository {
	return &roleRepository{
		DB: db,
	}
}

func (s *roleRepository) GetRoleById(ctx context.Context, id uint) (*model.Role, error) {
	var role *model.Role

	err := s.DB.First(&role, id).Error

	return role, err
}
