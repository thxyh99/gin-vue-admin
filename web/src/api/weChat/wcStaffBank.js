import service from '@/utils/request'

// @Tags WcStaffBank
// @Summary 创建银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffBank true "创建银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffBank/createWcStaffBank [post]
export const createWcStaffBank = (data) => {
  return service({
    url: '/wcStaffBank/createWcStaffBank',
    method: 'post',
    data
  })
}

// @Tags WcStaffBank
// @Summary 删除银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffBank true "删除银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffBank/deleteWcStaffBank [delete]
export const deleteWcStaffBank = (params) => {
  return service({
    url: '/wcStaffBank/deleteWcStaffBank',
    method: 'delete',
    params
  })
}

// @Tags WcStaffBank
// @Summary 批量删除银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffBank/deleteWcStaffBank [delete]
export const deleteWcStaffBankByIds = (params) => {
  return service({
    url: '/wcStaffBank/deleteWcStaffBankByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffBank
// @Summary 更新银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffBank true "更新银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffBank/updateWcStaffBank [put]
export const updateWcStaffBank = (data) => {
  return service({
    url: '/wcStaffBank/updateWcStaffBank',
    method: 'put',
    data
  })
}

// @Tags WcStaffBank
// @Summary 用id查询银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffBank true "用id查询银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffBank/findWcStaffBank [get]
export const findWcStaffBank = (params) => {
  return service({
    url: '/wcStaffBank/findWcStaffBank',
    method: 'get',
    params
  })
}

// @Tags WcStaffBank
// @Summary 分页获取银行卡信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取银行卡信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffBank/getWcStaffBankList [get]
export const getWcStaffBankList = (params) => {
  return service({
    url: '/wcStaffBank/getWcStaffBankList',
    method: 'get',
    params
  })
}
