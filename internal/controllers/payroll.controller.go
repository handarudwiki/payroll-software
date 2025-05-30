package controllers

import (
	"fmt"
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
		GenerateExcel(ctx *gin.Context)
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
	fmt.Println("Create Request:", createRequest)

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

func (c *payroll) GenerateExcel(ctx *gin.Context) {
	var baseQuery dto.BaseQuery

	page, limit := utils.GetPaginationParams(ctx)

	baseQuery.Page = page
	baseQuery.Limit = limit

	employeeId := ctx.Query("employee_id")

	if employeeId != "" {
		intEmployeeID, err := strconv.Atoi(employeeId)
		if err != nil {
			utils.ResponseError(ctx, "Invalid employee ID", http.StatusBadRequest)
			return
		}
		baseQuery.EmployeeID = &intEmployeeID
	}

	fileBytes, err := c.service.GenerateExcel(ctx, baseQuery)

	if err != nil {
		statusCode := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), statusCode)
		return
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename=report.xlsx")
	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}
