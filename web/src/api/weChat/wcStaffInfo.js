import service from '@/utils/request'

// @Tags WcStaffInfo
// @Summary 创建个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffInfo true "创建个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffInfo/createWcStaffInfo [post]
export const createWcStaffInfo = (data) => {
  return service({
    url: '/wcStaffInfo/createWcStaffInfo',
    method: 'post',
    data
  })
}

// @Tags WcStaffInfo
// @Summary 删除个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffInfo true "删除个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffInfo/deleteWcStaffInfo [delete]
export const deleteWcStaffInfo = (params) => {
  return service({
    url: '/wcStaffInfo/deleteWcStaffInfo',
    method: 'delete',
    params
  })
}

// @Tags WcStaffInfo
// @Summary 批量删除个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffInfo/deleteWcStaffInfo [delete]
export const deleteWcStaffInfoByIds = (params) => {
  return service({
    url: '/wcStaffInfo/deleteWcStaffInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffInfo
// @Summary 更新个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffInfo true "更新个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffInfo/updateWcStaffInfo [put]
export const updateWcStaffInfo = (data) => {
  return service({
    url: '/wcStaffInfo/updateWcStaffInfo',
    method: 'put',
    data
  })
}

// @Tags WcStaffInfo
// @Summary 用id查询个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffInfo true "用id查询个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffInfo/findWcStaffInfo [get]
export const findWcStaffInfo = (params) => {
  return service({
    url: '/wcStaffInfo/findWcStaffInfo',
    method: 'get',
    params
  })
}

// @Tags WcStaffInfo
// @Summary 分页获取个人信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取个人信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffInfo/getWcStaffInfoList [get]
export const getWcStaffInfoList = (params) => {
  return service({
    url: '/wcStaffInfo/getWcStaffInfoList',
    method: 'get',
    params
  })
}
