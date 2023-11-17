import service from '@/utils/request'

// @Tags WkInfo
// @Summary 创建WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkInfo true "创建WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wkinfo/createWkInfo [post]
export const createWkInfo = (data) => {
  return service({
    url: '/wkinfo/createWkInfo',
    method: 'post',
    data
  })
}

// @Tags WkInfo
// @Summary 删除WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkInfo true "删除WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wkinfo/deleteWkInfo [delete]
export const deleteWkInfo = (data) => {
  return service({
    url: '/wkinfo/deleteWkInfo',
    method: 'delete',
    data
  })
}

// @Tags WkInfo
// @Summary 批量删除WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wkinfo/deleteWkInfo [delete]
export const deleteWkInfoByIds = (data) => {
  return service({
    url: '/wkinfo/deleteWkInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags WkInfo
// @Summary 更新WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkInfo true "更新WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wkinfo/updateWkInfo [put]
export const updateWkInfo = (data) => {
  return service({
    url: '/wkinfo/updateWkInfo',
    method: 'put',
    data
  })
}

// @Tags WkInfo
// @Summary 用id查询WkInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WkInfo true "用id查询WkInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wkinfo/findWkInfo [get]
export const findWkInfo = (params) => {
  return service({
    url: '/wkinfo/findWkInfo',
    method: 'get',
    params
  })
}

// @Tags WkInfo
// @Summary 分页获取WkInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WkInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wkinfo/getWkInfoList [get]
export const getWkInfoList = (params) => {
  return service({
    url: '/wkinfo/getWkInfoList',
    method: 'get',
    params
  })
}
