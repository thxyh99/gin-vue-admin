import service from '@/utils/request'

// @Tags WcPosition
// @Summary 创建职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcPosition true "创建职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcPosition/createWcPosition [post]
export const createWcPosition = (data) => {
  return service({
    url: '/wcPosition/createWcPosition',
    method: 'post',
    data
  })
}

// @Tags WcPosition
// @Summary 删除职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcPosition true "删除职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcPosition/deleteWcPosition [delete]
export const deleteWcPosition = (params) => {
  return service({
    url: '/wcPosition/deleteWcPosition',
    method: 'delete',
    params
  })
}

// @Tags WcPosition
// @Summary 批量删除职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcPosition/deleteWcPosition [delete]
export const deleteWcPositionByIds = (params) => {
  return service({
    url: '/wcPosition/deleteWcPositionByIds',
    method: 'delete',
    params
  })
}

// @Tags WcPosition
// @Summary 更新职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcPosition true "更新职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcPosition/updateWcPosition [put]
export const updateWcPosition = (data) => {
  return service({
    url: '/wcPosition/updateWcPosition',
    method: 'put',
    data
  })
}

// @Tags WcPosition
// @Summary 用id查询职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcPosition true "用id查询职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcPosition/findWcPosition [get]
export const findWcPosition = (params) => {
  return service({
    url: '/wcPosition/findWcPosition',
    method: 'get',
    params
  })
}

// @Tags WcPosition
// @Summary 分页获取职位管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取职位管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcPosition/getWcPositionList [get]
export const getWcPositionList = (params) => {
  return service({
    url: '/wcPosition/getWcPositionList',
    method: 'get',
    params
  })
}
