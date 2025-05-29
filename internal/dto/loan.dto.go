package dto

type CreateLoan struct {
	EmployeeID         int     `json:"employee_id" validate:"required"`
	TotalAmount        float64 `json:"total_amount" validate:"required"`
	MonthlyInstallment float64 `json:"monthly_installment" validate:"required"`
	StartDate          string  `json:"start_date" validate:"required,datetime=2006-01-02"`
	Status             string  `json:"status" validate:"required,oneof=active cancelled paid"`
}

type UpdateLoan struct {
	EmployeeID         int     `json:"employee_id" validate:"required"`
	TotalAmount        float64 `json:"total_amount" validate:"required"`
	MonthlyInstallment float64 `json:"monthly_installment" validate:"required"`
	RemainingAmount    float64 `json:"remaining_amount" validate:"required"`
	StartDate          string  `json:"start_date" validate:"required,datetime=2006-01-02"`
	Status             string  `json:"status" validate:"required,oneof=active cancelled paid"`
}
