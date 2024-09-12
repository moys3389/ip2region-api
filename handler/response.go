package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseSuccess(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{
		Status:  200,
		Message: "success",
		Data:    data,
	})
}

func ResponseFail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, Response{
		Status:  500,
		Message: err.Error(),
	})
}

func ResponseErrorParam(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Status:  400,
		Message: "param is error",
	})
}
