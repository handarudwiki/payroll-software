package responses

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/models"
)

type Leave struct {
	ID         int       `json:"id"`
	EmployeeID int       `json:"employee_id"`
	Type       string    `json:"type"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Status     string    `json:"status"`
	Employee   *Employee `json:"employee"`
}

func NewLeaveResponse(leave models.Leave) Leave {

	var employee Employee

	if leave.Employee != nil {
		employee = NewEmployeeResponse(*leave.Employee)
	}

	return Leave{
		ID:         leave.ID,
		EmployeeID: leave.EmployeeID,
		Type:       string(leave.Type),
		StartDate:  leave.StartDate,
		EndDate:    leave.EndDate,
		Status:     string(leave.Status),
		Employee:   &employee,
	}
}

func NewLeaveResponses(leaves []models.Leave) []Leave {
	responses := make([]Leave, len(leaves))
	for i, leave := range leaves {
		responses[i] = NewLeaveResponse(leave)
	}
	return responses
}
