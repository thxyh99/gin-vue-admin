import service from '@/utils/request'

// @Tags WcStaffAdjustlevelApplication
// @Summary 创建调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAdjustlevelApplication true "创建调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffAdjustlevelApplication/createWcStaffAdjustlevelApplication [post]
export const createWcStaffAdjustlevelApplication = (data) => {
  return service({
    url: '/wcStaffAdjustlevelApplication/createWcStaffAdjustlevelApplication',
    method: 'post',
    data
  })
}

// @Tags WcStaffAdjustlevelApplication
// @Summary 删除调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAdjustlevelApplication true "删除调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAdjustlevelApplication/deleteWcStaffAdjustlevelApplication [delete]
export const deleteWcStaffAdjustlevelApplication = (params) => {
  return service({
    url: '/wcStaffAdjustlevelApplication/deleteWcStaffAdjustlevelApplication',
    method: 'delete',
    params
  })
}

// @Tags WcStaffAdjustlevelApplication
// @Summary 批量删除调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAdjustlevelApplication/deleteWcStaffAdjustlevelApplication [delete]
export const deleteWcStaffAdjustlevelApplicationByIds = (params) => {
  return service({
    url: '/wcStaffAdjustlevelApplication/deleteWcStaffAdjustlevelApplicationByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffAdjustlevelApplication
// @Summary 更新调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAdjustlevelApplication true "更新调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffAdjustlevelApplication/updateWcStaffAdjustlevelApplication [put]
export const updateWcStaffAdjustlevelApplication = (data) => {
  return service({
    url: '/wcStaffAdjustlevelApplication/updateWcStaffAdjustlevelApplication',
    method: 'put',
    data
  })
}

// @Tags WcStaffAdjustlevelApplication
// @Summary 用id查询调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffAdjustlevelApplication true "用id查询调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffAdjustlevelApplication/findWcStaffAdjustlevelApplication [get]
export const findWcStaffAdjustlevelApplication = (params) => {
  return service({
    url: '/wcStaffAdjustlevelApplication/findWcStaffAdjustlevelApplication',
    method: 'get',
    params
  })
}

// @Tags WcStaffAdjustlevelApplication
// @Summary 分页获取调级申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取调级申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAdjustlevelApplication/getWcStaffAdjustlevelApplicationList [get]
export const getWcStaffAdjustlevelApplicationList = (params) => {
  return service({
    url: '/wcStaffAdjustlevelApplication/getWcStaffAdjustlevelApplicationList',
    method: 'get',
    params
  })
}
