package models

import "time"

type Payroll struct {
	ID              int             `json:"id"`
	EmployeeID      int             `json:"employee_id"`
	Employee        Employee        `json:"employee" gorm:"foreignKey:EmployeeID;references:ID"`
	Period          time.Time       `json:"period"`
	TotalAllowances float64         `json:"total_allowance"`
	TotalDeductions float64         `json:"total_deduction"`
	NetSalary       float64         `json:"net_salary"`
	BaseSalary      float64         `json:"base_salary" gorm:"not null"`
	PayslipDetail   []PayslipDetail `json:"payslip_details" gorm:"foreignKey:PayrollID;references:ID"`
	GeneratedAt     time.Time       `json:"generate_at" gorm:"default:false"`
	CreatedAt       time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}
