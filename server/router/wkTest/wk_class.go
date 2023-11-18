package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WkClassRouter struct {
}

// InitWkClassRouter 初始化 WkClass 路由信息
func (s *WkClassRouter) InitWkClassRouter(Router *gin.RouterGroup) {
	wClassRouter := Router.Group("wClass").Use(middleware.OperationRecord())
	wClassRouterWithoutRecord := Router.Group("wClass")
	var wClassApi = v1.ApiGroupApp.WkTestApiGroup.WkClassApi
	{
		wClassRouter.POST("createWkClass", wClassApi.CreateWkClass)             // 新建WkClass
		wClassRouter.DELETE("deleteWkClass", wClassApi.DeleteWkClass)           // 删除WkClass
		wClassRouter.DELETE("deleteWkClassByIds", wClassApi.DeleteWkClassByIds) // 批量删除WkClass
		wClassRouter.PUT("updateWkClass", wClassApi.UpdateWkClass)              // 更新WkClass
	}
	{
		wClassRouterWithoutRecord.GET("findWkClass", wClassApi.FindWkClass)       // 根据ID获取WkClass
		wClassRouterWithoutRecord.GET("getWkClassList", wClassApi.GetWkClassList) // 获取WkClass列表
	}
}
