import service from '@/utils/request'

// @Tags WcSalary
// @Summary 创建工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalary true "创建工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcSalary/createWcSalary [post]
export const createWcSalary = (data) => {
  return service({
    url: '/wcSalary/createWcSalary',
    method: 'post',
    data
  })
}

// @Tags WcSalary
// @Summary 删除工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalary true "删除工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalary/deleteWcSalary [delete]
export const deleteWcSalary = (params) => {
  return service({
    url: '/wcSalary/deleteWcSalary',
    method: 'delete',
    params
  })
}

// @Tags WcSalary
// @Summary 批量删除工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalary/deleteWcSalary [delete]
export const deleteWcSalaryByIds = (params) => {
  return service({
    url: '/wcSalary/deleteWcSalaryByIds',
    method: 'delete',
    params
  })
}

// @Tags WcSalary
// @Summary 更新工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalary true "更新工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcSalary/updateWcSalary [put]
export const updateWcSalary = (data) => {
  return service({
    url: '/wcSalary/updateWcSalary',
    method: 'put',
    data
  })
}

// @Tags WcSalary
// @Summary 用id查询工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcSalary true "用id查询工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcSalary/findWcSalary [get]
export const findWcSalary = (params) => {
  return service({
    url: '/wcSalary/findWcSalary',
    method: 'get',
    params
  })
}

// @Tags WcSalary
// @Summary 分页获取工资单模版列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工资单模版列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalary/getWcSalaryList [get]
export const getWcSalaryList = (params) => {
  return service({
    url: '/wcSalary/getWcSalaryList',
    method: 'get',
    params
  })
}
