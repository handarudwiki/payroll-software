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
	Department interface {
		FindAll(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}
	departmentController struct {
		departmentService services.Department
	}
)

func NewDepartmentController(departmentService services.Department) *departmentController {
	return &departmentController{
		departmentService: departmentService,
	}
}
func (c *departmentController) FindAll(ctx *gin.Context) {
	page, limit := utils.GetPaginationParams(ctx)

	search := ctx.Query("search")

	base := dto.BaseQuery{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

	departments, meta, err := c.departmentService.FindAll(ctx, base)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponsePagination(ctx, departments, meta)
}

func (c *departmentController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	departmentID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid department ID", http.StatusBadRequest)
		return
	}
	department, err := c.departmentService.FindByID(ctx, departmentID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, department)
}

func (c *departmentController) Create(ctx *gin.Context) {
	var departmentRequest dto.CreateDepartment
	if err := ctx.ShouldBindJSON(&departmentRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(departmentRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	department, err := c.departmentService.Create(ctx, departmentRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, department)
}
func (c *departmentController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	departmentID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid department ID", http.StatusBadRequest)
		return
	}

	var departmentRequest dto.UpdateDepartment
	if err := ctx.ShouldBindJSON(&departmentRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(departmentRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	department, err := c.departmentService.Update(ctx, departmentID, departmentRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, department)
}

func (c *departmentController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	departmentID, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, "Invalid department ID", http.StatusBadRequest)
		return
	}

	err = c.departmentService.Delete(ctx, departmentID)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}
	utils.ResponseSuccess(ctx, "Department deleted successfully")
}
