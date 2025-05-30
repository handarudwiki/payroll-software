package utils_test

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/database/connections"
	"gorm.io/gorm"
)

type TestApp struct {
	DB     *gorm.DB
	Router *gin.Engine
	Config *config.Config
}

var (
	testDB     *gorm.DB
	testConfig *config.Config
)

func SetupTestSuite() (*gorm.DB, *config.Config) {
	if testDB != nil && testConfig != nil {
		return testDB, testConfig
	}

	cfg := config.LoadTestConfig()

	db, err := connections.GetDatabaseConnection(cfg)

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	testDB = db
	testConfig = cfg

	return testDB, testConfig

}

func NewTestApp(t *testing.T) *TestApp {
	db, cfg := SetupTestSuite()

	if db == nil || cfg == nil {
		t.Fatal("Failed to set up test suite")
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.SetTrustedProxies(nil) // Disable trusted proxies for testing

	return &TestApp{
		DB:     db,
		Router: router,
		Config: cfg,
	}
}

func TearDownSuite() {
	if testDB != nil {
		sqlDB, err := testDB.DB()
		if err != nil {
			log.Printf("failed to get sql DB from gorm: %v", err)
			return
		}

		err = sqlDB.Close()
		if err != nil {
			log.Printf("failed to close test DB: %v", err)
		} else {
			log.Println("test DB closed successfully")
		}

		testDB = nil
	}
}
