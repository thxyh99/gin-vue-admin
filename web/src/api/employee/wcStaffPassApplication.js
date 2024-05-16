import service from '@/utils/request'

// @Tags WcStaffPassApplication
// @Summary 创建转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffPassApplication true "创建转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffPassApplication/createWcStaffPassApplication [post]
export const createWcStaffPassApplication = (data) => {
  return service({
    url: '/wcStaffPassApplication/createWcStaffPassApplication',
    method: 'post',
    data
  })
}

// @Tags WcStaffPassApplication
// @Summary 删除转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffPassApplication true "删除转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffPassApplication/deleteWcStaffPassApplication [delete]
export const deleteWcStaffPassApplication = (params) => {
  return service({
    url: '/wcStaffPassApplication/deleteWcStaffPassApplication',
    method: 'delete',
    params
  })
}

// @Tags WcStaffPassApplication
// @Summary 批量删除转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffPassApplication/deleteWcStaffPassApplication [delete]
export const deleteWcStaffPassApplicationByIds = (params) => {
  return service({
    url: '/wcStaffPassApplication/deleteWcStaffPassApplicationByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffPassApplication
// @Summary 更新转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffPassApplication true "更新转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffPassApplication/updateWcStaffPassApplication [put]
export const updateWcStaffPassApplication = (data) => {
  return service({
    url: '/wcStaffPassApplication/updateWcStaffPassApplication',
    method: 'put',
    data
  })
}

// @Tags WcStaffPassApplication
// @Summary 用id查询转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffPassApplication true "用id查询转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffPassApplication/findWcStaffPassApplication [get]
export const findWcStaffPassApplication = (params) => {
  return service({
    url: '/wcStaffPassApplication/findWcStaffPassApplication',
    method: 'get',
    params
  })
}

// @Tags WcStaffPassApplication
// @Summary 分页获取转正申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取转正申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffPassApplication/getWcStaffPassApplicationList [get]
export const getWcStaffPassApplicationList = (params) => {
  return service({
    url: '/wcStaffPassApplication/getWcStaffPassApplicationList',
    method: 'get',
    params
  })
}
