import service from '@/utils/request'

// @Tags WcRank
// @Summary 创建职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcRank true "创建职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcRank/createWcRank [post]
export const createWcRank = (data) => {
  return service({
    url: '/wcRank/createWcRank',
    method: 'post',
    data
  })
}

// @Tags WcRank
// @Summary 删除职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcRank true "删除职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcRank/deleteWcRank [delete]
export const deleteWcRank = (params) => {
  return service({
    url: '/wcRank/deleteWcRank',
    method: 'delete',
    params
  })
}

// @Tags WcRank
// @Summary 批量删除职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcRank/deleteWcRank [delete]
export const deleteWcRankByIds = (params) => {
  return service({
    url: '/wcRank/deleteWcRankByIds',
    method: 'delete',
    params
  })
}

// @Tags WcRank
// @Summary 更新职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcRank true "更新职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcRank/updateWcRank [put]
export const updateWcRank = (data) => {
  return service({
    url: '/wcRank/updateWcRank',
    method: 'put',
    data
  })
}

// @Tags WcRank
// @Summary 用id查询职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcRank true "用id查询职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcRank/findWcRank [get]
export const findWcRank = (params) => {
  return service({
    url: '/wcRank/findWcRank',
    method: 'get',
    params
  })
}

// @Tags WcRank
// @Summary 分页获取职级管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取职级管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcRank/getWcRankList [get]
export const getWcRankList = (params) => {
  return service({
    url: '/wcRank/getWcRankList',
    method: 'get',
    params
  })
}

// @Tags WcRank
// @Summary 分页获取职级类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取职级类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcRank/getRankTypeList [get]
export const getRankTypeList = (params) => {
  return service({
    url: '/wcRank/getRankTypeList',
    method: 'get',
    params
  })
}
