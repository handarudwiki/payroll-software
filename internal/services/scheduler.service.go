package services

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

type (
	Scheduler interface {
		ScheduleAttendance(ctx context.Context) (err error)
	}
	scheduler struct {
		attendanceRepo repositories.Attendance
		employeeRepo   repositories.Employee
		leaveRepo      repositories.Leave
	}
)

func NewSchedulerService(attendanceRepo repositories.Attendance, employeeRepo repositories.Employee, leaveRepo repositories.Leave) Scheduler {
	return &scheduler{
		attendanceRepo: attendanceRepo,
		employeeRepo:   employeeRepo,
		leaveRepo:      leaveRepo,
	}
}

func (s *scheduler) ScheduleAttendance(ctx context.Context) (err error) {
	// Get all employees
	employees, err := s.employeeRepo.FindAllActive(ctx)
	if err != nil {
		return err
	}

	attendancesToday, err := s.attendanceRepo.TodayAttendance(ctx)

	if err != nil {
		return err
	}
	attended := make(map[int]bool)
	for _, attendance := range attendancesToday {
		attended[attendance.EmployeeID] = true
	}

	var notAttendEmployeeIds []int
	for _, employee := range employees {
		if !attended[employee.ID] {
			notAttendEmployeeIds = append(notAttendEmployeeIds, employee.ID)
		}
	}

	leaves, err := s.leaveRepo.FindByEmployeeIDS(ctx, notAttendEmployeeIds)
	if err != nil {
		return err
	}

	leaveMap := make(map[int]models.Leave)
	for _, leave := range leaves {
		leaveMap[leave.EmployeeID] = leave
	}

	var attendanceToCreate []models.Attendance
	for _, employeeId := range notAttendEmployeeIds {
		status := string(models.Absent)

		if leave, exists := leaveMap[employeeId]; exists {
			switch leave.Type {
			case models.LeaveTypeSick:
				status = string(models.Sick)
			case models.LeaveTypeAnnual:
				status = string(models.Present)
			case models.LeaveTypeMaternity:
				status = string(models.OnLeave)
			default:
				status = string(models.Absent)
			}

			attendance := models.Attendance{
				EmployeeID:   employeeId,
				Status:       models.AttendaceStatus(status),
				WorkingHours: 0, // Default working hours, can be updated later
			}
			attendanceToCreate = append(attendanceToCreate, attendance)
		}
	}

	_, err = s.attendanceRepo.BulkCreate(ctx, attendanceToCreate)
	if err != nil {
		return
	}

	return
}
