package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type Department struct {
	ID          int       `json:"id" `
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewDepartmentFromCreateDepartment(data dto.CreateDepartment) Department {
	return Department{
		Name:        data.Name,
		Description: data.Description,
	}
}

func NewDepartmentFromUpdateDepartment(data dto.UpdateDepartment) Department {
	return Department{
		Name:        data.Name,
		Description: data.Description,
	}
}
