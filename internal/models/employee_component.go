package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type EmployeeComponent struct {
	ID                int              `json:"id" gorm:"primaryKey"`
	EmployeeID        int              `json:"employee_id" gorm:"not null"`
	SalaryComponentID int              `json:"salary_component_id" gorm:"not null"`
	Amount            float64          `json:"amount" gorm:"not null"`
	IsActive          bool             `json:"is_active" gorm:"not null"`
	CustomOverride    string           `json:"custom_override" gorm:"not null"`
	Employee          *Employee        `json:"employee" gorm:"foreignKey:EmployeeID;references:ID"`
	SalaryComponent   *SalaryComponent `json:"salary_component" gorm:"foreignKey:SalaryComponentID;references:ID"`
	CreatedAt         time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewEmployeeComponentFromCreateEmployeeComponent(data dto.CreateEmployeeComponent) EmployeeComponent {
	return EmployeeComponent{
		EmployeeID:        data.EmployeeID,
		SalaryComponentID: data.SalaryComponentID,
		Amount:            data.Amount,
		IsActive:          data.IsActive,
		CustomOverride:    data.CustomOverride,
	}
}

func NewEmployeeComponentFromUpdateEmployeeComponent(data dto.UpdateEmployeeComponent) EmployeeComponent {
	return EmployeeComponent{
		EmployeeID:        data.EmployeeID,
		SalaryComponentID: data.SalaryComponentID,
		Amount:            data.Amount,
		IsActive:          data.IsActive,
		CustomOverride:    data.CustomOverride,
	}
}
