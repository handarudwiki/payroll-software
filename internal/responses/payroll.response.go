package responses

import "github.com/handarudwiki/payroll-sistem/internal/models"

type PayrollDetail struct {
	ComponentID   int                     `json:"component_id"`
	ComponentType string                  `json:"component_type"`
	Amount        float64                 `json:"amount"`
	Component     SalaryComponentResponse `json:"component"`
}

type Payroll struct {
	ID              int             `json:"id"`
	EmployeeID      int             `json:"employee_id"`
	Employee        Employee        `json:"employee"`
	Period          string          `json:"period"`
	TotalAllowances float64         `json:"total_allowance"`
	TotalDeductions float64         `json:"total_deduction"`
	BaseSalary      float64         `json:"base_salary"`
	NetSalary       float64         `json:"net_salary"`
	GenerateAt      string          `json:"generated_at"`
	PayrollDetail   []PayrollDetail `json:"payslip_details"`
}

func NewPayrollDetails(details []models.PayslipDetail) []PayrollDetail {
	var res []PayrollDetail
	for _, detail := range details {
		res = append(res, PayrollDetail{
			ComponentID:   detail.ComponentID,
			ComponentType: string(detail.ComponentType),
			Amount:        detail.Amount,
			Component:     NewSalaryComponentResponse(detail.Component),
		})
	}
	return res
}

func NewPayroll(payroll models.Payroll) Payroll {
	return Payroll{
		ID:              payroll.ID,
		EmployeeID:      payroll.EmployeeID,
		Employee:        NewEmployeeResponse(payroll.Employee),
		Period:          payroll.Period.Format("2006-01"),
		TotalAllowances: payroll.TotalAllowances,
		BaseSalary:      payroll.BaseSalary,
		TotalDeductions: payroll.TotalDeductions,
		NetSalary:       payroll.NetSalary,
		GenerateAt:      payroll.GeneratedAt.Format("2006-01-02 15:04:05"),
		PayrollDetail:   NewPayrollDetails(payroll.PayslipDetail),
	}
}

func NewPayrolls(payrolls []models.Payroll) []Payroll {
	var res []Payroll
	for _, payroll := range payrolls {
		res = append(res, NewPayroll(payroll))
	}
	return res
}
