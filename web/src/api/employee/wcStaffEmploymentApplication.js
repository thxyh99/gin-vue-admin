import service from '@/utils/request'

// @Tags WcStaffEmploymentApplication
// @Summary 创建入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffEmploymentApplication true "创建入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffEmploymentApplication/createWcStaffEmploymentApplication [post]
export const createWcStaffEmploymentApplication = (data) => {
  return service({
    url: '/wcStaffEmploymentApplication/createWcStaffEmploymentApplication',
    method: 'post',
    data
  })
}

// @Tags WcStaffEmploymentApplication
// @Summary 删除入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffEmploymentApplication true "删除入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffEmploymentApplication/deleteWcStaffEmploymentApplication [delete]
export const deleteWcStaffEmploymentApplication = (params) => {
  return service({
    url: '/wcStaffEmploymentApplication/deleteWcStaffEmploymentApplication',
    method: 'delete',
    params
  })
}

// @Tags WcStaffEmploymentApplication
// @Summary 批量删除入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffEmploymentApplication/deleteWcStaffEmploymentApplication [delete]
export const deleteWcStaffEmploymentApplicationByIds = (params) => {
  return service({
    url: '/wcStaffEmploymentApplication/deleteWcStaffEmploymentApplicationByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffEmploymentApplication
// @Summary 更新入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffEmploymentApplication true "更新入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffEmploymentApplication/updateWcStaffEmploymentApplication [put]
export const updateWcStaffEmploymentApplication = (data) => {
  return service({
    url: '/wcStaffEmploymentApplication/updateWcStaffEmploymentApplication',
    method: 'put',
    data
  })
}

// @Tags WcStaffEmploymentApplication
// @Summary 用id查询入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffEmploymentApplication true "用id查询入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffEmploymentApplication/findWcStaffEmploymentApplication [get]
export const findWcStaffEmploymentApplication = (params) => {
  return service({
    url: '/wcStaffEmploymentApplication/findWcStaffEmploymentApplication',
    method: 'get',
    params
  })
}

// @Tags WcStaffEmploymentApplication
// @Summary 分页获取入职申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取入职申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffEmploymentApplication/getWcStaffEmploymentApplicationList [get]
export const getWcStaffEmploymentApplicationList = (params) => {
  return service({
    url: '/wcStaffEmploymentApplication/getWcStaffEmploymentApplicationList',
    method: 'get',
    params
  })
}

// @Tags WcStaffEmploymentApplication
// @Summary 创建入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffEmploymentApplication true "创建入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffEmploymentApplication/createWcStaffEmploymentApplication [post]
export const CreateOAStaffEmploymentApplication = (data) => {
  return service({
    url: '/wcStaffEmploymentApplication/CreateOAStaffEmploymentApplication',
    method: 'get',
    data
  })
}

