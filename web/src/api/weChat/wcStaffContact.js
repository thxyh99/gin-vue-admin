import service from '@/utils/request'

// @Tags WcStaffContact
// @Summary 创建紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffContact true "创建紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffContact/createWcStaffContact [post]
export const createWcStaffContact = (data) => {
  return service({
    url: '/wcStaffContact/createWcStaffContact',
    method: 'post',
    data
  })
}

// @Tags WcStaffContact
// @Summary 删除紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffContact true "删除紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffContact/deleteWcStaffContact [delete]
export const deleteWcStaffContact = (params) => {
  return service({
    url: '/wcStaffContact/deleteWcStaffContact',
    method: 'delete',
    params
  })
}

// @Tags WcStaffContact
// @Summary 批量删除紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffContact/deleteWcStaffContact [delete]
export const deleteWcStaffContactByIds = (params) => {
  return service({
    url: '/wcStaffContact/deleteWcStaffContactByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffContact
// @Summary 更新紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffContact true "更新紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffContact/updateWcStaffContact [put]
export const updateWcStaffContact = (data) => {
  return service({
    url: '/wcStaffContact/updateWcStaffContact',
    method: 'put',
    data
  })
}

// @Tags WcStaffContact
// @Summary 用id查询紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffContact true "用id查询紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffContact/findWcStaffContact [get]
export const findWcStaffContact = (params) => {
  return service({
    url: '/wcStaffContact/findWcStaffContact',
    method: 'get',
    params
  })
}

// @Tags WcStaffContact
// @Summary 分页获取紧急联系人列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取紧急联系人列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffContact/getWcStaffContactList [get]
export const getWcStaffContactList = (params) => {
  return service({
    url: '/wcStaffContact/getWcStaffContactList',
    method: 'get',
    params
  })
}
