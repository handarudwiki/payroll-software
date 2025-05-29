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
	Attendance interface {
		FindAll(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}
	attendance struct {
		service services.Attendance
	}
)

func NewAttendanceController(service services.Attendance) Attendance {
	return &attendance{
		service: service,
	}
}

func (c *attendance) FindAll(ctx *gin.Context) {
	page, limit := utils.GetPaginationParams(ctx)

	search := ctx.Query("search")

	base := dto.BaseQuery{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

	attendances, meta, err := c.service.FindAll(ctx, base)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponsePagination(ctx, attendances, meta)
}

func (c *attendance) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	attendanceID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid attendance ID", http.StatusBadRequest)
		return
	}
	attendance, err := c.service.FindByID(ctx, attendanceID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, attendance)
}

func (c *attendance) Create(ctx *gin.Context) {
	var createRequest dto.CreateAttendance

	createRequest.EmployeeID = ctx.GetInt("id")

	createResponse, err := c.service.Create(ctx, createRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, createResponse)
}

func (c *attendance) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	attendanceID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid attendance ID", http.StatusBadRequest)
		return
	}

	var updateRequest dto.UpdateAttendance
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(updateRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	updateResponse, err := c.service.Update(ctx, attendanceID, updateRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, updateResponse)
}

func (c *attendance) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	attendanceID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid attendance ID", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(ctx, attendanceID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, nil)
}
