import service from '@/utils/request'

// @Tags WcStaffDepartment
// @Summary 创建员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffDepartment true "创建员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffDepartment/createWcStaffDepartment [post]
export const createWcStaffDepartment = (data) => {
  return service({
    url: '/wcStaffDepartment/createWcStaffDepartment',
    method: 'post',
    data
  })
}

// @Tags WcStaffDepartment
// @Summary 删除员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffDepartment true "删除员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffDepartment/deleteWcStaffDepartment [delete]
export const deleteWcStaffDepartment = (params) => {
  return service({
    url: '/wcStaffDepartment/deleteWcStaffDepartment',
    method: 'delete',
    params
  })
}

// @Tags WcStaffDepartment
// @Summary 批量删除员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffDepartment/deleteWcStaffDepartment [delete]
export const deleteWcStaffDepartmentByIds = (params) => {
  return service({
    url: '/wcStaffDepartment/deleteWcStaffDepartmentByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffDepartment
// @Summary 更新员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffDepartment true "更新员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffDepartment/updateWcStaffDepartment [put]
export const updateWcStaffDepartment = (data) => {
  return service({
    url: '/wcStaffDepartment/updateWcStaffDepartment',
    method: 'put',
    data
  })
}

// @Tags WcStaffDepartment
// @Summary 用id查询员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffDepartment true "用id查询员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffDepartment/findWcStaffDepartment [get]
export const findWcStaffDepartment = (params) => {
  return service({
    url: '/wcStaffDepartment/findWcStaffDepartment',
    method: 'get',
    params
  })
}

// @Tags WcStaffDepartment
// @Summary 分页获取员工部门管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取员工部门管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffDepartment/getWcStaffDepartmentList [get]
export const getWcStaffDepartmentList = (params) => {
  return service({
    url: '/wcStaffDepartment/getWcStaffDepartmentList',
    method: 'get',
    params
  })
}
