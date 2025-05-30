package seeds

import (
	"context"
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

func AttendanceSeed(attendanceRepo repositories.Attendance, employeeRepo repositories.Employee) (err error) {
	employees, err := employeeRepo.FindAllActive(context.Background())
	if err != nil {
		return err
	}

	attendances := []models.Attendance{
		{
			EmployeeID:   employees[0].ID,
			Date:         time.Now(),
			Status:       models.Present,
			WorkingHours: 8,
		},
		{
			EmployeeID:   employees[1].ID,
			Date:         time.Now(),
			Status:       models.Present,
			WorkingHours: 8,
		},
		{
			EmployeeID:   employees[2].ID,
			Date:         time.Now(),
			Status:       models.Late,
			WorkingHours: 6,
		},
		{
			EmployeeID:   employees[3].ID,
			Date:         time.Now(),
			Status:       models.Absent,
			WorkingHours: 0,
		},
		{
			EmployeeID:   employees[4].ID,
			Date:         time.Now(),
			Status:       models.Present,
			WorkingHours: 9,
		},
		{
			EmployeeID:   employees[5].ID,
			Date:         time.Now(),
			Status:       models.Present,
			WorkingHours: 8,
		},
		{
			EmployeeID:   employees[6].ID,
			Date:         time.Now(),
			Status:       models.Late,
			WorkingHours: 7,
		},
		{
			EmployeeID:   employees[7].ID,
			Date:         time.Now(),
			Status:       models.Present,
			WorkingHours: 8,
		},
		{
			EmployeeID:   employees[8].ID,
			Date:         time.Now(),
			Status:       models.Absent,
			WorkingHours: 0,
		},
		{
			EmployeeID:   employees[9].ID,
			Date:         time.Now(),
			Status:       models.Present,
			WorkingHours: 8,
		},
		{
			EmployeeID:   employees[10].ID,
			Date:         time.Now(),
			Status:       models.Late,
			WorkingHours: 6,
		},
	}

	_, err = attendanceRepo.BulkCreate(context.Background(), attendances)

	if err != nil {
		return err
	}

	return nil
}
