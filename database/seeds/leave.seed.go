package seeds

import (
	"context"
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

func LeaveSeed(leaveRepo repositories.Leave, employeeRepo repositories.Employee) (err error) {

	employees, err := employeeRepo.FindAllActive(context.Background())
	if err != nil {
		return err
	}

	leaves := []models.Leave{
		{
			EmployeeID: employees[0].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 5),
			Status:     models.LeaveStatusPending,
			Type:       models.LeaveTypeAnnual,
		},
		{
			EmployeeID: employees[1].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 3),
			Status:     models.LeaveStatusApproved,
			Type:       models.LeaveTypeAnnual,
		},
		{
			EmployeeID: employees[2].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 2),
			Status:     models.LeaveStatusRejected,
			Type:       models.LeaveTypeAnnual,
		},
		{
			EmployeeID: employees[3].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 4),
			Status:     models.LeaveStatusPending,
			Type:       models.LeaveTypeMaternity,
		},
		{
			EmployeeID: employees[4].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 1, 0),
			Status:     models.LeaveStatusApproved,
			Type:       models.LeaveTypeMaternity,
		},
		{
			EmployeeID: employees[5].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 7),
			Status:     models.LeaveStatusPending,
			Type:       models.LeaveTypeSick,
		},
		{
			EmployeeID: employees[6].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 10),
			Status:     models.LeaveStatusApproved,
			Type:       models.LeaveTypeSick,
		},
		{
			EmployeeID: employees[7].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 1),
			Status:     models.LeaveStatusRejected,
			Type:       models.LeaveTypeAnnual,
		},
		{
			EmployeeID: employees[8].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 6),
			Status:     models.LeaveStatusPending,
			Type:       models.LeaveTypeSick,
		},
		{
			EmployeeID: employees[9].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 2),
			Status:     models.LeaveStatusApproved,
			Type:       models.LeaveTypeAnnual,
		},
		{
			EmployeeID: employees[10].ID,
			StartDate:  time.Now(),
			EndDate:    time.Now().AddDate(0, 0, 4),
			Status:     models.LeaveStatusPending,
			Type:       models.LeaveTypeMaternity,
		},
	}

	_, err = leaveRepo.BulkCreate(context.Background(), leaves)
	if err != nil {
		return err
	}

	return nil
}
