package e2e

import (
	"log"
	"os"
	"testing"

	utils_test "github.com/handarudwiki/payroll-sistem/tests/utils"
)

func TestMain(m *testing.M) {

	db, _ := utils_test.SetupTestSuite()

	exitCode := m.Run()

	err := utils_test.CleanAllTables(db)
	if err != nil {
		log.Fatal("Failed to clean up test database: ", err)
	}
	utils_test.TearDownSuite()

	os.Exit(exitCode)

}
