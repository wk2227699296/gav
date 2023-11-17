package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WkInfoRouter struct {
}

// InitWkInfoRouter 初始化 WkInfo 路由信息
func (s *WkInfoRouter) InitWkInfoRouter(Router *gin.RouterGroup) {
	wkinfoRouter := Router.Group("wkinfo").Use(middleware.OperationRecord())
	wkinfoRouterWithoutRecord := Router.Group("wkinfo")
	var wkinfoApi = v1.ApiGroupApp.WkTestApiGroup.WkInfoApi
	{
		wkinfoRouter.POST("createWkInfo", wkinfoApi.CreateWkInfo)             // 新建WkInfo
		wkinfoRouter.DELETE("deleteWkInfo", wkinfoApi.DeleteWkInfo)           // 删除WkInfo
		wkinfoRouter.DELETE("deleteWkInfoByIds", wkinfoApi.DeleteWkInfoByIds) // 批量删除WkInfo
		wkinfoRouter.PUT("updateWkInfo", wkinfoApi.UpdateWkInfo)              // 更新WkInfo
	}
	{
		wkinfoRouterWithoutRecord.GET("findWkInfo", wkinfoApi.FindWkInfo)       // 根据ID获取WkInfo
		wkinfoRouterWithoutRecord.GET("getWkInfoList", wkinfoApi.GetWkInfoList) // 获取WkInfo列表
	}
}
