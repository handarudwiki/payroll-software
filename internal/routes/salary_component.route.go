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

func InitSalaryComponent(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	salaryComponentRepos := repositories.NewSalaryComponentRepository(db)
	salaryComponentService := services.NewSalaryComponentService(salaryComponentRepos)
	salaryComponentController := controllers.NewSalaryComponentController(salaryComponentService)

	// Salary Component routes
	salaryComponent := router.Group("/salary-component")
	{
		salaryComponent.POST("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), salaryComponentController.Create)
		salaryComponent.GET("/", salaryComponentController.FindAll)
		salaryComponent.GET("/:id", salaryComponentController.FindByID)
		salaryComponent.PUT("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), salaryComponentController.Update)
		salaryComponent.DELETE("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), salaryComponentController.Delete)
	}
}
