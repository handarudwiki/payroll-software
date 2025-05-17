package dto

type CreateDepartment struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateDepartment struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
