package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type LeaveStatus string
type LeaveType string

const (
	LeaveStatusPending  LeaveStatus = "pending"
	LeaveStatusApproved LeaveStatus = "approved"
	LeaveStatusRejected LeaveStatus = "rejected"

	LeaveTypeAnnual    LeaveType = "annual"
	LeaveTypeSick      LeaveType = "sick"
	LeaveTypeMaternity LeaveType = "unpaid"
)

type Leave struct {
	ID         int         `json:"id"`
	EmployeeID int         `json:"employee_id"`
	Employee   *Employee   `json:"employee"`
	StartDate  time.Time   `json:"start_date"`
	EndDate    time.Time   `json:"end_date"`
	Status     LeaveStatus `json:"status"`
	Type       LeaveType   `json:"type"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func NewLeaveFromCreateLeave(data dto.CreateLeave) (Leave, error) {
	starDate, err := time.Parse("2006-01-02", data.StartDate)
	if err != nil {
		return Leave{}, err
	}
	endDate, err := time.Parse("2006-01-02", data.EndDate)
	if err != nil {
		return Leave{}, err
	}
	return Leave{
		EmployeeID: data.EmployeeID,
		StartDate:  starDate,
		EndDate:    endDate,
		Status:     LeaveStatusPending,
		Type:       LeaveType(data.Type),
	}, nil
}

func NewLeaveFromUpdateLeave(data dto.UpdateLeave) (Leave, error) {
	starDate, err := time.Parse("2006-01-02", data.StartDate)
	if err != nil {
		return Leave{}, err
	}
	endDate, err := time.Parse("2006-01-02", data.EndDate)
	if err != nil {
		return Leave{}, err
	}
	return Leave{
		EmployeeID: data.EmployeeID,
		StartDate:  starDate,
		EndDate:    endDate,
		Status:     LeaveStatus(data.Status),
		Type:       LeaveType(data.Type),
	}, nil
}
