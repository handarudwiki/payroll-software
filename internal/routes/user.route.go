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

func InitUser(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	userRepos := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepos, jwt)
	userController := controllers.NewUserController(userService)

	// User routes
	user := router.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
		user.PUT("/update", middlewares.AuthMiddleware(jwt), userController.Update)
		user.GET("/me", middlewares.AuthMiddleware(jwt), userController.Me)
		user.PUT("/change-password", middlewares.AuthMiddleware(jwt), userController.ChangePassword)
	}
}
