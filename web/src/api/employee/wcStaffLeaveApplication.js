import service from '@/utils/request'

// @Tags WcStaffLeaveApplication
// @Summary 创建离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffLeaveApplication true "创建离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffLeaveApplication/createWcStaffLeaveApplication [post]
export const createWcStaffLeaveApplication = (data) => {
  return service({
    url: '/wcStaffLeaveApplication/createWcStaffLeaveApplication',
    method: 'post',
    data
  })
}

// @Tags WcStaffLeaveApplication
// @Summary 删除离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffLeaveApplication true "删除离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffLeaveApplication/deleteWcStaffLeaveApplication [delete]
export const deleteWcStaffLeaveApplication = (params) => {
  return service({
    url: '/wcStaffLeaveApplication/deleteWcStaffLeaveApplication',
    method: 'delete',
    params
  })
}

// @Tags WcStaffLeaveApplication
// @Summary 批量删除离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffLeaveApplication/deleteWcStaffLeaveApplication [delete]
export const deleteWcStaffLeaveApplicationByIds = (params) => {
  return service({
    url: '/wcStaffLeaveApplication/deleteWcStaffLeaveApplicationByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffLeaveApplication
// @Summary 更新离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffLeaveApplication true "更新离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffLeaveApplication/updateWcStaffLeaveApplication [put]
export const updateWcStaffLeaveApplication = (data) => {
  return service({
    url: '/wcStaffLeaveApplication/updateWcStaffLeaveApplication',
    method: 'put',
    data
  })
}

// @Tags WcStaffLeaveApplication
// @Summary 用id查询离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffLeaveApplication true "用id查询离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffLeaveApplication/findWcStaffLeaveApplication [get]
export const findWcStaffLeaveApplication = (params) => {
  return service({
    url: '/wcStaffLeaveApplication/findWcStaffLeaveApplication',
    method: 'get',
    params
  })
}

// @Tags WcStaffLeaveApplication
// @Summary 分页获取离职申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取离职申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffLeaveApplication/getWcStaffLeaveApplicationList [get]
export const getWcStaffLeaveApplicationList = (params) => {
  return service({
    url: '/wcStaffLeaveApplication/getWcStaffLeaveApplicationList',
    method: 'get',
    params
  })
}
