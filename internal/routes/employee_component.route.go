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

func InitEmployeeComponent(db *gorm.DB, jwt config.JWT, app *gin.Engine) {
	employeeComponentRepos := repositories.NewEmployeeComponentRepository(db)
	salaryComponentRepo := repositories.NewSalaryComponentRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	employeeComponentService := services.NewEmployeeComponentService(employeeComponentRepos, employeeRepo, salaryComponentRepo)
	employeeComponentController := controllers.NewEmployeeComponentController(employeeComponentService)

	// Employee Component routes
	employeeComponent := app.Group("/employee-component")
	{
		employeeComponent.POST("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), employeeComponentController.Create)
		employeeComponent.GET("/", employeeComponentController.FindAll)
		employeeComponent.GET("/:id", employeeComponentController.FindByID)
		employeeComponent.PUT("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), employeeComponentController.Update)
		employeeComponent.DELETE("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), employeeComponentController.Delete)
	}
}
