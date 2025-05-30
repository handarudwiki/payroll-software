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
	Payroll interface {
		Create(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		FindAll(ctx *gin.Context)
	}
	payroll struct {
		service services.Payroll
	}
)

func NewPayrollController(service services.Payroll) Payroll {
	return &payroll{
		service: service,
	}
}
func (c *payroll) Create(ctx *gin.Context) {
	var createRequest dto.CreatePayroll
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}
	errors := utils.ValidateRequest(createRequest)
	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}
	err := c.service.Create(ctx, createRequest)

	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, "Payroll created successfully")
}
func (c *payroll) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	payroll, err := c.service.FindByID(ctx, intID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, payroll)
}

func (c *payroll) FindAll(ctx *gin.Context) {
	var baseQuery dto.BaseQuery

	page, limit := utils.GetPaginationParams(ctx)

	baseQuery.Page = page
	baseQuery.Limit = limit

	res, meta, err := c.service.FindAll(ctx, baseQuery)

	if err != nil {
		statusCode := utils.GetHttpStatusCode(err)

		utils.ResponseError(ctx, err.Error(), statusCode)
		return
	}

	utils.ResponsePagination(ctx, res, meta)
}
