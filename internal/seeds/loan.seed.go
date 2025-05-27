package seeds

import (
	"context"
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

func LoanSeed(loanRepo repositories.Loan, employeeRepo repositories.Employee) (err error) {
	employees, err := employeeRepo.FindAllActive(context.Background())
	if err != nil {
		return err
	}

	loans := []models.Loan{
		{
			EmployeeID:         employees[0].ID,
			TotalAmount:        1000000,
			MonthlyInstallment: 200000,
			RemainingAmount:    800000,
			StartDate:          time.Now(),
			Status:             models.LoanStatusActive,
		},
	}

	_, err = loanRepo.BulkCreate(context.Background(), loans)
	if err != nil {
		return err
	}

	return nil
}
