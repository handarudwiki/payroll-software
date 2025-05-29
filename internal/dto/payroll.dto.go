package dto

type CreatePayroll struct {
	Period     string `json:"period" type:"date" validate:"required,date"`
	IsAll      bool   `json:"is_all"`
	EmployeIDS []int  `json:"employee_ids" `
}
