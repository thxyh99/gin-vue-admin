import service from '@/utils/request'

// @Tags WcStaffSalary
// @Summary 创建工资列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffSalary true "创建工资列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffSalary/createWcStaffSalary [post]
export const createWcStaffSalary = (data) => {
  return service({
    url: '/wcStaffSalary/createWcStaffSalary',
    method: 'post',
    data
  })
}

// @Tags WcStaffSalary
// @Summary 删除工资列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffSalary true "删除工资列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffSalary/deleteWcStaffSalary [delete]
export const deleteWcStaffSalary = (params) => {
  return service({
    url: '/wcStaffSalary/deleteWcStaffSalary',
    method: 'delete',
    params
  })
}

// @Tags WcStaffSalary
// @Summary 批量删除工资列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工资列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffSalary/deleteWcStaffSalary [delete]
export const deleteWcStaffSalaryByIds = (params) => {
  return service({
    url: '/wcStaffSalary/deleteWcStaffSalaryByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffSalary
// @Summary 更新工资列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffSalary true "更新工资列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffSalary/updateWcStaffSalary [put]
export const updateWcStaffSalary = (data) => {
  return service({
    url: '/wcStaffSalary/updateWcStaffSalary',
    method: 'put',
    data
  })
}

// @Tags WcStaffSalary
// @Summary 用id查询工资列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffSalary true "用id查询工资列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffSalary/findWcStaffSalary [get]
export const findWcStaffSalary = (params) => {
  return service({
    url: '/wcStaffSalary/findWcStaffSalary',
    method: 'get',
    params
  })
}

// @Tags WcStaffSalary
// @Summary 分页获取工资列表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工资列表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffSalary/getWcStaffSalaryList [get]
export const getWcStaffSalaryList = (params) => {
  return service({
    url: '/wcStaffSalary/getWcStaffSalaryList',
    method: 'get',
    params
  })
}
