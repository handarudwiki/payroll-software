package main

import (
	"fmt"

	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/database/connections"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"github.com/handarudwiki/payroll-sistem/internal/seeds"
)

func main() {
	cfg := config.GetConfig()

	db, err := connections.GetDatabaseConnection(cfg)

	if err != nil {
		panic(err)
	}

	userRepo := repositories.NewUserRepository(db)
	positionRepo := repositories.NewPositionRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	leaveRepo := repositories.NewLeaveRepository(db)
	loanRepo := repositories.NewLoanRepository(db)
	salaryComponentRepo := repositories.NewSalaryComponentRepository(db)
	departmentRepo := repositories.NewDepartmentRepository(db)
	employeeComponentRepo := repositories.NewEmployeeComponentRepository(db)
	attendanceRepo := repositories.NewAttendanceRepository(db)

	fmt.Println("Seeding database...")
	seeds.UserSeed(userRepo)
	seeds.PositionSeed(positionRepo)
	seeds.DepartmentSeed(departmentRepo)
	seeds.EmployeeSeed(employeeRepo, departmentRepo, positionRepo, userRepo)
	seeds.LeaveSeed(leaveRepo, employeeRepo)
	seeds.LoanSeed(loanRepo, employeeRepo)
	seeds.SalaryComponentSeed(salaryComponentRepo)
	seeds.EmployeeComponentSeed(employeeComponentRepo, salaryComponentRepo, employeeRepo)
	seeds.AttendanceSeed(attendanceRepo, employeeRepo)
	fmt.Println("Database seeding completed successfully!")
}
