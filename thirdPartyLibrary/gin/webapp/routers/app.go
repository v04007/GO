package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"webapp/controllers"
	"webapp/logger"
)

func SetupRouters() *gin.Engine {
	gin.SetMode(viper.GetString("app.mode"))
	//r := gin.Default()
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(false))

	r.Static("dist/", "./dist/dist")
	r.LoadHTMLGlob("dist/index.html")

	//r.GET("/index", controllers.IndexHandler)
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	api := r.Group("/api")
	{
		api.POST("signup", controllers.SignUPHandler)
		api.POST("/login", controllers.LoginHandler)
		api.POST("/questions", controllers.QuestionListsHandler)
		api.GET("/question", controllers.QuestionDetailHandler)
		api.GET("/answers", controllers.AnswerListHandler)

		api.Use(controllers.JWTAuthMiddleware()) //需要登录后才能访问
		api.GET("/needLogin", controllers.NeedLoginHandler)
		api.POST("/question", controllers.QuestionSubmitHandler)
		api.GET("/getCategoryList", controllers.CategoryListHandler)
		api.POST("/answer", controllers.AnswerCommitHandler)
	}
	//IncludeAdminRoutes (r) //引入其他路由文件中定义的路由
	return r
}
