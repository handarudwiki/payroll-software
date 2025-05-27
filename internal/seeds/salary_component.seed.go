package seeds

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

func SalaryComponentSeed(salaryComponentRepo repositories.SalaryComponent) (err error) {
	salaryComponents := []models.SalaryComponent{
		{
			Name:        "Gaji Pokok",
			IsRecurring: true,
			Type:        models.SalaryTypeAllowance,
		},
	}

	_, err = salaryComponentRepo.BulkCreate(context.Background(), salaryComponents)
	if err != nil {
		return err
	}

	return nil
}
