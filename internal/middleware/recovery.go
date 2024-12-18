package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"rummy-logic-v3/global"
	"rummy-logic-v3/pkg/xerr"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCaller(1).WithCallersFrames().Errorf(c, s, err)
				app.NewResponse(c).ToResponse(xerr.ServerError)
				c.Abort()
			}
		}()
	}
}
