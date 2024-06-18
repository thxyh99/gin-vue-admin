import service from '@/utils/request'

// @Tags WcThirdOaPush
// @Summary 创建OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcThirdOaPush true "创建OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcThirdOaPush/createWcThirdOaPush [post]
export const createWcThirdOaPush = (data) => {
  return service({
    url: '/wcThirdOaPush/createWcThirdOaPush',
    method: 'post',
    data
  })
}

// @Tags WcThirdOaPush
// @Summary 删除OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcThirdOaPush true "删除OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcThirdOaPush/deleteWcThirdOaPush [delete]
export const deleteWcThirdOaPush = (params) => {
  return service({
    url: '/wcThirdOaPush/deleteWcThirdOaPush',
    method: 'delete',
    params
  })
}

// @Tags WcThirdOaPush
// @Summary 批量删除OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcThirdOaPush/deleteWcThirdOaPush [delete]
export const deleteWcThirdOaPushByIds = (params) => {
  return service({
    url: '/wcThirdOaPush/deleteWcThirdOaPushByIds',
    method: 'delete',
    params
  })
}

// @Tags WcThirdOaPush
// @Summary 更新OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcThirdOaPush true "更新OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcThirdOaPush/updateWcThirdOaPush [put]
export const updateWcThirdOaPush = (data) => {
  return service({
    url: '/wcThirdOaPush/updateWcThirdOaPush',
    method: 'put',
    data
  })
}

// @Tags WcThirdOaPush
// @Summary 用id查询OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcThirdOaPush true "用id查询OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcThirdOaPush/findWcThirdOaPush [get]
export const findWcThirdOaPush = (params) => {
  return service({
    url: '/wcThirdOaPush/findWcThirdOaPush',
    method: 'get',
    params
  })
}

// @Tags WcThirdOaPush
// @Summary 分页获取OA回调日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OA回调日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcThirdOaPush/getWcThirdOaPushList [get]
export const getWcThirdOaPushList = (params) => {
  return service({
    url: '/wcThirdOaPush/getWcThirdOaPushList',
    method: 'get',
    params
  })
}
