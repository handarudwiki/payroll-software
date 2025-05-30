package seeds

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

func EmployeeComponentSeed(employeeComponentRepo repositories.EmployeeComponent, salaryComponentRepo repositories.SalaryComponent, employeeRepo repositories.Employee) (err error) {
	salaryComponents, err := salaryComponentRepo.FindAllOnly(context.Background())
	if err != nil {
		return err
	}
	employees, err := employeeRepo.FindAllActive(context.Background())
	if err != nil {
		return err
	}

	employeeComponents := []models.EmployeeComponent{
		{
			EmployeeID:        employees[0].ID,
			SalaryComponentID: salaryComponents[0].ID,
			Amount:            1000000,
			CustomOverride:    "Tunjangan Khusus A",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[1].ID,
			SalaryComponentID: salaryComponents[1].ID,
			Amount:            1500000,
			CustomOverride:    "Tunjangan Khusus B",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[2].ID,
			SalaryComponentID: salaryComponents[2].ID,
			Amount:            1200000,
			CustomOverride:    "Bonus Produktivitas",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[3].ID,
			SalaryComponentID: salaryComponents[3].ID,
			Amount:            1800000,
			CustomOverride:    "Insentif Lembur",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[4].ID,
			SalaryComponentID: salaryComponents[4].ID,
			Amount:            1300000,
			CustomOverride:    "Tunjangan Transportasi",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[5].ID,
			SalaryComponentID: salaryComponents[5].ID,
			Amount:            1100000,
			CustomOverride:    "Tunjangan Komunikasi",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[6].ID,
			SalaryComponentID: salaryComponents[6].ID,
			Amount:            1700000,
			CustomOverride:    "Bonus Tahunan",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[2].ID,
			SalaryComponentID: salaryComponents[2].ID,
			Amount:            1250000,
			CustomOverride:    "Tunjangan Makan",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[2].ID,
			SalaryComponentID: salaryComponents[2].ID,
			Amount:            1400000,
			CustomOverride:    "Tunjangan Skill",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[2].ID,
			SalaryComponentID: salaryComponents[2].ID,
			Amount:            1600000,
			CustomOverride:    "Tunjangan Kinerja",
			IsActive:          true,
		},
		{
			EmployeeID:        employees[2].ID,
			SalaryComponentID: salaryComponents[2].ID,
			Amount:            1550000,
			CustomOverride:    "Tunjangan Proyek",
			IsActive:          true,
		},
	}

	_, err = employeeComponentRepo.BulkCreate(context.Background(), employeeComponents)
	if err != nil {
		return err
	}

	return nil

}
