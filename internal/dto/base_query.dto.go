package dto

import "time"

type BaseQuery struct {
	Limit        int        `json:"limit" form:"limit"`
	Page         int        `json:"page" form:"page"`
	Search       string     `json:"search" form:"search"`
	EmployeeID   *int       `json:"employee_id" from:"search"`
	Period       *time.Time `json:"period" form:"period"`
	DepartmentID *int       `json:"department_id" form:"department_id"`
	PositionID   *int       `json:"position_id" form:"position_id"`
}
