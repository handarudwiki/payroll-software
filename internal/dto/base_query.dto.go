package dto

type BaseQuery struct {
	Limit        int    `json:"limit" form:"limit"`
	Page         int    `json:"page" form:"page"`
	Search       string `json:"search" form:"search"`
	DepartmentID *int   `json:"department_id" form:"department_id"`
	PositionID   *int   `json:"position_id" form:"position_id"`
}
