package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IncludeAdminRoutes(r *gin.Engine) {
	r.GET("/xxx", func(c *gin.Context) {
		c.String(http.StatusOK, "admin/xxx")
	})
}
