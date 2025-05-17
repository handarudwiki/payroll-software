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
	EmployeeComponent interface {
		Create(ctx context.Context, dto dto.CreateEmployeeComponent) (res responses.EmployeeComponent, err error)
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.EmployeeComponent, meta commons.Pagination, err error)
		FindByID(ctx context.Context, id int) (res responses.EmployeeComponent, err error)
		Update(ctx context.Context, id int, dto dto.UpdateEmployeeComponent) (res responses.EmployeeComponent, err error)
		Delete(ctx context.Context, id int) (err error)
	}

	employeeComponent struct {
		repo                repositories.EmployeeComponent
		employeeRepo        repositories.Employee
		salaryComponentRepo repositories.SalaryComponent
	}
)

func NewEmployeeComponentService(repo repositories.EmployeeComponent, employeeRepo repositories.Employee, salaryComponentRepo repositories.SalaryComponent) EmployeeComponent {
	return &employeeComponent{
		repo:                repo,
		employeeRepo:        employeeRepo,
		salaryComponentRepo: salaryComponentRepo,
	}
}

func (s *employeeComponent) Create(ctx context.Context, dto dto.CreateEmployeeComponent) (res responses.EmployeeComponent, err error) {
	_, err = s.employeeRepo.FindByID(ctx, dto.EmployeeID)
	if err != nil {
		return res, err
	}

	_, err = s.salaryComponentRepo.FindByID(ctx, dto.SalaryComponentID)
	if err != nil {
		return res, err
	}

	employeeComponent := models.NewEmployeeComponentFromCreateEmployeeComponent(dto)

	employeeComponent, err = s.repo.Create(ctx, employeeComponent)
	if err != nil {
		return res, err
	}

	res = responses.NewEmployeeComponentResponse(employeeComponent)
	return res, nil
}

func (s *employeeComponent) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.EmployeeComponent, meta commons.Pagination, err error) {
	employeeComponents, totalData, err := s.repo.FindAll(ctx, base)
	if err != nil {
		return res, meta, err
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	res = responses.NewEmployeeComponentResponses(employeeComponents)
	return res, meta, nil
}

func (s *employeeComponent) FindByID(ctx context.Context, id int) (res responses.EmployeeComponent, err error) {
	employeeComponent, err := s.repo.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	res = responses.NewEmployeeComponentResponse(employeeComponent)
	return res, nil
}

func (s *employeeComponent) Update(ctx context.Context, id int, dto dto.UpdateEmployeeComponent) (res responses.EmployeeComponent, err error) {
	_, err = s.repo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	_, err = s.salaryComponentRepo.FindByID(ctx, dto.SalaryComponentID)
	if err != nil {
		return res, err
	}

	employeeComponent := models.NewEmployeeComponentFromUpdateEmployeeComponent(dto)

	employeeComponent, err = s.repo.Update(ctx, id, employeeComponent)
	if err != nil {
		return res, err
	}

	res = responses.NewEmployeeComponentResponse(employeeComponent)
	return res, nil
}

func (s *employeeComponent) Delete(ctx context.Context, id int) (err error) {
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
