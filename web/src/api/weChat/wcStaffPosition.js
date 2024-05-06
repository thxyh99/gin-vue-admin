import service from '@/utils/request'

// @Tags WcStaffPosition
// @Summary 创建员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffPosition true "创建员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffPosition/createWcStaffPosition [post]
export const createWcStaffPosition = (data) => {
  return service({
    url: '/wcStaffPosition/createWcStaffPosition',
    method: 'post',
    data
  })
}

// @Tags WcStaffPosition
// @Summary 删除员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffPosition true "删除员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffPosition/deleteWcStaffPosition [delete]
export const deleteWcStaffPosition = (params) => {
  return service({
    url: '/wcStaffPosition/deleteWcStaffPosition',
    method: 'delete',
    params
  })
}

// @Tags WcStaffPosition
// @Summary 批量删除员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffPosition/deleteWcStaffPosition [delete]
export const deleteWcStaffPositionByIds = (params) => {
  return service({
    url: '/wcStaffPosition/deleteWcStaffPositionByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffPosition
// @Summary 更新员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffPosition true "更新员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffPosition/updateWcStaffPosition [put]
export const updateWcStaffPosition = (data) => {
  return service({
    url: '/wcStaffPosition/updateWcStaffPosition',
    method: 'put',
    data
  })
}

// @Tags WcStaffPosition
// @Summary 用id查询员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffPosition true "用id查询员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffPosition/findWcStaffPosition [get]
export const findWcStaffPosition = (params) => {
  return service({
    url: '/wcStaffPosition/findWcStaffPosition',
    method: 'get',
    params
  })
}

// @Tags WcStaffPosition
// @Summary 分页获取员工职位管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取员工职位管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffPosition/getWcStaffPositionList [get]
export const getWcStaffPositionList = (params) => {
  return service({
    url: '/wcStaffPosition/getWcStaffPositionList',
    method: 'get',
    params
  })
}
