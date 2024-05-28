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

type WcStaffSocialApi struct {
}

var wcStaffSocialService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffSocialService


// CreateWcStaffSocial 创建社保公积金管理
// @Tags WcStaffSocial
// @Summary 创建社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffSocial true "创建社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffSocial/createWcStaffSocial [post]
func (wcStaffSocialApi *WcStaffSocialApi) CreateWcStaffSocial(c *gin.Context) {
	var wcStaffSocial weChat.WcStaffSocial
	err := c.ShouldBindJSON(&wcStaffSocial)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffSocialService.CreateWcStaffSocial(&wcStaffSocial); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffSocial 删除社保公积金管理
// @Tags WcStaffSocial
// @Summary 删除社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffSocial true "删除社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffSocial/deleteWcStaffSocial [delete]
func (wcStaffSocialApi *WcStaffSocialApi) DeleteWcStaffSocial(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffSocialService.DeleteWcStaffSocial(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffSocialByIds 批量删除社保公积金管理
// @Tags WcStaffSocial
// @Summary 批量删除社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffSocial/deleteWcStaffSocialByIds [delete]
func (wcStaffSocialApi *WcStaffSocialApi) DeleteWcStaffSocialByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffSocialService.DeleteWcStaffSocialByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffSocial 更新社保公积金管理
// @Tags WcStaffSocial
// @Summary 更新社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffSocial true "更新社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffSocial/updateWcStaffSocial [put]
func (wcStaffSocialApi *WcStaffSocialApi) UpdateWcStaffSocial(c *gin.Context) {
	var wcStaffSocial weChat.WcStaffSocial
	err := c.ShouldBindJSON(&wcStaffSocial)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffSocialService.UpdateWcStaffSocial(wcStaffSocial); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffSocial 用id查询社保公积金管理
// @Tags WcStaffSocial
// @Summary 用id查询社保公积金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffSocial true "用id查询社保公积金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffSocial/findWcStaffSocial [get]
func (wcStaffSocialApi *WcStaffSocialApi) FindWcStaffSocial(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffSocial, err := wcStaffSocialService.GetWcStaffSocial(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffSocial": rewcStaffSocial}, c)
	}
}

// GetWcStaffSocialList 分页获取社保公积金管理列表
// @Tags WcStaffSocial
// @Summary 分页获取社保公积金管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSocialSearch true "分页获取社保公积金管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffSocial/getWcStaffSocialList [get]
func (wcStaffSocialApi *WcStaffSocialApi) GetWcStaffSocialList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffSocialSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffSocialService.GetWcStaffSocialInfoList(pageInfo); err != nil {
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

// GetWcStaffSocialPublic 不需要鉴权的社保公积金管理接口
// @Tags WcStaffSocial
// @Summary 不需要鉴权的社保公积金管理接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSocialSearch true "分页获取社保公积金管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffSocial/getWcStaffSocialList [get]
func (wcStaffSocialApi *WcStaffSocialApi) GetWcStaffSocialPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的社保公积金管理接口信息",
    }, "获取成功", c)
}
