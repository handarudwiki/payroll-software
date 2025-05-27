package schedulers

import (
	"context"
	"fmt"

	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"github.com/handarudwiki/payroll-sistem/internal/services"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func InitSchedulerAttendance(db *gorm.DB) {
	attendanceRepo := repositories.NewAttendanceRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	leaveRepo := repositories.NewLeaveRepository(db)
	schedulersService := services.NewSchedulerService(attendanceRepo, employeeRepo, leaveRepo)

	c := cron.New()

	c.AddFunc("0 11 * * 1-5", func() {
		ctx := context.Background()
		err := schedulersService.ScheduleAttendance(ctx)
		if err != nil {
			fmt.Printf("ERROR SCHEDULING ATTENDANCE : %v\n", err)
		}
	})

	c.Start()
}
