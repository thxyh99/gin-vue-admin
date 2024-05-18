import service from '@/utils/request'

// @Tags WcStaffAjustlevelApplication
// @Summary 创建调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAjustlevelApplication true "创建调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffAjustlevelApplication/createWcStaffAjustlevelApplication [post]
export const createWcStaffAjustlevelApplication = (data) => {
  return service({
    url: '/wcStaffAjustlevelApplication/createWcStaffAjustlevelApplication',
    method: 'post',
    data
  })
}

// @Tags WcStaffAjustlevelApplication
// @Summary 删除调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAjustlevelApplication true "删除调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAjustlevelApplication/deleteWcStaffAjustlevelApplication [delete]
export const deleteWcStaffAjustlevelApplication = (params) => {
  return service({
    url: '/wcStaffAjustlevelApplication/deleteWcStaffAjustlevelApplication',
    method: 'delete',
    params
  })
}

// @Tags WcStaffAjustlevelApplication
// @Summary 批量删除调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAjustlevelApplication/deleteWcStaffAjustlevelApplication [delete]
export const deleteWcStaffAjustlevelApplicationByIds = (params) => {
  return service({
    url: '/wcStaffAjustlevelApplication/deleteWcStaffAjustlevelApplicationByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffAjustlevelApplication
// @Summary 更新调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAjustlevelApplication true "更新调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffAjustlevelApplication/updateWcStaffAjustlevelApplication [put]
export const updateWcStaffAjustlevelApplication = (data) => {
  return service({
    url: '/wcStaffAjustlevelApplication/updateWcStaffAjustlevelApplication',
    method: 'put',
    data
  })
}

// @Tags WcStaffAjustlevelApplication
// @Summary 用id查询调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffAjustlevelApplication true "用id查询调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffAjustlevelApplication/findWcStaffAjustlevelApplication [get]
export const findWcStaffAjustlevelApplication = (params) => {
  return service({
    url: '/wcStaffAjustlevelApplication/findWcStaffAjustlevelApplication',
    method: 'get',
    params
  })
}

// @Tags WcStaffAjustlevelApplication
// @Summary 分页获取调级申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取调级申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAjustlevelApplication/getWcStaffAjustlevelApplicationList [get]
export const getWcStaffAjustlevelApplicationList = (params) => {
  return service({
    url: '/wcStaffAjustlevelApplication/getWcStaffAjustlevelApplicationList',
    method: 'get',
    params
  })
}
