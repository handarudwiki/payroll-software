package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/services"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
)

type (
	Leave interface {
		Create(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		FindAll(ctx *gin.Context)
	}
	leave struct {
		service services.Leave
	}
)

func NewLeaveController(service services.Leave) Leave {
	return &leave{
		service: service,
	}
}

func (c *leave) Create(ctx *gin.Context) {
	var createRequest dto.CreateLeave
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	createRequest.EmployeeID = ctx.GetInt("id")

	errors := utils.ValidateRequest(createRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	createResponse, err := c.service.Create(ctx, createRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, createResponse)
}
func (c *leave) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	leave, err := c.service.FindByID(ctx, intID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, leave)
}
func (c *leave) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateRequest dto.UpdateLeave
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(updateRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	updateResponse, err := c.service.Update(ctx, intID, updateRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, updateResponse)
}
func (c *leave) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(ctx, intID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, "Leave deleted successfully")
}
func (c *leave) FindAll(ctx *gin.Context) {
	var baseQuery dto.BaseQuery
	if err := ctx.ShouldBindQuery(&baseQuery); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(baseQuery)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	leaves, totalData, err := c.service.FindAll(ctx, baseQuery)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponsePagination(ctx, leaves, totalData)
}
