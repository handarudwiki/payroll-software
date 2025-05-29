package dto

type CreateLeave struct {
	EmployeeID int    `json:"employee_id,omitempty" validate:"required"`
	StartDate  string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate    string `json:"end_date" validate:"required,datetime=2006-01-02"`
	Type       string `json:"type" validate:"required,oneof=annual sick unpaid"`
}

type UpdateLeave struct {
	EmployeeID int    `json:"employee_id" validate:"required"`
	StartDate  string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate    string `json:"end_date" validate:"required,datetime=2006-01-02"`
	Status     string `json:"status" validate:"required,oneof=pending approved rejected"`
	Type       string `json:"type" validate:"required,oneof=annual sick unpaid"`
}
