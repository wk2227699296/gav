import service from '@/utils/request'

// @Tags WkClass
// @Summary 创建WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkClass true "创建WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wClass/createWkClass [post]
export const createWkClass = (data) => {
  return service({
    url: '/wClass/createWkClass',
    method: 'post',
    data
  })
}

// @Tags WkClass
// @Summary 删除WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkClass true "删除WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wClass/deleteWkClass [delete]
export const deleteWkClass = (data) => {
  return service({
    url: '/wClass/deleteWkClass',
    method: 'delete',
    data
  })
}

// @Tags WkClass
// @Summary 批量删除WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wClass/deleteWkClass [delete]
export const deleteWkClassByIds = (data) => {
  return service({
    url: '/wClass/deleteWkClassByIds',
    method: 'delete',
    data
  })
}

// @Tags WkClass
// @Summary 更新WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WkClass true "更新WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wClass/updateWkClass [put]
export const updateWkClass = (data) => {
  return service({
    url: '/wClass/updateWkClass',
    method: 'put',
    data
  })
}

// @Tags WkClass
// @Summary 用id查询WkClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WkClass true "用id查询WkClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wClass/findWkClass [get]
export const findWkClass = (params) => {
  return service({
    url: '/wClass/findWkClass',
    method: 'get',
    params
  })
}

// @Tags WkClass
// @Summary 分页获取WkClass列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WkClass列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wClass/getWkClassList [get]
export const getWkClassList = (params) => {
  return service({
    url: '/wClass/getWkClassList',
    method: 'get',
    params
  })
}
