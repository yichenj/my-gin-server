package api

import (
	"net/http"
	"strconv"

	"github.com/my-gin-server/api/vo"
	"github.com/my-gin-server/base/apperror"
	"github.com/my-gin-server/service"
	"github.com/my-gin-server/service/dto"

	"github.com/gin-gonic/gin"
)

type DemoAPI struct {
	demoService *service.DemoService
}

func NewDemoAPI(demoService *service.DemoService) *DemoAPI {
	return &DemoAPI{demoService: demoService}
}

func (api *DemoAPI) Create(ctx *gin.Context) {
	var demoVO vo.Demo
	if err := ctx.BindJSON(&demoVO); err != nil {
		ctx.JSON(http.StatusBadRequest, apperror.New(apperror.ERR_INPUT_FORMAT, err.Error()))
		return
	}

	extraInfo := ""
	middlewareInfo, exist := ctx.Get("middleware")
	if exist {
		extraInfo, _ = middlewareInfo.(string)
	}

	demoDTO := dto.DemoDTO{Name: demoVO.Name, Description: demoVO.Description, ExtraInfo: extraInfo}
	demoID, err := api.demoService.Create(&demoDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": demoID})
}

func (api *DemoAPI) Delete(ctx *gin.Context) {
	demoId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, apperror.ErrInvalidResource)
		return
	}

	err = api.demoService.DeleteById(demoId)
	if err != nil {
		if err == apperror.ErrNotFound {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": demoId})
}

func (api *DemoAPI) DeleteRange(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	numOfDeletion, err := api.demoService.DeleteRange(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"NumOfDeletion": numOfDeletion})
}

func (api *DemoAPI) List(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	demos, total, err := api.demoService.List(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Total":    total,
		"DemoList": demos,
	})
}

func (api *DemoAPI) Query(ctx *gin.Context) {
	demoId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, apperror.ErrInvalidResource)
		return
	}

	demo, err := api.demoService.QueryById(demoId)
	if err != nil {
		if err == apperror.ErrNotFound {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, demo)
}
