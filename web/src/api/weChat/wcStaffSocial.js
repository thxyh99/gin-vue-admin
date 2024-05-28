import service from '@/utils/request'

// @Tags WcStaffSocial
// @Summary 创建社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffSocial true "创建社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffSocial/createWcStaffSocial [post]
export const createWcStaffSocial = (data) => {
  return service({
    url: '/wcStaffSocial/createWcStaffSocial',
    method: 'post',
    data
  })
}

// @Tags WcStaffSocial
// @Summary 删除社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffSocial true "删除社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffSocial/deleteWcStaffSocial [delete]
export const deleteWcStaffSocial = (params) => {
  return service({
    url: '/wcStaffSocial/deleteWcStaffSocial',
    method: 'delete',
    params
  })
}

// @Tags WcStaffSocial
// @Summary 批量删除社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffSocial/deleteWcStaffSocial [delete]
export const deleteWcStaffSocialByIds = (params) => {
  return service({
    url: '/wcStaffSocial/deleteWcStaffSocialByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffSocial
// @Summary 更新社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffSocial true "更新社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffSocial/updateWcStaffSocial [put]
export const updateWcStaffSocial = (data) => {
  return service({
    url: '/wcStaffSocial/updateWcStaffSocial',
    method: 'put',
    data
  })
}

// @Tags WcStaffSocial
// @Summary 用id查询社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffSocial true "用id查询社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffSocial/findWcStaffSocial [get]
export const findWcStaffSocial = (params) => {
  return service({
    url: '/wcStaffSocial/findWcStaffSocial',
    method: 'get',
    params
  })
}

// @Tags WcStaffSocial
// @Summary 分页获取社保公积金管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取社保公积金管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffSocial/getWcStaffSocialList [get]
export const getWcStaffSocialList = (params) => {
  return service({
    url: '/wcStaffSocial/getWcStaffSocialList',
    method: 'get',
    params
  })
}
