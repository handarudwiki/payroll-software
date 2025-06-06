package responses

import "github.com/handarudwiki/payroll-sistem/internal/models"

type Employee struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	Phone        string      `json:"phone_number"`
	PositionID   int         `json:"position_id"`
	NIK          string      `json:"nik"`
	DepartmentID int         `json:"department_id"`
	HireDate     string      `json:"hire_date"`
	Status       string      `json:"status"`
	Department   *Department `json:"department"`
	Position     *Position   `json:"position"`
}

func NewEmployeeResponse(employee models.Employee) Employee {

	var department Department
	var position Position

	if employee.Department != nil {
		department = NewDepartment(*employee.Department)
	}

	if employee.Position != nil {
		position = NewPositionResponse(*employee.Position)
	}

	return Employee{
		ID:           employee.ID,
		Name:         employee.Name,
		Email:        employee.Email,
		Phone:        employee.Phone,
		PositionID:   employee.PositionID,
		NIK:          employee.NIK,
		DepartmentID: employee.DepartmentID,
		HireDate:     employee.HireDate,
		Status:       string(employee.Status),
		Department:   &department,
		Position:     &position,
	}
}

func NewEmployeeListResponse(employees []models.Employee) []Employee {
	var employeeResponses []Employee
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, NewEmployeeResponse(employee))
	}
	return employeeResponses
}
