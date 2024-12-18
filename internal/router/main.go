package router

import (
	"github.com/gin-gonic/gin"
	"rummy-logic-v3/global"
	"rummy-logic-v3/internal/controller"
	"rummy-logic-v3/internal/middleware"
)

func NewRouter() (r *gin.Engine) {
	r = gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}

	r.Use(middleware.Tracing())
	r.Use(middleware.Recovery()) // 异常捕获
	r.Use(middleware.AccessLog())
	//r.Use(middleware.CheckSign())
	r.Use(middleware.Translations())
	//r.Use(middleware.CheckIp())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	group := r.RouterGroup.Group("/api/v1")
	{
		controller.RegHandRouter(group)
	}
	return
}
