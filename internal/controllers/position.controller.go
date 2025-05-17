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
	Position interface {
		Create(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		FindAll(ctx *gin.Context)
	}
	position struct {
		service services.Position
	}
)

func NewPositionController(service services.Position) Position {
	return &position{
		service: service,
	}
}

func (c *position) Create(ctx *gin.Context) {
	var createRequest dto.CreatePosition
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

func (c *position) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	positionResponse, err := c.service.FindByID(ctx, id)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, positionResponse)
}
func (c *position) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateRequest dto.UpdatePosition
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(updateRequest)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	updateResponse, err := c.service.Update(ctx, id, updateRequest)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, updateResponse)
}
func (c *position) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ResponseError(ctx, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(ctx, id)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponseSuccess(ctx, "Position deleted successfully")
}
func (c *position) FindAll(ctx *gin.Context) {
	var base dto.BaseQuery
	if err := ctx.ShouldBindQuery(&base); err != nil {
		utils.ResponseError(ctx, "Invalid request", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(base)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	positionResponse, meta, err := c.service.FindAll(ctx, base)
	if err != nil {
		code := utils.GetHttpStatusCode(err)
		utils.ResponseError(ctx, err.Error(), code)
		return
	}

	utils.ResponsePagination(ctx, positionResponse, meta)
}
