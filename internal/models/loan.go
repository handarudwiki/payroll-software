package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type LoanStatus string

const (
	LoanStatusActive   LoanStatus = "active"
	LoanStatusInactive LoanStatus = "cancelled"
	LoanStatusPaid     LoanStatus = "paid"
)

type Loan struct {
	ID                 int        `json:"id"`
	EmployeeID         int        `json:"employee_id"`
	TotalAmount        float64    `json:"total_amount"`
	MonthlyInstallment float64    `json:"monthly_installment"`
	RemainingAmount    float64    `json:"remaining_amount"`
	StartDate          time.Time  `json:"start_date"`
	Status             LoanStatus `json:"status"`
	Employee           Employee   `json:"employee"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

func NewLoanFromCreateLoan(data dto.CreateLoan) (Loan, error) {
	startDate, err := time.Parse("2006-01-02", data.StartDate)
	if err != nil {
		return Loan{}, err
	}
	return Loan{
		EmployeeID:         data.EmployeeID,
		TotalAmount:        data.TotalAmount,
		MonthlyInstallment: data.MonthlyInstallment,
		RemainingAmount:    data.RemainingAmount,
		StartDate:          startDate,
		Status:             LoanStatus(data.Status),
	}, nil
}

func NewLoanFromUpdateLoan(data dto.UpdateLoan) (Loan, error) {
	startDate, err := time.Parse("2006-01-02", data.StartDate)
	if err != nil {
		return Loan{}, err
	}
	return Loan{
		EmployeeID:         data.EmployeeID,
		TotalAmount:        data.TotalAmount,
		MonthlyInstallment: data.MonthlyInstallment,
		RemainingAmount:    data.RemainingAmount,
		StartDate:          startDate,
		Status:             LoanStatus(data.Status),
	}, nil
}
