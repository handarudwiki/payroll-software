package seeds

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

func DepartmentSeed(repo repositories.Department) (err error) {
	departments := []models.Department{
		{
			Name:        "IT",
			Description: "Information Technology Department",
		},
		{
			Name:        "HR",
			Description: "Human Resources Department",
		},
		{
			Name:        "Finance",
			Description: "Finance and Accounting Department",
		},
		{
			Name:        "Marketing",
			Description: "Marketing and Communications Department",
		},
		{
			Name:        "Sales",
			Description: "Sales and Business Development Department",
		},
		{
			Name:        "Operations",
			Description: "Operations and Logistics Department",
		},
		{
			Name:        "Customer Service",
			Description: "Customer Support and Service Department",
		},
		{
			Name:        "Legal",
			Description: "Legal and Compliance Department",
		},
		{
			Name:        "Procurement",
			Description: "Procurement and Vendor Management Department",
		},
		{
			Name:        "R&D",
			Description: "Research and Development Department",
		},
		{
			Name:        "Administration",
			Description: "Administrative Support Department",
		},
	}

	_, err = repo.BulkCreate(context.Background(), departments)
	if err != nil {
		return err
	}
	return
}
