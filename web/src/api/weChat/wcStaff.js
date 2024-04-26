import service from '@/utils/request'

// @Tags WcStaff
// @Summary 创建账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaff true "创建账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaff/createWcStaff [post]
export const createWcStaff = (data) => {
  return service({
    url: '/wcStaff/createWcStaff',
    method: 'post',
    data
  })
}

// @Tags WcStaff
// @Summary 删除账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaff true "删除账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaff/deleteWcStaff [delete]
export const deleteWcStaff = (params) => {
  return service({
    url: '/wcStaff/deleteWcStaff',
    method: 'delete',
    params
  })
}

// @Tags WcStaff
// @Summary 批量删除账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaff/deleteWcStaff [delete]
export const deleteWcStaffByIds = (params) => {
  return service({
    url: '/wcStaff/deleteWcStaffByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaff
// @Summary 更新账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaff true "更新账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaff/updateWcStaff [put]
export const updateWcStaff = (data) => {
  return service({
    url: '/wcStaff/updateWcStaff',
    method: 'put',
    data
  })
}

// @Tags WcStaff
// @Summary 用id查询账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaff true "用id查询账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaff/findWcStaff [get]
export const findWcStaff = (params) => {
  return service({
    url: '/wcStaff/findWcStaff',
    method: 'get',
    params
  })
}

// @Tags WcStaff
// @Summary 分页获取账号信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取账号信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaff/getWcStaffList [get]
export const getWcStaffList = (params) => {
  return service({
    url: '/wcStaff/getWcStaffList',
    method: 'get',
    params
  })
}
