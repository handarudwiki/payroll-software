package main

import (
	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/database/connections"
)

func main() {
	cfg := config.GetConfig()

	db, err := connections.GetDatabaseConnection(cfg)

	if err != nil {
		panic(err)
	}

}
