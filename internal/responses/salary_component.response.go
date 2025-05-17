package responses

import "github.com/handarudwiki/payroll-sistem/internal/models"

type SalaryComponentResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	IsRecurring bool   `json:"is_recurring"`
}

func NewSalaryComponentResponse(salaryComponent models.SalaryComponent) SalaryComponentResponse {
	return SalaryComponentResponse{
		ID:          salaryComponent.ID,
		Name:        salaryComponent.Name,
		Type:        salaryComponent.Name,
		IsRecurring: salaryComponent.IsRecurring,
	}
}

func NewSalaryComponentResponses(salaryComponents []models.SalaryComponent) []SalaryComponentResponse {
	responses := make([]SalaryComponentResponse, len(salaryComponents))
	for i, salaryComponent := range salaryComponents {
		responses[i] = NewSalaryComponentResponse(salaryComponent)
	}
	return responses
}
