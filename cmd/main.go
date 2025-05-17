package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/database/connections"
	"github.com/handarudwiki/payroll-sistem/internal/routes"
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

	app.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
}
