package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/services"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
)

type (
	User interface {
		Login(ctx *gin.Context)
		Register(ctx *gin.Context)
		Update(ctx *gin.Context)
		Me(ctx *gin.Context)
		ChangePassword(ctx *gin.Context)
	}
	userController struct {
		userService services.User
	}
)

func NewUserController(userService services.User) *userController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Login(ctx *gin.Context) {
	var loginRequest dto.Login
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(loginRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	loginResponse, err := c.userService.Login(ctx, loginRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, loginResponse)
}

func (c *userController) Register(ctx *gin.Context) {
	var registerRequest dto.Register
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(registerRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	userResponse, err := c.userService.Register(ctx, registerRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, userResponse)
}

func (c *userController) Update(ctx *gin.Context) {
	var updateRequest dto.UpdateUser
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(updateRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}
	id := ctx.GetInt("id")
	userResponse, err := c.userService.Update(ctx, id, updateRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, userResponse)
}

func (c *userController) Me(ctx *gin.Context) {
	id := ctx.GetInt("id")
	userResponse, err := c.userService.Me(ctx, id)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, userResponse)
}

func (c *userController) ChangePassword(ctx *gin.Context) {
	var updatePasswordRequest dto.UpdatePassword
	if err := ctx.ShouldBindJSON(&updatePasswordRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(updatePasswordRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	id := ctx.GetInt("id")
	userResponse, err := c.userService.UpdatePassword(ctx, id, updatePasswordRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, userResponse)
}
