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
	SalaryComponent interface {
		Create(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		FindAll(ctx *gin.Context)
	}

	salaryComponent struct {
		service services.SalaryComponent
	}
)

func NewSalaryComponentController(service services.SalaryComponent) SalaryComponent {
	return &salaryComponent{
		service: service,
	}
}

func (c *salaryComponent) Create(ctx *gin.Context) {
	var createRequest dto.CreateSalaryComponent
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

func (c *salaryComponent) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	salaryComponent, err := c.service.FindByID(ctx, idInt)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, salaryComponent)
}
func (c *salaryComponent) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateRequest dto.UpdateSalaryComponent
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(updateRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	updateResponse, err := c.service.Update(ctx, idInt, updateRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, updateResponse)
}
func (c *salaryComponent) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(ctx, idInt)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, "Salary component deleted successfully")
}

func (c *salaryComponent) FindAll(ctx *gin.Context) {
	page, limit := utils.GetPaginationParams(ctx)
	search := ctx.Query("search")

	baseQuery := dto.BaseQuery{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

	salaryComponents, meta, err := c.service.FindAll(ctx, baseQuery)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponsePagination(ctx, salaryComponents, meta)
}
