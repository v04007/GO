package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"webapp/dao/mysql"
)

func CategoryListHandler(c *gin.Context) {
	data, err := mysql.GetCategoryList()
	if err != nil {
		zap.L().Error("mysql.GetCategoryList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	zap.L().Debug("get category list success", zap.Any("data", data))
	ResponseSuccess(c, data)
}
