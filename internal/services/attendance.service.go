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
	Attendance interface {
		FindByID(ctx context.Context, id int) (res responses.Attendance, err error)
		Create(ctx context.Context, dto dto.CreateAttendance) (res responses.Attendance, err error)
		Update(ctx context.Context, id int, dto dto.UpdateAttendance) (res responses.Attendance, err error)
		Delete(ctx context.Context, id int) (err error)
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Attendance, meta commons.Pagination, err error)
	}

	attendance struct {
		repo         repositories.Attendance
		employeeRepo repositories.Employee
	}
)

func NewAttendanceService(repo repositories.Attendance, employeeRepo repositories.Employee) Attendance {
	return &attendance{
		repo:         repo,
		employeeRepo: employeeRepo,
	}
}
func (s *attendance) FindByID(ctx context.Context, id int) (res responses.Attendance, err error) {
	attendance, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return res, err
	}

	res = responses.NewAttendanceResponse(attendance)

	return res, nil
}

func (s *attendance) Create(ctx context.Context, dto dto.CreateAttendance) (res responses.Attendance, err error) {
	_, err = s.employeeRepo.FindByID(ctx, dto.EmployeeID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	attendance, err := models.NewAttendanceFromCreateAttendance(dto)

	if err != nil {
		return res, err
	}

	attendance, err = s.repo.Create(ctx, attendance)
	if err != nil {
		return res, err
	}

	res = responses.NewAttendanceResponse(attendance)

	return res, nil
}

func (s *attendance) Update(ctx context.Context, id int, dto dto.UpdateAttendance) (res responses.Attendance, err error) {
	attendance, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return res, err
	}
	attendance, err = models.NewAttendanceFromUpdateAttendance(dto)
	if err != nil {
		return res, err
	}
	attendance, err = s.repo.Update(ctx, id, attendance)
	if err != nil {
		return res, err
	}
	res = responses.NewAttendanceResponse(attendance)
	return res, nil
}

func (s *attendance) Delete(ctx context.Context, id int) (err error) {
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
func (s *attendance) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Attendance, meta commons.Pagination, err error) {
	attendances, totalData, err := s.repo.FindAll(ctx, base)
	if err != nil {
		return res, meta, err
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	res = responses.NewAttendanceResponses(attendances)

	return res, meta, nil
}
