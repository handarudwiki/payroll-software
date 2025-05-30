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

func InitDepartment(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	departmentRepos := repositories.NewDepartmentRepository(db)
	departmentService := services.NewDepartmentService(departmentRepos)
	departmentController := controllers.NewDepartmentController(departmentService)

	// Department routes
	department := router.Group("/department")
	{
		department.GET("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), departmentController.FindAll)
		department.GET("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), departmentController.FindByID)
		department.POST("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), departmentController.Create)
		department.PUT("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), departmentController.Update)
		department.DELETE("/:id", middlewares.AuthMiddleware(jwt), departmentController.Delete)
	}
}
