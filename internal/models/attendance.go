package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type AttendaceStatus string

const (
	Present       AttendaceStatus = "present"
	Absent        AttendaceStatus = "absent"
	Sick          AttendaceStatus = "sick"
	OnLeave       AttendaceStatus = "on_leave"
	Late          AttendaceStatus = "late"
	AbsentPenalty                 = 50000
	LatePenalty                   = 20000
)

type Attendance struct {
	ID           int             `json:"id"`
	EmployeeID   int             `json:"employee_id"`
	Employee     Employee        `json:"employee"`
	Date         time.Time       `json:"date"`
	Status       AttendaceStatus `json:"status"`
	WorkingHours int             `json:"working_hours"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

func NewAttendanceFromCreateAttendance(data dto.CreateAttendance) (Attendance, error) {

	date := time.Now()

	var status AttendaceStatus

	var working_hours int

	if date.Hour() > 8 {
		status = Late
		working_hours = 8 - (date.Hour() - 8)
	} else {
		status = Present
		working_hours = 8
	}

	return Attendance{
		EmployeeID:   data.EmployeeID,
		Date:         date,
		Status:       status,
		WorkingHours: working_hours,
	}, nil
}

func NewAttendanceFromUpdateAttendance(data dto.UpdateAttendance) (Attendance, error) {
	date, err := time.Parse("2006-01-02", data.Date)
	if err != nil {
		date = time.Now() // Fallback to current time if parsing fails
	}
	if err != nil {
		return Attendance{}, err
	}
	return Attendance{
		EmployeeID:   data.EmployeeID,
		Date:         date,
		Status:       AttendaceStatus(data.Status),
		WorkingHours: data.WorkingHours,
	}, nil
}
