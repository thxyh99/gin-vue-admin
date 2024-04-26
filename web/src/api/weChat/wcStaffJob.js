import service from '@/utils/request'

// @Tags WcStaffJob
// @Summary 创建工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffJob true "创建工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffJob/createWcStaffJob [post]
export const createWcStaffJob = (data) => {
  return service({
    url: '/wcStaffJob/createWcStaffJob',
    method: 'post',
    data
  })
}

// @Tags WcStaffJob
// @Summary 删除工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffJob true "删除工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffJob/deleteWcStaffJob [delete]
export const deleteWcStaffJob = (params) => {
  return service({
    url: '/wcStaffJob/deleteWcStaffJob',
    method: 'delete',
    params
  })
}

// @Tags WcStaffJob
// @Summary 批量删除工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffJob/deleteWcStaffJob [delete]
export const deleteWcStaffJobByIds = (params) => {
  return service({
    url: '/wcStaffJob/deleteWcStaffJobByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffJob
// @Summary 更新工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffJob true "更新工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffJob/updateWcStaffJob [put]
export const updateWcStaffJob = (data) => {
  return service({
    url: '/wcStaffJob/updateWcStaffJob',
    method: 'put',
    data
  })
}

// @Tags WcStaffJob
// @Summary 用id查询工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffJob true "用id查询工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffJob/findWcStaffJob [get]
export const findWcStaffJob = (params) => {
  return service({
    url: '/wcStaffJob/findWcStaffJob',
    method: 'get',
    params
  })
}

// @Tags WcStaffJob
// @Summary 分页获取工作信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工作信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffJob/getWcStaffJobList [get]
export const getWcStaffJobList = (params) => {
  return service({
    url: '/wcStaffJob/getWcStaffJobList',
    method: 'get',
    params
  })
}
