package controllers

import (
	"strings"
	"webapp/pkg/myJWT"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	ContextUserIDKey = "userID"
)

//基于JWT实现的登录认证中间件
//对于需要登陆才能访问的API来说
//该中间件需要从请求头中获取JWT Token
//如果没有Token--> /Login
//如果Token期--> /Login
//从JWT中解析我们需要的UserID字段 -->根据userID我们就能从数据库查询到当前请求的用户是谁

//JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//客户端携带Token有三种方式1,放在请求头(推荐)2,放在请求体3,放在URI
		//这里假设Token放在Header的中Authorization: Bearer token_string
		//这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
			c.Abort()
			return
		}
		//按空格分隔
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ResponseErrorWithMsg(c, CodeInvalidToken, "请求头中auth格式问题")
			c.Abort()
			return
		}
		//parts[1]是获取到的tokenString,我们使用之前定义好的解析JWT函数来解析他
		mc, err := myJWT.ParseToken(parts[1])
		if err != nil {
			ResponseError(c, CodeInvalidToken)
			zap.L().Warn("invalid JWT token", zap.Error(err))
			c.Abort()
			return
		}
		//当前请求的username信息保存到请求的上下文c上
		c.Set(ContextUserIDKey, mc.UserID)
		c.Next() //后续处理的函数可以通过使用c.GET("userID")来获取当前请求的用户信息
		//返回响应的时候可以做Token/Cookie续期
	}
}

//1基于Cookie和Session认证的中间件
//对于需要登陆才能访问的API来说
//该中间件需要从请求中获取Cookie值
//如果没有Cookie-->/Login
//拿到Cookie值取session数据中找对应的数据,找不到(1.session期了2,无效的cookie值) --> /Login:
//session值也可以通过c.Set()直接 赋值 到上下文 c 上
