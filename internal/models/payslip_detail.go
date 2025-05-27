package models

import "time"

type PayslipDetail struct {
	ID            int                 `json:"id" gorm:"primaryKey"`
	PayrollID     int                 `json:"payroll_id" gorm:"not null"`
	ComponentID   int                 `json:"component_id" gorm:"not null"`
	ComponentType SalaryComponentType `json:"component_type" gorm:"not null"`
	Amount        float64             `json:"amount" gorm:"not null"`
	CreatedAt     time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time           `json:"updated_at" gorm:"autoUpdateTime"`
	Payroll       Payroll             `json:"payroll" gorm:"foreignKey:PayrollID;references:ID"`
	Component     SalaryComponent     `json:"component" gorm:"foreignKey:ComponentID;references:ID"`
}
