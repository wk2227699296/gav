package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wkTest"
	wkTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/wkTest/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WkClassApi struct {
}

var wClassService = service.ServiceGroupApp.WkTestServiceGroup.WkClassService

// CreateWkClass 创建WkClass
// @Tags WkClass
// @Summary 创建WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkClass true "创建WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":global.Translate("general.getDataSuccess")}"
// @Router /wClass/createWkClass [post]
func (wClassApi *WkClassApi) CreateWkClass(c *gin.Context) {
	var wClass wkTest.WkClass
	err := c.ShouldBindJSON(&wClass)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wClassService.CreateWkClass(&wClass); err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccss"), c)
	}
}

// DeleteWkClass 删除WkClass
// @Tags WkClass
// @Summary 删除WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkClass true "删除WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wClass/deleteWkClass [delete]
func (wClassApi *WkClassApi) DeleteWkClass(c *gin.Context) {
	var wClass wkTest.WkClass
	err := c.ShouldBindJSON(&wClass)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wClassService.DeleteWkClass(wClass); err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deletFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
	}
}

// DeleteWkClassByIds 批量删除WkClass
// @Tags WkClass
// @Summary 批量删除WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wClass/deleteWkClassByIds [delete]
func (wClassApi *WkClassApi) DeleteWkClassByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wClassService.DeleteWkClassByIds(IDS); err != nil {
		global.GVA_LOG.Error(global.Translate("sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_operation_record.batchDeleteFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("sys_operation_record.batchDeleteSuccess"), c)
	}
}

// UpdateWkClass 更新WkClass
// @Tags WkClass
// @Summary 更新WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkClass true "更新WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wClass/updateWkClass [put]
func (wClassApi *WkClassApi) UpdateWkClass(c *gin.Context) {
	var wClass wkTest.WkClass
	err := c.ShouldBindJSON(&wClass)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wClassService.UpdateWkClass(wClass); err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.updateSuccess"), c)
	}
}

// FindWkClass 用id查询WkClass
// @Tags WkClass
// @Summary 用id查询WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wkTest.WkClass true "用id查询WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wClass/findWkClass [get]
func (wClassApi *WkClassApi) FindWkClass(c *gin.Context) {
	var wClass wkTest.WkClass
	err := c.ShouldBindQuery(&wClass)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewClass, err := wClassService.GetWkClass(wClass.ID); err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
	} else {
		response.OkWithData(gin.H{"rewClass": rewClass}, c)
	}
}

// GetWkClassList 分页获取WkClass列表
// @Tags WkClass
// @Summary 分页获取WkClass列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wkTestReq.WkClassSearch true "分页获取WkClass列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wClass/getWkClassList [get]
func (wClassApi *WkClassApi) GetWkClassList(c *gin.Context) {
	var pageInfo wkTestReq.WkClassSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wClassService.GetWkClassInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, global.Translate("general.getDataSuccess"), c)
	}
}
