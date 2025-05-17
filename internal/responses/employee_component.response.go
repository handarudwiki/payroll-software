package responses

import "github.com/handarudwiki/payroll-sistem/internal/models"

type EmployeeComponent struct {
	ID                int                     `json:"id"`
	EmployeeID        int                     `json:"employee_id"`
	SalaryComponentID int                     `json:"salary_component_id"`
	Amount            float64                 `json:"amount"`
	IsActive          bool                    `json:"is_active"`
	CustomOverride    string                  `json:"custom_override"`
	Employee          Employee                `json:"employee"`
	SalaryComponent   SalaryComponentResponse `json:"salary_component"`
}

func NewEmployeeComponentResponse(employeeComponent models.EmployeeComponent) EmployeeComponent {
	return EmployeeComponent{
		ID:                employeeComponent.ID,
		EmployeeID:        employeeComponent.EmployeeID,
		SalaryComponentID: employeeComponent.SalaryComponentID,
		Amount:            employeeComponent.Amount,
		IsActive:          employeeComponent.IsActive,
		CustomOverride:    employeeComponent.CustomOverride,
		Employee:          NewEmployeeResponse(employeeComponent.Employee),
		SalaryComponent:   NewSalaryComponentResponse(employeeComponent.SalaryComponent),
	}
}

func NewEmployeeComponentResponses(employeeComponents []models.EmployeeComponent) []EmployeeComponent {
	responses := make([]EmployeeComponent, len(employeeComponents))
	for i, employeeComponent := range employeeComponents {
		responses[i] = NewEmployeeComponentResponse(employeeComponent)
	}
	return responses
}
