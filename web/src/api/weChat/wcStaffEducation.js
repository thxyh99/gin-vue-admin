import service from '@/utils/request'

// @Tags WcStaffEducation
// @Summary 创建学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffEducation true "创建学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffEducation/createWcStaffEducation [post]
export const createWcStaffEducation = (data) => {
  return service({
    url: '/wcStaffEducation/createWcStaffEducation',
    method: 'post',
    data
  })
}

// @Tags WcStaffEducation
// @Summary 删除学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffEducation true "删除学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffEducation/deleteWcStaffEducation [delete]
export const deleteWcStaffEducation = (params) => {
  return service({
    url: '/wcStaffEducation/deleteWcStaffEducation',
    method: 'delete',
    params
  })
}

// @Tags WcStaffEducation
// @Summary 批量删除学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffEducation/deleteWcStaffEducation [delete]
export const deleteWcStaffEducationByIds = (params) => {
  return service({
    url: '/wcStaffEducation/deleteWcStaffEducationByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffEducation
// @Summary 更新学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffEducation true "更新学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffEducation/updateWcStaffEducation [put]
export const updateWcStaffEducation = (data) => {
  return service({
    url: '/wcStaffEducation/updateWcStaffEducation',
    method: 'put',
    data
  })
}

// @Tags WcStaffEducation
// @Summary 用id查询学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffEducation true "用id查询学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffEducation/findWcStaffEducation [get]
export const findWcStaffEducation = (params) => {
  return service({
    url: '/wcStaffEducation/findWcStaffEducation',
    method: 'get',
    params
  })
}

// @Tags WcStaffEducation
// @Summary 分页获取学历信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取学历信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffEducation/getWcStaffEducationList [get]
export const getWcStaffEducationList = (params) => {
  return service({
    url: '/wcStaffEducation/getWcStaffEducationList',
    method: 'get',
    params
  })
}
