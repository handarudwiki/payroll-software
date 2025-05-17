package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/internal/controllers"
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
	}
}
