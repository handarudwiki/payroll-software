package dto

type CreateAttendance struct {
	EmployeeID   int    `json:"employee_id" validate:"required"`
	Date         string `json:"date" validate:"required,datetime=2006-01-02"`
	Status       string `json:"status" validate:"required,oneof=present absent sick on_leave late"`
	WorkingHours int    `json:"working_hours" validate:"required"`
}

type UpdateAttendance struct {
	EmployeeID   int    `json:"employee_id" validate:"required"`
	Date         string `json:"date" validate:"required,datetime=2006-01-02"`
	Status       string `json:"status" validate:"required,oneof=present absent sick on_leave late"`
	WorkingHours int    `json:"working_hours" validate:"required"`
}
