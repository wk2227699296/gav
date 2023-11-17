package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WkStudentRouter struct {
}

// InitWkStudentRouter 初始化 wkStudent 路由信息
func (s *WkStudentRouter) InitWkStudentRouter(Router *gin.RouterGroup) {
	wkStudentRouter := Router.Group("wkStudent").Use(middleware.OperationRecord())
	wkStudentRouterWithoutRecord := Router.Group("wkStudent")
	var wkStudentApi = v1.ApiGroupApp.WkTestApiGroup.WkStudentApi
	{
		wkStudentRouter.POST("createWkStudent", wkStudentApi.CreateWkStudent)             // 新建wkStudent
		wkStudentRouter.DELETE("deleteWkStudent", wkStudentApi.DeleteWkStudent)           // 删除wkStudent
		wkStudentRouter.DELETE("deleteWkStudentByIds", wkStudentApi.DeleteWkStudentByIds) // 批量删除wkStudent
		wkStudentRouter.PUT("updateWkStudent", wkStudentApi.UpdateWkStudent)              // 更新wkStudent
	}
	{
		wkStudentRouterWithoutRecord.GET("findWkStudent", wkStudentApi.FindWkStudent)       // 根据ID获取wkStudent
		wkStudentRouterWithoutRecord.GET("getWkStudentList", wkStudentApi.GetWkStudentList) // 获取wkStudent列表
	}
}
