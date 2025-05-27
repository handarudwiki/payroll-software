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

func InitLeave(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	leaveRepos := repositories.NewLeaveRepository(db)
	employeeRepos := repositories.NewEmployeeRepository(db)
	leaveService := services.NewLeaveService(leaveRepos, employeeRepos)
	leaveController := controllers.NewLeaveController(leaveService)

	// Leave routes
	leave := router.Group("/leave")
	{
		leave.POST("/", middlewares.AuthMiddleware(jwt), leaveController.Create)
		leave.GET("/", middlewares.AuthMiddleware(jwt), leaveController.FindAll)
		leave.GET("/:id", middlewares.AuthMiddleware(jwt), leaveController.FindByID)
		leave.PUT("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), leaveController.Update)
		leave.DELETE("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), leaveController.Delete)
	}
}
