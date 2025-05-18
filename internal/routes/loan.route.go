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

func InitLoan(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	loanRepos := repositories.NewLoanRepository(db)
	employeeRepos := repositories.NewEmployeeRepository(db)
	loanService := services.NewLoanService(loanRepos, employeeRepos)
	loanController := controllers.NewLoanController(loanService)

	// Loan routes
	loan := router.Group("/loan")
	{
		loan.POST("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), loanController.Create)
		loan.GET("/", loanController.FindAll)
		loan.GET("/:id", loanController.FindByID)
		loan.PUT("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), loanController.Update)
		loan.DELETE("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), loanController.Delete)
	}
}
