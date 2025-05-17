package dto

type CreateSalaryComponent struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required,oneof=allowance deduction"`
	IsRecurring bool   `json:"is_recurring"`
}
type UpdateSalaryComponent struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required,oneof=allowance deduction"`
	IsRecurring bool   `json:"is_recurring"`
}
