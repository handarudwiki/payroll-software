package seeds

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

func PositionSeed(repo repositories.Position) error {
	positions := []models.Position{
		{
			Name:       "Software Engineer",
			BaseSalary: 5000000,
		},
		{
			Name:       "Product Manager",
			BaseSalary: 7000000,
		},
		{
			Name:       "Data Scientist",
			BaseSalary: 6000000,
		},
		{
			Name:       "UI/UX Designer",
			BaseSalary: 4500000,
		},
		{
			Name:       "DevOps Engineer",
			BaseSalary: 5500000,
		},
		{
			Name:       "QA Engineer",
			BaseSalary: 4000000,
		},
		{
			Name:       "Backend Developer",
			BaseSalary: 5200000,
		},
		{
			Name:       "Frontend Developer",
			BaseSalary: 5100000,
		},
		{
			Name:       "Mobile Developer",
			BaseSalary: 5300000,
		},
		{
			Name:       "System Analyst",
			BaseSalary: 4800000,
		},
		{
			Name:       "IT Support Specialist",
			BaseSalary: 3500000,
		},
		{
			Name:       "Scrum Master",
			BaseSalary: 6000000,
		},
		{
			Name:       "Project Manager",
			BaseSalary: 6800000,
		},
		{
			Name:       "Cybersecurity Analyst",
			BaseSalary: 6200000,
		},
		{
			Name:       "Database Administrator",
			BaseSalary: 5400000,
		},
	}

	_, err := repo.BulkCreate(context.Background(), positions)
	if err != nil {
		return err
	}

	return nil
}
