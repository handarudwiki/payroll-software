package services

import (
	"context"
	"errors"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/models/commons"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"github.com/handarudwiki/payroll-sistem/internal/responses"
	"gorm.io/gorm"
)

type (
	Position interface {
		Create(ctx context.Context, dto dto.CreatePosition) (res responses.Position, err error)
		FindByID(ctx context.Context, id int) (res responses.Position, err error)
		Update(ctx context.Context, id int, dto dto.UpdatePosition) (res responses.Position, err error)
		Delete(ctx context.Context, id int) (err error)
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Position, meta commons.Pagination, err error)
	}
	position struct {
		repo repositories.Position
	}
)

func NewPositionService(repo repositories.Position) Position {
	return &position{
		repo: repo,
	}
}

func (s *position) Create(ctx context.Context, dto dto.CreatePosition) (res responses.Position, err error) {
	position := models.NewPositionFromCreatePosition(dto)

	position, err = s.repo.Create(ctx, position)
	if err != nil {
		return res, err
	}
	return responses.NewPositionResponse(position), nil
}

func (s *position) FindByID(ctx context.Context, id int) (res responses.Position, err error) {
	position, err := s.repo.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}
	return responses.NewPositionResponse(position), nil
}

func (s *position) Update(ctx context.Context, id int, dto dto.UpdatePosition) (res responses.Position, err error) {
	_, err = s.repo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	position := models.NewPositionFromUpdatePosition(dto)

	position, err = s.repo.Update(ctx, id, position)
	if err != nil {
		return res, err
	}
	return responses.NewPositionResponse(position), nil
}
func (s *position) Delete(ctx context.Context, id int) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *position) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Position, meta commons.Pagination, err error) {
	positions, totalData, err := s.repo.FindAll(ctx, base)
	if err != nil {
		return res, meta, err
	}

	for _, position := range positions {
		res = append(res, responses.NewPositionResponse(position))
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	return res, meta, nil
}
