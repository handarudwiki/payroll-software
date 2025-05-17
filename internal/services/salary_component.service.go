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
	SalaryComponent interface {
		Create(ctx context.Context, dto dto.CreateSalaryComponent) (res responses.SalaryComponentResponse, err error)
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.SalaryComponentResponse, meta commons.Pagination, err error)
		FindByID(ctx context.Context, id int) (res responses.SalaryComponentResponse, err error)
		Update(ctx context.Context, id int, dto dto.UpdateSalaryComponent) (res responses.SalaryComponentResponse, err error)
		Delete(ctx context.Context, id int) (err error)
	}

	salaryComponent struct {
		repo repositories.SalaryComponent
	}
)

func NewSalaryComponentService(repo repositories.SalaryComponent) SalaryComponent {
	return &salaryComponent{
		repo: repo,
	}
}

func (s *salaryComponent) Create(ctx context.Context, dto dto.CreateSalaryComponent) (res responses.SalaryComponentResponse, err error) {
	salaryComponet := models.NewSalaryCmponentFromCreateSalaryComponent(dto)

	salaryComponent, err := s.repo.Create(ctx, salaryComponet)
	if err != nil {
		return res, err
	}

	res = responses.NewSalaryComponentResponse(salaryComponent)
	return res, nil
}

func (s *salaryComponent) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.SalaryComponentResponse, meta commons.Pagination, err error) {
	salaryComponents, totalData, err := s.repo.FindAll(ctx, base)
	if err != nil {
		return res, meta, err
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	res = responses.NewSalaryComponentResponses(salaryComponents)
	return res, meta, nil
}

func (s *salaryComponent) FindByID(ctx context.Context, id int) (res responses.SalaryComponentResponse, err error) {
	salaryComponent, err := s.repo.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	res = responses.NewSalaryComponentResponse(salaryComponent)
	return res, nil
}

func (s *salaryComponent) Update(ctx context.Context, id int, dto dto.UpdateSalaryComponent) (res responses.SalaryComponentResponse, err error) {

	salaryComponent, err := s.repo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	salaryComponent = models.NewSalaryComponentFromUpdateSalaryComponent(dto)

	salaryComponent, err = s.repo.Update(ctx, id, salaryComponent)
	if err != nil {
		return res, err
	}

	res = responses.NewSalaryComponentResponse(salaryComponent)
	return res, nil
}

func (s *salaryComponent) Delete(ctx context.Context, id int) (err error) {
	_, err = s.repo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return commons.ErrNotfound
	}

	if err != nil {
		return err
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
