package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type StatusEmployee string

const (
	STATUS_ACTIVE     StatusEmployee = "active"
	STATUS_INACTIVE   StatusEmployee = "inactive"
	STATUS_TERMINATED StatusEmployee = "terminated"
)

type Employee struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	Email             string              `json:"email"`
	Phone             string              `json:"phone"`
	NIK               string              `json:"nik"`
	PositionID        int                 `json:"position_id"`
	DepartmentID      int                 `json:"department_id"`
	HireDate          string              `json:"hire_date"`
	UserID            int                 `json:"user_id"`
	Status            StatusEmployee      `json:"status"`
	Loans             []Loan              `json:"loans,omitempty"`
	Attendances       []Attendance        `json:"attendances,omitempty"`
	EmployeeComponent []EmployeeComponent `json:"employee_components,omitempty"`
	Leaves            []Leave             `json:"leaves,omitempty"`
	Department        Department          `json:"department"`
	Position          Position            `json:"position"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
}

func NewEmployeeFromCreateEmployee(data dto.CreateEmployee) Employee {
	return Employee{
		Name:         data.Name,
		Email:        data.Email,
		Phone:        data.Phone,
		NIK:          data.NIK,
		PositionID:   data.PositionID,
		DepartmentID: data.DepartmentID,
		HireDate:     data.HireDate,
		Status:       StatusEmployee(data.Status),
	}
}

func NewEmployeeFromUpdateEmployee(data dto.UpdateEmployee) Employee {
	return Employee{
		Name:         data.Name,
		Email:        data.Email,
		Phone:        data.Phone,
		NIK:          data.NIK,
		PositionID:   data.PositionID,
		DepartmentID: data.DepartmentID,
		HireDate:     data.HireDate,
		Status:       StatusEmployee(data.Status),
	}
}
