package router

import (
	"github.com/my-gin-server/api"
	"github.com/my-gin-server/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterDemoAPI(e *gin.Engine, demoApi *api.DemoAPI) {
	middlewareGroup := e.Group("")
	middlewareGroup.Use(middleware.Middleware())
	middlewareGroup.POST("/demo", demoApi.Create)
	middlewareGroup.DELETE("/demo", demoApi.DeleteRange)
	middlewareGroup.DELETE("/demo/:id", demoApi.Delete)

	nonMiddlewareGroup := e.Group("")
	nonMiddlewareGroup.GET("/demo", demoApi.List)
	nonMiddlewareGroup.GET("/demo/:id", demoApi.Query)
}
