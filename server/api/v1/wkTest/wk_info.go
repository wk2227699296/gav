package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wkTest"
	wkTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/wkTest/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WkInfoApi struct {
}

var wkinfoService = service.ServiceGroupApp.WkTestServiceGroup.WkInfoService

// CreateWkInfo 创建WkInfo
// @Tags WkInfo
// @Summary 创建WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkInfo true "创建WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":global.Translate("general.getDataSuccess")}"
// @Router /wkinfo/createWkInfo [post]
func (wkinfoApi *WkInfoApi) CreateWkInfo(c *gin.Context) {
	var wkinfo wkTest.WkInfo
	err := c.ShouldBindJSON(&wkinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wkinfo.CreatedBy = utils.GetUserID(c)
	if err := wkinfoService.CreateWkInfo(&wkinfo); err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccss"), c)
	}
}

// DeleteWkInfo 删除WkInfo
// @Tags WkInfo
// @Summary 删除WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkInfo true "删除WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wkinfo/deleteWkInfo [delete]
func (wkinfoApi *WkInfoApi) DeleteWkInfo(c *gin.Context) {
	var wkinfo wkTest.WkInfo
	err := c.ShouldBindJSON(&wkinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wkinfo.DeletedBy = utils.GetUserID(c)
	if err := wkinfoService.DeleteWkInfo(wkinfo); err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deletFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
	}
}

// DeleteWkInfoByIds 批量删除WkInfo
// @Tags WkInfo
// @Summary 批量删除WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wkinfo/deleteWkInfoByIds [delete]
func (wkinfoApi *WkInfoApi) DeleteWkInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := wkinfoService.DeleteWkInfoByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error(global.Translate("sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_operation_record.batchDeleteFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("sys_operation_record.batchDeleteSuccess"), c)
	}
}

// UpdateWkInfo 更新WkInfo
// @Tags WkInfo
// @Summary 更新WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkInfo true "更新WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wkinfo/updateWkInfo [put]
func (wkinfoApi *WkInfoApi) UpdateWkInfo(c *gin.Context) {
	var wkinfo wkTest.WkInfo
	err := c.ShouldBindJSON(&wkinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wkinfo.UpdatedBy = utils.GetUserID(c)
	if err := wkinfoService.UpdateWkInfo(wkinfo); err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.updateSuccess"), c)
	}
}

// FindWkInfo 用id查询WkInfo
// @Tags WkInfo
// @Summary 用id查询WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wkTest.WkInfo true "用id查询WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wkinfo/findWkInfo [get]
func (wkinfoApi *WkInfoApi) FindWkInfo(c *gin.Context) {
	var wkinfo wkTest.WkInfo
	err := c.ShouldBindQuery(&wkinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewkinfo, err := wkinfoService.GetWkInfo(wkinfo.ID); err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
	} else {
		response.OkWithData(gin.H{"rewkinfo": rewkinfo}, c)
	}
}

// GetWkInfoList 分页获取WkInfo列表
// @Tags WkInfo
// @Summary 分页获取WkInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wkTestReq.WkInfoSearch true "分页获取WkInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wkinfo/getWkInfoList [get]
func (wkinfoApi *WkInfoApi) GetWkInfoList(c *gin.Context) {
	var pageInfo wkTestReq.WkInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wkinfoService.GetWkInfoInfoList(pageInfo); err != nil {
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
