package services

import (
	"context"
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"github.com/handarudwiki/payroll-sistem/internal/responses"
	"gorm.io/gorm"
)

type (
	Payroll interface {
		Create(ctx context.Context, dto dto.CreatePayroll) (err error)
		FindByID(ctx context.Context, id int) (res responses.Payroll, err error)
	}

	payroll struct {
		repo         repositories.Payroll
		payslipRepo  repositories.PayslipDetail
		employeeRepo repositories.Employee
		db           *gorm.DB
	}
)

func NewPayrollService(repo repositories.Payroll, payslipRepo repositories.PayslipDetail, employeeRepo repositories.Employee, db *gorm.DB) Payroll {
	return &payroll{
		repo:         repo,
		payslipRepo:  payslipRepo,
		employeeRepo: employeeRepo,
		db:           db,
	}
}
func (s *payroll) Create(ctx context.Context, dto dto.CreatePayroll) (err error) {

	tx := s.db.Begin()

	if tx.Error != nil {
		return tx.Error
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
		if len(employee.Loans) > 0 {
			for _, loan := range employee.Loans {

				if loan.Status == "active" && loan.StartDate.Before(periodDate) {
					totalDeduction += loan.MonthlyInstallment
					loan.RemainingAmount -= loan.MonthlyInstallment

					if loan.RemainingAmount <= 0 {
						loan.Status = "paid"
					}
				}
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

		payroll := models.Payroll{
			EmployeeID:      employee.ID,
			Period:          periodDate,
			NetSalary:       totalSalary,
			TotalAllowances: totalAllowance,
			TotalDeductions: totalDeduction,
			Generated:       time.Now(),
		}

		payroll, err := s.repo.Create(ctx, tx, payroll)
		if err != nil {
			return err
		}

		var payslipDetails []models.PayslipDetail

		for _, comp := range employee.EmployeeComponent {
			payslipDetail := models.PayslipDetail{
				PayrollID:     payroll.ID,
				ComponentID:   comp.SalaryComponentID,
				Amount:        comp.Amount,
				ComponentType: comp.SalaryComponent.Type,
			}

			payslipDetails = append(payslipDetails, payslipDetail)
		}

		s.payslipRepo.BulkCreate(ctx, tx, payslipDetails)
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
