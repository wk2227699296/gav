import service from '@/utils/request'

// @Tags WkStudent
// @Summary 创建wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkStudent true "创建wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wkStudent/createWkStudent [post]
export const createWkStudent = (data) => {
  return service({
    url: '/wkStudent/createWkStudent',
    method: 'post',
    data
  })
}

// @Tags WkStudent
// @Summary 删除wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkStudent true "删除wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wkStudent/deleteWkStudent [delete]
export const deleteWkStudent = (data) => {
  return service({
    url: '/wkStudent/deleteWkStudent',
    method: 'delete',
    data
  })
}

// @Tags WkStudent
// @Summary 批量删除wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wkStudent/deleteWkStudent [delete]
export const deleteWkStudentByIds = (data) => {
  return service({
    url: '/wkStudent/deleteWkStudentByIds',
    method: 'delete',
    data
  })
}

// @Tags WkStudent
// @Summary 更新wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkStudent true "更新wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wkStudent/updateWkStudent [put]
export const updateWkStudent = (data) => {
  return service({
    url: '/wkStudent/updateWkStudent',
    method: 'put',
    data
  })
}

// @Tags WkStudent
// @Summary 用id查询wkStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WkStudent true "用id查询wkStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wkStudent/findWkStudent [get]
export const findWkStudent = (params) => {
  return service({
    url: '/wkStudent/findWkStudent',
    method: 'get',
    params
  })
}

// @Tags WkStudent
// @Summary 分页获取wkStudent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wkStudent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wkStudent/getWkStudentList [get]
export const getWkStudentList = (params) => {
  return service({
    url: '/wkStudent/getWkStudentList',
    method: 'get',
    params
  })
}
