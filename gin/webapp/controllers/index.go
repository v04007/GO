package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	//zap.L().Error("this is a error")
	//zap.L().Debug("this is index handler")
	//zap.L().Info("this is a test log")
	//time.Sleep(time.Second * 5)
	c.String(http.StatusOK, viper.GetString("app.ver"))
}

func NeedLoginHandler(c *gin.Context) {
	//因为之前的JWT中间件在 上下文C 中保存了 当前登录的userID
	//我在这个函数中就可以通过 上下文c 获取当前登录的用户
	v, ok := c.Get(ContextUserIDKey)
	fmt.Println("value:", v, "ok:", ok)
	if !ok {
		//没有取到userID,-->/login
		ResponseError(c, CodeNotLogin)
		return
	}
	userID := v.(uint64)
	ResponseSuccess(c, userID)
}
