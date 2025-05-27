package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type SalaryComponentType string

const (
	SalaryTypeAllowance SalaryComponentType = "allowance"
	SalaryTypeDeduction SalaryComponentType = "deduction"
)

type SalaryComponent struct {
	ID          int                 `json:"id" gorm:"primaryKey"`
	Name        string              `json:"name" gorm:"not null"`
	Type        SalaryComponentType `json:"type" gorm:"not null"`
	IsRecurring bool                `json:"is_recurring" gorm:"not null"`
	CreatedAt   time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time           `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewSalaryCmponentFromCreateSalaryComponent(dto dto.CreateSalaryComponent) SalaryComponent {
	return SalaryComponent{
		Name:        dto.Name,
		Type:        SalaryComponentType(dto.Type),
		IsRecurring: dto.IsRecurring,
	}
}

func NewSalaryComponentFromUpdateSalaryComponent(dto dto.UpdateSalaryComponent) SalaryComponent {
	return SalaryComponent{
		Name:        dto.Name,
		Type:        SalaryComponentType(dto.Type),
		IsRecurring: dto.IsRecurring,
	}
}
