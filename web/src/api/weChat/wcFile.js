import service from '@/utils/request'

// @Tags WcFile
// @Summary 创建文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcFile true "创建文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcFile/createWcFile [post]
export const createWcFile = (data) => {
  return service({
    url: '/wcFile/createWcFile',
    method: 'post',
    data
  })
}

// @Tags WcFile
// @Summary 删除文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcFile true "删除文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcFile/deleteWcFile [delete]
export const deleteWcFile = (params) => {
  return service({
    url: '/wcFile/deleteWcFile',
    method: 'delete',
    params
  })
}

// @Tags WcFile
// @Summary 批量删除文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcFile/deleteWcFile [delete]
export const deleteWcFileByIds = (params) => {
  return service({
    url: '/wcFile/deleteWcFileByIds',
    method: 'delete',
    params
  })
}

// @Tags WcFile
// @Summary 更新文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcFile true "更新文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcFile/updateWcFile [put]
export const updateWcFile = (data) => {
  return service({
    url: '/wcFile/updateWcFile',
    method: 'put',
    data
  })
}

// @Tags WcFile
// @Summary 用id查询文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcFile true "用id查询文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcFile/findWcFile [get]
export const findWcFile = (params) => {
  return service({
    url: '/wcFile/findWcFile',
    method: 'get',
    params
  })
}

// @Tags WcFile
// @Summary 分页获取文件管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文件管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcFile/getWcFileList [get]
export const getWcFileList = (params) => {
  return service({
    url: '/wcFile/getWcFileList',
    method: 'get',
    params
  })
}
