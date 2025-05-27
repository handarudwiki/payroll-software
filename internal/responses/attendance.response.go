package responses

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/models"
)

type Attendance struct {
	ID           int       `json:"id"`
	EmployeeID   int       `json:"employee_id"`
	Employee     Employee  `json:"employee"`
	Date         time.Time `json:"date"`
	Status       string    `json:"status"`
	WorkingHours int       `json:"working_hours"`
}

func NewAttendanceResponse(attendance models.Attendance) Attendance {
	return Attendance{
		ID:           attendance.ID,
		EmployeeID:   attendance.EmployeeID,
		Employee:     NewEmployeeResponse(attendance.Employee),
		Date:         attendance.Date,
		Status:       string(attendance.Status),
		WorkingHours: attendance.WorkingHours,
	}
}

func NewAttendanceResponses(attendances []models.Attendance) []Attendance {
	responses := make([]Attendance, len(attendances))
	for i, attendance := range attendances {
		responses[i] = NewAttendanceResponse(attendance)
	}
	return responses
}
