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
	Department interface {
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Department, meta commons.Pagination, err error)
		FindByID(ctx context.Context, id int) (responses.Department, error)
		Create(ctx context.Context, dto dto.CreateDepartment) (responses.Department, error)
		Update(ctx context.Context, id int, dto dto.UpdateDepartment) (responses.Department, error)
		Delete(ctx context.Context, id int) error
	}
	department struct {
		departmentRepository repositories.Department
	}
)

func NewDepartmentService(departmentRepository repositories.Department) Department {
	return &department{
		departmentRepository: departmentRepository,
	}
}

func (s *department) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Department, meta commons.Pagination, err error) {
	departments, totalData, err := s.departmentRepository.FindAll(ctx, base)
	if err != nil {
		return nil, commons.Pagination{}, err
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	res = responses.NewDepartments(departments)

	return res, meta, nil
}

func (s *department) FindByID(ctx context.Context, id int) (responses.Department, error) {
	department, err := s.departmentRepository.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return responses.Department{}, commons.ErrNotfound
	}

	if err != nil {
		return responses.Department{}, err
	}

	res := responses.NewDepartment(department)

	return res, nil
}

func (s *department) Create(ctx context.Context, dto dto.CreateDepartment) (responses.Department, error) {
	department := models.NewDepartmentFromCreateDepartment(dto)

	department, err := s.departmentRepository.Create(ctx, department)
	if err != nil {
		return responses.Department{}, err
	}

	res := responses.NewDepartment(department)

	return res, nil
}

func (s *department) Update(ctx context.Context, id int, dto dto.UpdateDepartment) (responses.Department, error) {
	department := models.NewDepartmentFromUpdateDepartment(dto)

	_, err := s.departmentRepository.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return responses.Department{}, commons.ErrNotfound
	}

	if err != nil {
		return responses.Department{}, err
	}

	department, err = s.departmentRepository.Update(ctx, id, department)
	if err != nil {
		return responses.Department{}, err
	}

	res := responses.NewDepartment(department)

	return res, nil
}
func (s *department) Delete(ctx context.Context, id int) error {
	_, err := s.departmentRepository.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return commons.ErrNotfound
	}

	if err != nil {
		return err
	}

	err = s.departmentRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
