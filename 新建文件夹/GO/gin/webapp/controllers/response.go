package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    MyCode      `json:"code"`    //业务状态码
	Message string      `json:"message"` //提示信息
	Date    interface{} `json:"date"`    //数据
}

// ResponseError 根据错误状态码返回响应
func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Date:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

// ResponseErrorWithMsg 返回自定义类型错误信息
func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, errMsg string) {
	rd := &ResponseData{
		Code:    code,
		Message: errMsg,
		Date:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

// ResponseSuccess 返回正常响应
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Date:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
