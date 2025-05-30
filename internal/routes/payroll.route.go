package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/internal/controllers"
	"github.com/handarudwiki/payroll-sistem/internal/middlewares"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"github.com/handarudwiki/payroll-sistem/internal/services"
	"gorm.io/gorm"
)

func InitPayroll(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	payrollRepos := repositories.NewPayrollRepository(db)
	payslipDetailRepos := repositories.NewPayslipDetailRepository(db)
	emailRepos := repositories.NewEmployeeRepository(db)
	leaveRepo := repositories.NewLeaveRepository(db)
	loanRepo := repositories.NewLoanRepository(db)

	payrollService := services.NewPayrollService(payrollRepos, payslipDetailRepos, emailRepos, leaveRepo, loanRepo, db)
	payrollController := controllers.NewPayrollController(payrollService)

	// Payroll routes
	payroll := router.Group("/payroll")
	{
		payroll.POST("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), payrollController.Create)
		payroll.GET("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), payrollController.FindByID)
		payroll.GET("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), payrollController.FindAll)
		payroll.GET("/excel", payrollController.GenerateExcel)
	}
}
