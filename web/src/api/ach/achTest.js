import service from '@/utils/request'

// @Tags AchTest
// @Summary 创建绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AchTest true "创建绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /achTest/createAchTest [post]
export const createAchTest = (data) => {
  return service({
    url: '/achTest/createAchTest',
    method: 'post',
    data
  })
}

// @Tags AchTest
// @Summary 删除绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AchTest true "删除绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /achTest/deleteAchTest [delete]
export const deleteAchTest = (params) => {
  return service({
    url: '/achTest/deleteAchTest',
    method: 'delete',
    params
  })
}

// @Tags AchTest
// @Summary 批量删除绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /achTest/deleteAchTest [delete]
export const deleteAchTestByIds = (params) => {
  return service({
    url: '/achTest/deleteAchTestByIds',
    method: 'delete',
    params
  })
}

// @Tags AchTest
// @Summary 更新绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AchTest true "更新绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /achTest/updateAchTest [put]
export const updateAchTest = (data) => {
  return service({
    url: '/achTest/updateAchTest',
    method: 'put',
    data
  })
}

// @Tags AchTest
// @Summary 用id查询绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AchTest true "用id查询绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /achTest/findAchTest [get]
export const findAchTest = (params) => {
  return service({
    url: '/achTest/findAchTest',
    method: 'get',
    params
  })
}

// @Tags AchTest
// @Summary 分页获取绩效测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取绩效测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /achTest/getAchTestList [get]
export const getAchTestList = (params) => {
  return service({
    url: '/achTest/getAchTestList',
    method: 'get',
    params
  })
}
