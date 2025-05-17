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

func InitAttendance(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	attendanceRepos := repositories.NewAttendanceRepository(db)
	employeeRepos := repositories.NewEmployeeRepository(db)
	attendanceService := services.NewAttendanceService(attendanceRepos, employeeRepos)
	attendanceController := controllers.NewAttendanceController(attendanceService)

	// Attendance routes
	attendance := router.Group("/attendance")
	{
		attendance.POST("/", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), attendanceController.Create)
		attendance.GET("/", attendanceController.FindAll)
		attendance.GET("/:id", attendanceController.FindByID)
		attendance.PUT("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), attendanceController.Update)
		attendance.DELETE("/:id", middlewares.AuthMiddleware(jwt), middlewares.AuthorizationMiddleware("admin"), attendanceController.Delete)
	}
}
