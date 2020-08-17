package main

import (
	"github.com/my-gin-server/api"
	"github.com/my-gin-server/dao"
	"github.com/my-gin-server/router"
	"github.com/my-gin-server/service"

	"github.com/gin-gonic/gin"
)

func InitWorld(e *gin.Engine) {
	demoDao := dao.NewDemoDao()
	demoAPI := api.NewDemoAPI(service.NewDemoService(demoDao))

	router.RegisterDemoAPI(e, demoAPI)
}
