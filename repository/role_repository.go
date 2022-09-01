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

func (s *roleRepository) GetRoles(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role

	err := s.DB.Find(&roles).Error

	return roles, err
}
