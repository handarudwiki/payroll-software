package responses

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/models"
)

type Loan struct {
	ID                 int       `json:"id"`
	EmployeeID         int       `json:"employee_id"`
	TotalAmount        float64   `json:"total_amount"`
	MonthlyInstallment float64   `json:"monthly_installment"`
	RemainingAmount    float64   `json:"remaining_amount"`
	StartDate          time.Time `json:"start_date"`
	Status             string    `json:"status"`
	Employee           Employee  `json:"employee"`
}

func NewLoanResponse(loan models.Loan) Loan {
	return Loan{
		ID:                 loan.ID,
		EmployeeID:         loan.EmployeeID,
		TotalAmount:        loan.TotalAmount,
		MonthlyInstallment: loan.MonthlyInstallment,
		RemainingAmount:    loan.RemainingAmount,
		StartDate:          loan.StartDate,
		Status:             string(loan.Status),
		Employee:           NewEmployeeResponse(loan.Employee),
	}
}

func NewLoanResponses(loans []models.Loan) []Loan {
	responses := make([]Loan, len(loans))
	for i, loan := range loans {
		responses[i] = NewLoanResponse(loan)
	}
	return responses
}
