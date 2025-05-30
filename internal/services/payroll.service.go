package services

import (
	"context"
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/models/commons"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"github.com/handarudwiki/payroll-sistem/internal/responses"
	"gorm.io/gorm"
)

type (
	Payroll interface {
		Create(ctx context.Context, dto dto.CreatePayroll) (err error)
		FindByID(ctx context.Context, id int) (res responses.Payroll, err error)
		FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Payroll, meta commons.Pagination, err error)
	}

	payroll struct {
		repo         repositories.Payroll
		payslipRepo  repositories.PayslipDetail
		employeeRepo repositories.Employee
		leaveRepo    repositories.Leave
		loanRepo     repositories.Loan
		db           *gorm.DB
	}
)

func NewPayrollService(repo repositories.Payroll, payslipRepo repositories.PayslipDetail, employeeRepo repositories.Employee, leaveRepo repositories.Leave, loanRepo repositories.Loan, db *gorm.DB) Payroll {
	return &payroll{
		repo:         repo,
		payslipRepo:  payslipRepo,
		employeeRepo: employeeRepo,
		loanRepo:     loanRepo,
		leaveRepo:    leaveRepo,
		db:           db,
	}
}
func (s *payroll) Create(ctx context.Context, dto dto.CreatePayroll) (err error) {

	tx := s.db.Begin()

	if tx.Error != nil {
		err = tx.Error
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = r.(error)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
			if err != nil {
				tx.Rollback()
			}
		}
	}()

	periodDate, err := time.Parse("2006-01-02", dto.Period)

	if err != nil {
		return err
	}

	var employees []models.Employee

	if dto.IsAll {
		employees, err = s.employeeRepo.FindAllActive(ctx)
		if err != nil {
			return err
		}

	} else {
		employees, err = s.employeeRepo.FindByIDSActive(ctx, dto.EmployeIDS)
		if err != nil {
			return err
		}
	}

	for _, employee := range employees {
		baseSalary := employee.Position.BaseSalary

		totalAllowance := float64(0)
		totalDeduction := float64(0)

		//handle component salary
		for _, comp := range employee.EmployeeComponent {
			if comp.SalaryComponent.Type == "allowance" {
				totalAllowance += comp.Amount
			} else if comp.SalaryComponent.Type == "deduction" {
				totalDeduction += comp.Amount
			}
		}

		//handle loan
		for _, loan := range employee.Loans {

			if loan.Status == "active" && loan.StartDate.Before(periodDate) {
				totalDeduction += loan.MonthlyInstallment
				loan.RemainingAmount -= loan.MonthlyInstallment

				if loan.RemainingAmount <= 0 {
					loan.Status = "paid"
				}
			}
			_, err = s.loanRepo.Update(ctx, loan.ID, loan)
			if err != nil {
				return err
			}
		}

		absentDays := 0
		lateDays := 0

		//handle attendance
		if len(employee.Attendances) > 0 {
			for _, attendance := range employee.Attendances {
				if attendance.Date.Month() == periodDate.Month() && attendance.Date.Year() == periodDate.Year() {
					if attendance.Status == models.Absent {
						absentDays++
					}
					if attendance.Status == models.Late {
						lateDays++
					}
				}
			}
		}

		absentPenalty := float64(absentDays * models.AbsentPenalty)
		latePenalty := float64(lateDays * models.LatePenalty)

		totalDeduction += absentPenalty + latePenalty
		totalSalary := baseSalary + totalAllowance - totalDeduction

		leaves, err := s.leaveRepo.FindByEmployeeIDMaternity(ctx, employee.ID)

		if err != nil {
			return err
		}

		if len(leaves) > 0 {
			totalSalary = 0
			totalAllowance = 0
			totalDeduction = 0
		}

		payroll := models.Payroll{
			EmployeeID:      employee.ID,
			Period:          periodDate,
			NetSalary:       totalSalary,
			TotalAllowances: totalAllowance,
			BaseSalary:      baseSalary,
			TotalDeductions: totalDeduction,
			GeneratedAt:     time.Now(),
		}

		payroll, err = s.repo.Create(ctx, tx, payroll)
		if err != nil {
			return err
		}

		var payslipDetails []models.PayslipDetail

		if len(leaves) == 0 {
			for _, comp := range employee.EmployeeComponent {
				payslipDetail := models.PayslipDetail{
					PayrollID:         payroll.ID,
					SalaryComponentID: comp.SalaryComponentID,
					Amount:            comp.Amount,
					ComponentType:     comp.SalaryComponent.Type,
				}

				payslipDetails = append(payslipDetails, payslipDetail)
			}

			err = s.payslipRepo.BulkCreate(ctx, tx, payslipDetails)
			if err != nil {
				return err
			}
		}
	}

	return
}

func (s *payroll) FindByID(ctx context.Context, id int) (res responses.Payroll, err error) {
	payroll, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return
	}

	res = responses.NewPayroll(payroll)

	return
}

func (s *payroll) FindAll(ctx context.Context, base dto.BaseQuery) (res []responses.Payroll, meta commons.Pagination, err error) {
	payrolls, totalData, err := s.repo.FindAll(ctx, base)
	if err != nil {
		return res, meta, err
	}

	for _, payroll := range payrolls {
		res = append(res, responses.NewPayroll(payroll))
	}

	meta = commons.NewPagination(base.Page, base.Limit, int(totalData))

	res = responses.NewPayrolls(payrolls)

	return res, meta, nil
}
