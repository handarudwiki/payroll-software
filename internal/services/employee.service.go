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
	Employee interface {
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Employee, meta commons.Pagination, err error)
		FindByID(ctx context.Context, id int) (res responses.Employee, err error)
		Create(ctx context.Context, dto dto.CreateEmployee) (res responses.Employee, err error)
		Update(ctx context.Context, id int, dto dto.UpdateEmployee) (res responses.Employee, err error)
		Delete(ctx context.Context, id int) (err error)
	}

	employee struct {
		repo           repositories.Employee
		departmentRepo repositories.Department
		positionRepo   repositories.Position
		userRepo       repositories.User
	}
)

func NewEmployeeService(repo repositories.Employee, departmentRepo repositories.Department, positionRepo repositories.Position, userRepo repositories.User) Employee {
	return &employee{
		repo:           repo,
		departmentRepo: departmentRepo,
		positionRepo:   positionRepo,
		userRepo:       userRepo,
	}
}

func (s *employee) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Employee, meta commons.Pagination, err error) {
	employees, totalData, err := s.repo.FindAll(ctx, base)
	if err != nil {
		return res, meta, err
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	res = responses.NewEmployeeListResponse(employees)
	return res, meta, nil
}

func (s *employee) FindByID(ctx context.Context, id int) (res responses.Employee, err error) {
	employee, err := s.repo.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	res = responses.NewEmployeeResponse(employee)
	return res, nil
}
func (s *employee) Create(ctx context.Context, dto dto.CreateEmployee) (res responses.Employee, err error) {
	employee := models.NewEmployeeFromCreateEmployee(dto)

	department, err := s.departmentRepo.FindByID(ctx, dto.DepartmentID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}
	employee.Department = department

	position, err := s.positionRepo.FindByID(ctx, dto.PositionID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}
	employee.Position = position

	nik, err := s.repo.FindByNIK(ctx, dto.NIK)

	if nik.ID != 0 {
		return res, commons.ErrConflict
	}

	if err == nil {
		return res, err
	}

	email, err := s.repo.FindByEmail(ctx, dto.Email)
	if email.ID != 0 {
		return res, commons.ErrConflict
	}
	if err == nil {
		return res, err
	}
	phone, err := s.repo.FindByPhone(ctx, dto.Phone)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if phone.ID != 0 {
		return res, commons.ErrConflict
	}

	user := models.User{
		Name:     dto.Name,
		Username: dto.Email,
		Password: dto.Password,
		Role:     models.RoleUser,
	}

	user.EncryptPassword()

	user, err = s.userRepo.Create(ctx, user)
	if err != nil {
		return res, err
	}

	employee.UserID = user.ID

	employee, err = s.repo.Create(ctx, employee)
	if err != nil {
		return res, err
	}

	res = responses.NewEmployeeResponse(employee)
	return res, nil
}
func (s *employee) Update(ctx context.Context, id int, dto dto.UpdateEmployee) (res responses.Employee, err error) {
	_, err = s.repo.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}
	employee := models.NewEmployeeFromUpdateEmployee(dto)
	_, err = s.departmentRepo.FindByID(ctx, dto.DepartmentID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	_, err = s.positionRepo.FindByID(ctx, dto.PositionID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	nik, err := s.repo.FindByNIK(ctx, dto.NIK)
	if nik.ID != 0 && nik.ID != id {
		return res, commons.ErrConflict
	}
	if err == nil {
		return res, err
	}
	email, err := s.repo.FindByEmail(ctx, dto.Email)
	if email.ID != 0 && email.ID != id {
		return res, commons.ErrConflict
	}
	if err == nil {
		return res, err
	}
	phone, err := s.repo.FindByPhone(ctx, dto.Phone)
	if phone.ID != 0 && phone.ID != id {
		return res, commons.ErrConflict
	}
	if err == nil {
		return res, err
	}

	employee, err = s.repo.Update(ctx, id, employee)
	if err != nil {
		return res, err
	}

	res = responses.NewEmployeeResponse(employee)
	return res, nil

}

func (s *employee) Delete(ctx context.Context, id int) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
