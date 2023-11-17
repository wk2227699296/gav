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

type WkStudentApi struct {
}

var wkStudentService = service.ServiceGroupApp.WkTestServiceGroup.WkStudentService

// CreateWkStudent 创建wkStudent
// @Tags WkStudent
// @Summary 创建wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkStudent true "创建wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":global.Translate("general.getDataSuccess")}"
// @Router /wkStudent/createWkStudent [post]
func (wkStudentApi *WkStudentApi) CreateWkStudent(c *gin.Context) {
	var wkStudent wkTest.WkStudent
	err := c.ShouldBindJSON(&wkStudent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wkStudentService.CreateWkStudent(&wkStudent); err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccss"), c)
	}
}

// DeleteWkStudent 删除wkStudent
// @Tags WkStudent
// @Summary 删除wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkStudent true "删除wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wkStudent/deleteWkStudent [delete]
func (wkStudentApi *WkStudentApi) DeleteWkStudent(c *gin.Context) {
	var wkStudent wkTest.WkStudent
	err := c.ShouldBindJSON(&wkStudent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wkStudentService.DeleteWkStudent(wkStudent); err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deletFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
	}
}

// DeleteWkStudentByIds 批量删除wkStudent
// @Tags WkStudent
// @Summary 批量删除wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wkStudent/deleteWkStudentByIds [delete]
func (wkStudentApi *WkStudentApi) DeleteWkStudentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wkStudentService.DeleteWkStudentByIds(IDS); err != nil {
		global.GVA_LOG.Error(global.Translate("sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_operation_record.batchDeleteFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("sys_operation_record.batchDeleteSuccess"), c)
	}
}

// UpdateWkStudent 更新wkStudent
// @Tags WkStudent
// @Summary 更新wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wkTest.WkStudent true "更新wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wkStudent/updateWkStudent [put]
func (wkStudentApi *WkStudentApi) UpdateWkStudent(c *gin.Context) {
	var wkStudent wkTest.WkStudent
	err := c.ShouldBindJSON(&wkStudent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wkStudentService.UpdateWkStudent(wkStudent); err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.updateSuccess"), c)
	}
}

// FindWkStudent 用id查询wkStudent
// @Tags WkStudent
// @Summary 用id查询wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wkTest.WkStudent true "用id查询wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wkStudent/findWkStudent [get]
func (wkStudentApi *WkStudentApi) FindWkStudent(c *gin.Context) {
	var wkStudent wkTest.WkStudent
	err := c.ShouldBindQuery(&wkStudent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewkStudent, err := wkStudentService.GetWkStudent(wkStudent.ID); err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
	} else {
		response.OkWithData(gin.H{"rewkStudent": rewkStudent}, c)
	}
}

// GetWkStudentList 分页获取wkStudent列表
// @Tags WkStudent
// @Summary 分页获取wkStudent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wkTestReq.WkStudentSearch true "分页获取wkStudent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wkStudent/getWkStudentList [get]
func (wkStudentApi *WkStudentApi) GetWkStudentList(c *gin.Context) {
	var pageInfo wkTestReq.WkStudentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wkStudentService.GetWkStudentInfoList(pageInfo); err != nil {
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
