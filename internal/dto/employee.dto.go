package dto

type CreateEmployee struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	NIK          string `json:"nik" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	PositionID   int    `json:"position_id" validate:"required"`
	DepartmentID int    `json:"department_id" validate:"required"`
	HireDate     string `json:"hire_date" validate:"required,datetime=2006-01-02"`
	Status       string `json:"status" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type UpdateEmployee struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	NIK          string `json:"nik" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	PositionID   int    `json:"position_id" validate:"required"`
	DepartmentID int    `json:"department_id" validate:"required"`
	HireDate     string `json:"hire_date" validate:"required,datetime=2006-01-02"`
	Status       string `json:"status" validate:"required"`
}
