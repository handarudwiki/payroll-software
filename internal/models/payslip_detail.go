package models

import "time"

type PayslipDetail struct {
	ID                int                 `json:"id" `
	PayrollID         int                 `json:"payroll_id" `
	SalaryComponentID int                 `json:"salary_component_id" `
	ComponentType     SalaryComponentType `json:"component_type" `
	Amount            float64             `json:"amount" `
	CreatedAt         time.Time           `json:"created_at" `
	UpdatedAt         time.Time           `json:"updated_at"`
	Payroll           Payroll             `json:"payroll" `
	SalaryComponent   SalaryComponent     `json:"salary_component" `
}
