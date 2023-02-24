package handler

import (
	"github.com/gin-gonic/gin"
	"im/model"
	"net/http"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, model.ApiResult{
		Code:    200,
		Message: "成功",
		Data:    data,
	})
}

func Fail(c *gin.Context, errorResult model.ErrorResult, customErrorMessage ...string) {
	errorMessage := errorResult.ErrorMessage
	for _, str := range customErrorMessage {
		errorMessage += " " + str
	}

	c.JSON(http.StatusOK, model.ApiResult{
		Code:    errorResult.ErrorCode,
		Message: errorMessage,
	})
}
