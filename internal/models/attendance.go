package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type AttendaceStatus string

const (
	Present AttendaceStatus = "present"
	Absent  AttendaceStatus = "absent"
	Sick    AttendaceStatus = "sick"
	OnLeave AttendaceStatus = "on_leave"
	Late    AttendaceStatus = "late"
)

type Attendance struct {
	ID           int             `json:"id"`
	EmployeeID   int             `json:"employee_id"`
	Employee     Employee        `json:"employee"`
	Date         string          `json:"date"`
	Status       AttendaceStatus `json:"status"`
	WorkingHours int             `json:"working_hours"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

func NewAttendanceFromCreateAttendance(data dto.CreateAttendance) Attendance {
	return Attendance{
		EmployeeID:   data.EmployeeID,
		Date:         data.Date,
		Status:       AttendaceStatus(data.Status),
		WorkingHours: data.WorkingHours,
	}
}
func NewAttendanceFromUpdateAttendance(data dto.UpdateAttendance) Attendance {
	return Attendance{
		EmployeeID:   data.EmployeeID,
		Date:         data.Date,
		Status:       AttendaceStatus(data.Status),
		WorkingHours: data.WorkingHours,
	}
}
