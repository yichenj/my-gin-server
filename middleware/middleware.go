package middleware

import (
	"github.com/my-gin-server/base/applog"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Info.Println("This is a middleware, do sth interesting.")
		c.Set("middleware", "Ha ha ha...")
		c.Next()
	}
}
