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
	PayslipDetail   []PayslipDetail `json:"payslip_details" gorm:"foreignKey:PayrollID;references:ID"`
	Generated       time.Time       `json:"generated" gorm:"default:false"`
	CreatedAt       time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}
