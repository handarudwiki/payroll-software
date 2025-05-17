package dto

type CreatePosition struct {
	Name       string  `json:"name" validate:"required"`
	BaseSalary float64 `json:"base_salary" validate:"required,min=1"`
}

type UpdatePosition struct {
	Name       string  `json:"name" validate:"required"`
	BaseSalary float64 `json:"base_salary" validate:"required,min=1"`
}
