package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/database/connections"
	"github.com/handarudwiki/payroll-sistem/internal/routes"
	"github.com/handarudwiki/payroll-sistem/internal/schedulers"
)

func main() {
	cfg := config.GetConfig()

	db, err := connections.GetDatabaseConnection(cfg)

	if err != nil {
		panic(err)
	}

	app := gin.New()

	routes.InitUser(db, cfg.JWT, app)
	routes.InitDepartment(db, cfg.JWT, app)
	routes.InitPosition(db, cfg.JWT, app)
	routes.InitEmployee(db, cfg.JWT, app)
	routes.InitSalaryComponent(db, cfg.JWT, app)
	routes.InitEmployeeComponent(db, cfg.JWT, app)
	routes.InitAttendance(db, cfg.JWT, app)
	routes.InitLeave(db, cfg.JWT, app)
	routes.InitLoan(db, cfg.JWT, app)
	routes.InitPayroll(db, cfg.JWT, app)

	schedulers.InitSchedulerAttendance(db)

	app.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
}
