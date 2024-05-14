import service from '@/utils/request'

// @Tags WcStaffAgreement
// @Summary 创建合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAgreement true "创建合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffAgreement/createWcStaffAgreement [post]
export const createWcStaffAgreement = (data) => {
  return service({
    url: '/wcStaffAgreement/createWcStaffAgreement',
    method: 'post',
    data
  })
}

// @Tags WcStaffAgreement
// @Summary 删除合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAgreement true "删除合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAgreement/deleteWcStaffAgreement [delete]
export const deleteWcStaffAgreement = (params) => {
  return service({
    url: '/wcStaffAgreement/deleteWcStaffAgreement',
    method: 'delete',
    params
  })
}

// @Tags WcStaffAgreement
// @Summary 批量删除合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAgreement/deleteWcStaffAgreement [delete]
export const deleteWcStaffAgreementByIds = (params) => {
  return service({
    url: '/wcStaffAgreement/deleteWcStaffAgreementByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffAgreement
// @Summary 更新合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffAgreement true "更新合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffAgreement/updateWcStaffAgreement [put]
export const updateWcStaffAgreement = (data) => {
  return service({
    url: '/wcStaffAgreement/updateWcStaffAgreement',
    method: 'put',
    data
  })
}

// @Tags WcStaffAgreement
// @Summary 用id查询合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffAgreement true "用id查询合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffAgreement/findWcStaffAgreement [get]
export const findWcStaffAgreement = (params) => {
  return service({
    url: '/wcStaffAgreement/findWcStaffAgreement',
    method: 'get',
    params
  })
}

// @Tags WcStaffAgreement
// @Summary 分页获取合同信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取合同信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAgreement/getWcStaffAgreementList [get]
export const getWcStaffAgreementList = (params) => {
  return service({
    url: '/wcStaffAgreement/getWcStaffAgreementList',
    method: 'get',
    params
  })
}
