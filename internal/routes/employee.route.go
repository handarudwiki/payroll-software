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

func InitEmployee(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	employeeRepos := repositories.NewEmployeeRepository(db)
	userRepos := repositories.NewUserRepository(db)
	departmentRepo := repositories.NewDepartmentRepository(db)
	positionRepo := repositories.NewPositionRepository(db)
	employeeService := services.NewEmployeeService(employeeRepos, departmentRepo, positionRepo, userRepos)
	employeeController := controllers.NewEmployeeController(employeeService)

	// Employee routes
	employee := router.Group("/employee")
	{
		employee.POST("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), employeeController.Create)
		employee.GET("/", employeeController.FindAll)
		employee.GET("/:id", employeeController.FindByID)
		employee.PUT("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), employeeController.Update)
		employee.DELETE("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), employeeController.Delete)
	}
}
