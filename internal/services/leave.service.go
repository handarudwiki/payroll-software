package services

import (
	"context"
	"errors"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/models/commons"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"github.com/handarudwiki/payroll-sistem/internal/responses"
)

type (
	Leave interface {
		Create(ctx context.Context, dto dto.CreateLeave) (res responses.Leave, err error)
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Leave, meta commons.Pagination, err error)
		FindByID(ctx context.Context, id int) (res responses.Leave, err error)
		Update(ctx context.Context, id int, dto dto.UpdateLeave) (res responses.Leave, err error)
		Delete(ctx context.Context, id int) (err error)
	}

	leave struct {
		repo         repositories.Leave
		employeeRepo repositories.Employee
	}
)

func NewLeaveService(repo repositories.Leave, employeeRepo repositories.Employee) Leave {
	return &leave{
		repo:         repo,
		employeeRepo: employeeRepo,
	}
}

func (s *leave) Create(ctx context.Context, dto dto.CreateLeave) (res responses.Leave, err error) {
	_, err = s.employeeRepo.FindByID(ctx, dto.EmployeeID)
	if err != nil {
		return res, err
	}

	leave, err := models.NewLeaveFromCreateLeave(dto)

	if err != nil {
		return res, err
	}

	leave, err = s.repo.Create(ctx, leave)
	if err != nil {
		return res, err
	}

	res = responses.NewLeaveResponse(leave)
	return res, nil
}

func (s *leave) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Leave, meta commons.Pagination, err error) {
	leaves, totalData, err := s.repo.FindAll(ctx, base)
	if err != nil {
		return res, meta, err
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	res = responses.NewLeaveResponses(leaves)
	return res, meta, nil
}
func (s *leave) FindByID(ctx context.Context, id int) (res responses.Leave, err error) {

	leave, err := s.repo.FindByID(ctx, id)

	if errors.Is(err, commons.ErrNotfound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	res = responses.NewLeaveResponse(leave)
	return res, nil
}

func (s *leave) Update(ctx context.Context, id int, dto dto.UpdateLeave) (res responses.Leave, err error) {
	_, err = s.repo.FindByID(ctx, id)

	if errors.Is(err, commons.ErrNotfound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	leave, err := models.NewLeaveFromUpdateLeave(dto)

	if err != nil {
		return res, err
	}

	leave, err = s.repo.Update(ctx, id, leave)
	if err != nil {
		return res, err
	}

	res = responses.NewLeaveResponse(leave)
	return res, nil
}
func (s *leave) Delete(ctx context.Context, id int) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
