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

func InitPosition(db *gorm.DB, jwt config.JWT, app *gin.Engine) {
	positionRepos := repositories.NewPositionRepository(db)
	positionService := services.NewPositionService(positionRepos)
	positionController := controllers.NewPositionController(positionService)

	// Position routes
	position := app.Group("/position")
	{
		position.GET("/", positionController.FindAll)
		position.GET("/:id", positionController.FindByID)
		position.POST("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), positionController.Create)
		position.PUT("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), positionController.Update)
		position.DELETE("/:id", middlewares.AuthMiddleware(jwt), positionController.Delete)
	}

}
