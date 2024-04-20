import service from '@/utils/request'

// @Tags WcDepartment
// @Summary 创建wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcDepartment true "创建wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcDepartment/createWcDepartment [post]
export const createWcDepartment = (data) => {
  return service({
    url: '/wcDepartment/createWcDepartment',
    method: 'post',
    data
  })
}

// @Tags WcDepartment
// @Summary 删除wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcDepartment true "删除wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcDepartment/deleteWcDepartment [delete]
export const deleteWcDepartment = (params) => {
  return service({
    url: '/wcDepartment/deleteWcDepartment',
    method: 'delete',
    params
  })
}

// @Tags WcDepartment
// @Summary 批量删除wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcDepartment/deleteWcDepartment [delete]
export const deleteWcDepartmentByIds = (params) => {
  return service({
    url: '/wcDepartment/deleteWcDepartmentByIds',
    method: 'delete',
    params
  })
}

// @Tags WcDepartment
// @Summary 更新wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcDepartment true "更新wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcDepartment/updateWcDepartment [put]
export const updateWcDepartment = (data) => {
  return service({
    url: '/wcDepartment/updateWcDepartment',
    method: 'put',
    data
  })
}

// @Tags WcDepartment
// @Summary 用id查询wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcDepartment true "用id查询wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcDepartment/findWcDepartment [get]
export const findWcDepartment = (params) => {
  return service({
    url: '/wcDepartment/findWcDepartment',
    method: 'get',
    params
  })
}

// @Tags WcDepartment
// @Summary 分页获取wcDepartment表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wcDepartment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcDepartment/getWcDepartmentList [get]
export const getWcDepartmentList = (params) => {
  return service({
    url: '/wcDepartment/getWcDepartmentList',
    method: 'get',
    params
  })
}
