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
	Loan interface {
		Create(ctx *gin.Context)
		FindAll(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	loan struct {
		service services.Loan
	}
)

func NewLoanController(service services.Loan) Loan {
	return &loan{
		service: service,
	}
}

func (c *loan) Create(ctx *gin.Context) {
	var createRequest dto.CreateLoan
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
func (c *loan) FindAll(ctx *gin.Context) {
	baseQuery := dto.BaseQuery{
		Page:  1,
		Limit: 10,
	}

	if err := ctx.ShouldBindQuery(&baseQuery); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	loans, meta, err := c.service.FindAll(ctx, baseQuery)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponsePagination(ctx, loans, meta)
}
func (c *loan) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	loan, err := c.service.FindByID(ctx, intID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, loan)
}
func (c *loan) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateRequest dto.UpdateLoan
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
func (c *loan) Delete(ctx *gin.Context) {
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

	utils.ResponseSuccess(ctx, "Loan deleted successfully")
}
