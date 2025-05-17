package responses

import "github.com/handarudwiki/payroll-sistem/internal/models"

type Department struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewDepartment(department models.Department) Department {
	return Department{
		ID:          department.ID,
		Name:        department.Name,
		Description: department.Description,
	}
}

func NewDepartments(departments []models.Department) []Department {
	var res []Department
	for _, department := range departments {
		res = append(res, NewDepartment(department))
	}
	return res
}
