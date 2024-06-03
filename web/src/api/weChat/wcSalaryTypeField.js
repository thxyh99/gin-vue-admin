import service from '@/utils/request'

// @Tags WcSalaryTypeField
// @Summary 创建工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalaryTypeField true "创建工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcSalaryTypeField/createWcSalaryTypeField [post]
export const createWcSalaryTypeField = (data) => {
  return service({
    url: '/wcSalaryTypeField/createWcSalaryTypeField',
    method: 'post',
    data
  })
}

// @Tags WcSalaryTypeField
// @Summary 删除工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalaryTypeField true "删除工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalaryTypeField/deleteWcSalaryTypeField [delete]
export const deleteWcSalaryTypeField = (params) => {
  return service({
    url: '/wcSalaryTypeField/deleteWcSalaryTypeField',
    method: 'delete',
    params
  })
}

// @Tags WcSalaryTypeField
// @Summary 批量删除工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalaryTypeField/deleteWcSalaryTypeField [delete]
export const deleteWcSalaryTypeFieldByIds = (params) => {
  return service({
    url: '/wcSalaryTypeField/deleteWcSalaryTypeFieldByIds',
    method: 'delete',
    params
  })
}

// @Tags WcSalaryTypeField
// @Summary 更新工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalaryTypeField true "更新工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcSalaryTypeField/updateWcSalaryTypeField [put]
export const updateWcSalaryTypeField = (data) => {
  return service({
    url: '/wcSalaryTypeField/updateWcSalaryTypeField',
    method: 'put',
    data
  })
}

// @Tags WcSalaryTypeField
// @Summary 用id查询工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcSalaryTypeField true "用id查询工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcSalaryTypeField/findWcSalaryTypeField [get]
export const findWcSalaryTypeField = (params) => {
  return service({
    url: '/wcSalaryTypeField/findWcSalaryTypeField',
    method: 'get',
    params
  })
}

// @Tags WcSalaryTypeField
// @Summary 分页获取工资类型款项列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工资类型款项列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalaryTypeField/getWcSalaryTypeFieldList [get]
export const getWcSalaryTypeFieldList = (params) => {
  return service({
    url: '/wcSalaryTypeField/getWcSalaryTypeFieldList',
    method: 'get',
    params
  })
}
