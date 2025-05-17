package dto

type CreateEmployeeComponent struct {
	EmployeeID        int     `json:"employee_id" validate:"required"`
	SalaryComponentID int     `json:"salary_component_id" validate:"required"`
	Amount            float64 `json:"amount" validate:"required"`
	IsActive          bool    `json:"is_active" `
	CustomOverride    string  `json:"custom_override" validate:"required"`
}

type UpdateEmployeeComponent struct {
	EmployeeID        int     `json:"employee_id" validate:"required"`
	SalaryComponentID int     `json:"salary_component_id" validate:"required"`
	Amount            float64 `json:"amount" validate:"required"`
	IsActive          bool    `json:"is_active" `
	CustomOverride    string  `json:"custom_override" validate:"required"`
}
