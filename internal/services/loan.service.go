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
	Loan interface {
		Create(ctx context.Context, dto dto.CreateLoan) (res responses.Loan, err error)
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Loan, meta commons.Pagination, err error)
		FindByID(ctx context.Context, id int) (res responses.Loan, err error)
		Update(ctx context.Context, id int, dto dto.UpdateLoan) (res responses.Loan, err error)
		Delete(ctx context.Context, id int) (err error)
	}

	loan struct {
		repo         repositories.Loan
		employeeRepo repositories.Employee
	}
)

func NewLoanService(repo repositories.Loan, employeeRepo repositories.Employee) Loan {
	return &loan{
		repo:         repo,
		employeeRepo: employeeRepo,
	}
}

func (s *loan) Create(ctx context.Context, dto dto.CreateLoan) (res responses.Loan, err error) {
	_, err = s.employeeRepo.FindByID(ctx, dto.EmployeeID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	loan, err := models.NewLoanFromCreateLoan(dto)
	if err != nil {
		return res, err
	}

	loan, err = s.repo.Create(ctx, loan)
	if err != nil {
		return res, err
	}

	res = responses.NewLoanResponse(loan)
	return res, nil
}
func (s *loan) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Loan, meta commons.Pagination, err error) {
	loans, totalData, err := s.repo.FindAll(ctx, base)
	if err != nil {
		return res, meta, err
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	res = responses.NewLoanResponses(loans)
	return res, meta, nil
}

func (s *loan) FindByID(ctx context.Context, id int) (res responses.Loan, err error) {
	loan, err := s.repo.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	res = responses.NewLoanResponse(loan)
	return res, nil
}
func (s *loan) Update(ctx context.Context, id int, dto dto.UpdateLoan) (res responses.Loan, err error) {
	_, err = s.repo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, commons.ErrNotfound
	}

	if err != nil {
		return res, err
	}

	loan, err := models.NewLoanFromUpdateLoan(dto)

	if err != nil {
		return res, err
	}

	loan, err = s.repo.Update(ctx, id, loan)
	if err != nil {
		return res, err
	}

	res = responses.NewLoanResponse(loan)
	return res, nil
}
func (s *loan) Delete(ctx context.Context, id int) (err error) {
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
