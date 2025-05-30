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
	Employee interface {
		FindAll(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}
	employee struct {
		service services.Employee
	}
)

func NewEmployeeController(service services.Employee) Employee {
	return &employee{
		service: service,
	}
}

func (c *employee) FindAll(ctx *gin.Context) {
	var base dto.BaseQuery

	page, limit := utils.GetPaginationParams(ctx)

	base.Page = page
	base.Limit = limit
	base.Search = ctx.Query("search")

	departmentId := ctx.Query("department_id")
	positionId := ctx.Query("position_id")

	if departmentId != "" {
		intDepartmentID, err := strconv.Atoi(departmentId)
		if err != nil {
			utils.ResponseError(ctx, "Invalid department ID", http.StatusBadRequest)
			return
		}
		base.DepartmentID = &intDepartmentID
	}

	if positionId != "" {
		intPositionId, err := strconv.Atoi(positionId)
		if err != nil {
			utils.ResponseError(ctx, "Invalid position ID", http.StatusBadRequest)
			return
		}

		base.PositionID = &intPositionId
	}

	employees, meta, err := c.service.FindAll(ctx, base)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponsePagination(ctx, employees, meta)

}
func (c *employee) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	employee, err := c.service.FindByID(ctx, intID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, employee)
}

func (c *employee) Create(ctx *gin.Context) {
	var createRequest dto.CreateEmployee
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

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
func (c *employee) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateRequest dto.UpdateEmployee
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

func (c *employee) Delete(ctx *gin.Context) {
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

	utils.ResponseSuccess(ctx, "Employee deleted successfully")
}
