package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
    weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type WcStaffMaterialsApi struct {
}

var wcStaffMaterialsService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffMaterialsService


// CreateWcStaffMaterials 创建证件材料
// @Tags WcStaffMaterials
// @Summary 创建证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffMaterials true "创建证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffMaterials/createWcStaffMaterials [post]
func (wcStaffMaterialsApi *WcStaffMaterialsApi) CreateWcStaffMaterials(c *gin.Context) {
	var wcStaffMaterials weChat.WcStaffMaterials
	err := c.ShouldBindJSON(&wcStaffMaterials)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffMaterialsService.CreateWcStaffMaterials(&wcStaffMaterials); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffMaterials 删除证件材料
// @Tags WcStaffMaterials
// @Summary 删除证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffMaterials true "删除证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffMaterials/deleteWcStaffMaterials [delete]
func (wcStaffMaterialsApi *WcStaffMaterialsApi) DeleteWcStaffMaterials(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffMaterialsService.DeleteWcStaffMaterials(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffMaterialsByIds 批量删除证件材料
// @Tags WcStaffMaterials
// @Summary 批量删除证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffMaterials/deleteWcStaffMaterialsByIds [delete]
func (wcStaffMaterialsApi *WcStaffMaterialsApi) DeleteWcStaffMaterialsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffMaterialsService.DeleteWcStaffMaterialsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffMaterials 更新证件材料
// @Tags WcStaffMaterials
// @Summary 更新证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffMaterials true "更新证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffMaterials/updateWcStaffMaterials [put]
func (wcStaffMaterialsApi *WcStaffMaterialsApi) UpdateWcStaffMaterials(c *gin.Context) {
	var wcStaffMaterials weChat.WcStaffMaterials
	err := c.ShouldBindJSON(&wcStaffMaterials)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffMaterialsService.UpdateWcStaffMaterials(wcStaffMaterials); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffMaterials 用id查询证件材料
// @Tags WcStaffMaterials
// @Summary 用id查询证件材料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffMaterials true "用id查询证件材料"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffMaterials/findWcStaffMaterials [get]
func (wcStaffMaterialsApi *WcStaffMaterialsApi) FindWcStaffMaterials(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffMaterials, err := wcStaffMaterialsService.GetWcStaffMaterials(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffMaterials": rewcStaffMaterials}, c)
	}
}

// GetWcStaffMaterialsList 分页获取证件材料列表
// @Tags WcStaffMaterials
// @Summary 分页获取证件材料列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffMaterialsSearch true "分页获取证件材料列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffMaterials/getWcStaffMaterialsList [get]
func (wcStaffMaterialsApi *WcStaffMaterialsApi) GetWcStaffMaterialsList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffMaterialsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffMaterialsService.GetWcStaffMaterialsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetWcStaffMaterialsPublic 不需要鉴权的证件材料接口
// @Tags WcStaffMaterials
// @Summary 不需要鉴权的证件材料接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffMaterialsSearch true "分页获取证件材料列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffMaterials/getWcStaffMaterialsList [get]
func (wcStaffMaterialsApi *WcStaffMaterialsApi) GetWcStaffMaterialsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的证件材料接口信息",
    }, "获取成功", c)
}
