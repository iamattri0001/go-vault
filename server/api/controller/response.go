package controller

import (
	customerrors "go-vault/custom_errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SendResponse(ctx *gin.Context, success bool, message string, data any, err error) {
	if err != nil {
		ctx.JSON(customerrors.GetCode(err), newResponse(success, message, nil, err.Error()))
		return
	}
	ctx.JSON(200, newResponse(success, message, data, ""))
}

func newResponse(success bool, message string, data interface{}, err string) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	}
}
