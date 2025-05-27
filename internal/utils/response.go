package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/payroll-sistem/internal/models/commons"
)

type response map[string]interface{}

type ErrValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var ErrToResponseCode = map[error]int{
	commons.ErrCredentials:   http.StatusUnauthorized,
	commons.ErrNotfound:      http.StatusNotFound,
	commons.ErrInvalidInput:  http.StatusBadRequest,
	commons.ErrConflict:      http.StatusConflict,
	commons.ErrInvalidToken:  http.StatusUnauthorized,
	commons.ErrWrongPassword: 401,
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, response{
		"data": data,
	})
}

func ResponsePagination(ctx *gin.Context, data interface{}, paginate commons.Pagination) {
	ctx.JSON(http.StatusOK, response{
		"data": data,
		"meta": paginate,
	})
}

func ResponseError(ctx *gin.Context, message string, code int) {
	ctx.JSON(code, response{
		"message": message,
	})
}

func ResponseValidationError(ctx *gin.Context, errors interface{}) {
	ctx.JSON(http.StatusBadRequest, response{
		"message": "Validation error",
		"errors":  errors,
	})
}

func GetHttpStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if code, ok := ErrToResponseCode[err]; ok {
		return code
	}
	return http.StatusInternalServerError
}
