import service from '@/utils/request'

// @Tags WcStaffTransferApplication
// @Summary 创建调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffTransferApplication true "创建调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffTransferApplication/createWcStaffTransferApplication [post]
export const createWcStaffTransferApplication = (data) => {
  return service({
    url: '/wcStaffTransferApplication/createWcStaffTransferApplication',
    method: 'post',
    data
  })
}

// @Tags WcStaffTransferApplication
// @Summary 删除调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffTransferApplication true "删除调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffTransferApplication/deleteWcStaffTransferApplication [delete]
export const deleteWcStaffTransferApplication = (params) => {
  return service({
    url: '/wcStaffTransferApplication/deleteWcStaffTransferApplication',
    method: 'delete',
    params
  })
}

// @Tags WcStaffTransferApplication
// @Summary 批量删除调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffTransferApplication/deleteWcStaffTransferApplication [delete]
export const deleteWcStaffTransferApplicationByIds = (params) => {
  return service({
    url: '/wcStaffTransferApplication/deleteWcStaffTransferApplicationByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffTransferApplication
// @Summary 更新调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffTransferApplication true "更新调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffTransferApplication/updateWcStaffTransferApplication [put]
export const updateWcStaffTransferApplication = (data) => {
  return service({
    url: '/wcStaffTransferApplication/updateWcStaffTransferApplication',
    method: 'put',
    data
  })
}

// @Tags WcStaffTransferApplication
// @Summary 用id查询调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffTransferApplication true "用id查询调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffTransferApplication/findWcStaffTransferApplication [get]
export const findWcStaffTransferApplication = (params) => {
  return service({
    url: '/wcStaffTransferApplication/findWcStaffTransferApplication',
    method: 'get',
    params
  })
}

// @Tags WcStaffTransferApplication
// @Summary 分页获取调动申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取调动申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffTransferApplication/getWcStaffTransferApplicationList [get]
export const getWcStaffTransferApplicationList = (params) => {
  return service({
    url: '/wcStaffTransferApplication/getWcStaffTransferApplicationList',
    method: 'get',
    params
  })
}
