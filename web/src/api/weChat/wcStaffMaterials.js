import service from '@/utils/request'

// @Tags WcStaffMaterials
// @Summary 创建证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffMaterials true "创建证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffMaterials/createWcStaffMaterials [post]
export const createWcStaffMaterials = (data) => {
  return service({
    url: '/wcStaffMaterials/createWcStaffMaterials',
    method: 'post',
    data
  })
}

// @Tags WcStaffMaterials
// @Summary 删除证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffMaterials true "删除证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffMaterials/deleteWcStaffMaterials [delete]
export const deleteWcStaffMaterials = (params) => {
  return service({
    url: '/wcStaffMaterials/deleteWcStaffMaterials',
    method: 'delete',
    params
  })
}

// @Tags WcStaffMaterials
// @Summary 批量删除证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffMaterials/deleteWcStaffMaterials [delete]
export const deleteWcStaffMaterialsByIds = (params) => {
  return service({
    url: '/wcStaffMaterials/deleteWcStaffMaterialsByIds',
    method: 'delete',
    params
  })
}

// @Tags WcStaffMaterials
// @Summary 更新证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WcStaffMaterials true "更新证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffMaterials/updateWcStaffMaterials [put]
export const updateWcStaffMaterials = (data) => {
  return service({
    url: '/wcStaffMaterials/updateWcStaffMaterials',
    method: 'put',
    data
  })
}

// @Tags WcStaffMaterials
// @Summary 用id查询证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WcStaffMaterials true "用id查询证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffMaterials/findWcStaffMaterials [get]
export const findWcStaffMaterials = (params) => {
  return service({
    url: '/wcStaffMaterials/findWcStaffMaterials',
    method: 'get',
    params
  })
}

// @Tags WcStaffMaterials
// @Summary 分页获取证件材料列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取证件材料列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffMaterials/getWcStaffMaterialsList [get]
export const getWcStaffMaterialsList = (params) => {
  return service({
    url: '/wcStaffMaterials/getWcStaffMaterialsList',
    method: 'get',
    params
  })
}
