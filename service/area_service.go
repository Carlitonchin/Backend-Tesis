package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type areaService struct {
	repo model.AreaRepository
}

func NewAreaService(area_repo model.AreaRepository) model.AreaService {
	return &areaService{
		repo: area_repo,
	}
}

func (s *areaService) AddArea(ctx context.Context, area *model.Area) (*model.Area, error) {
	return s.repo.CreateArea(ctx, area)
}

func (s *areaService) Clasify(ctx context.Context, question_id uint, area_id uint) error {
	return s.repo.Clasify(ctx, question_id, area_id)
}
