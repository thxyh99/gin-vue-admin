import service from '@/utils/request'

// @Tags WcSalaryTemplate
// @Summary 创建工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalaryTemplate true "创建工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcSalaryTemplate/createWcSalaryTemplate [post]
export const createWcSalaryTemplate = (data) => {
  return service({
    url: '/wcSalaryTemplate/createWcSalaryTemplate',
    method: 'post',
    data
  })
}

// @Tags WcSalaryTemplate
// @Summary 删除工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalaryTemplate true "删除工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalaryTemplate/deleteWcSalaryTemplate [delete]
export const deleteWcSalaryTemplate = (params) => {
  return service({
    url: '/wcSalaryTemplate/deleteWcSalaryTemplate',
    method: 'delete',
    params
  })
}

// @Tags WcSalaryTemplate
// @Summary 批量删除工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalaryTemplate/deleteWcSalaryTemplate [delete]
export const deleteWcSalaryTemplateByIds = (params) => {
  return service({
    url: '/wcSalaryTemplate/deleteWcSalaryTemplateByIds',
    method: 'delete',
    params
  })
}

// @Tags WcSalaryTemplate
// @Summary 更新工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcSalaryTemplate true "更新工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcSalaryTemplate/updateWcSalaryTemplate [put]
export const updateWcSalaryTemplate = (data) => {
  return service({
    url: '/wcSalaryTemplate/updateWcSalaryTemplate',
    method: 'put',
    data
  })
}

// @Tags WcSalaryTemplate
// @Summary 用id查询工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcSalaryTemplate true "用id查询工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcSalaryTemplate/findWcSalaryTemplate [get]
export const findWcSalaryTemplate = (params) => {
  return service({
    url: '/wcSalaryTemplate/findWcSalaryTemplate',
    method: 'get',
    params
  })
}

// @Tags WcSalaryTemplate
// @Summary 分页获取工资单模版列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工资单模版列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalaryTemplate/getWcSalaryTemplateList [get]
export const getWcSalaryTemplateList = (params) => {
  return service({
    url: '/wcSalaryTemplate/getWcSalaryTemplateList',
    method: 'get',
    params
  })
}
