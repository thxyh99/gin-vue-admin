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

type WcStaffApi struct {
}

var wcStaffService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffService


// CreateWcStaff 创建账号信息
// @Tags WcStaff
// @Summary 创建账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaff true "创建账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaff/createWcStaff [post]
func (wcStaffApi *WcStaffApi) CreateWcStaff(c *gin.Context) {
	var wcStaff weChat.WcStaff
	err := c.ShouldBindJSON(&wcStaff)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffService.CreateWcStaff(&wcStaff); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaff 删除账号信息
// @Tags WcStaff
// @Summary 删除账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaff true "删除账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaff/deleteWcStaff [delete]
func (wcStaffApi *WcStaffApi) DeleteWcStaff(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffService.DeleteWcStaff(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffByIds 批量删除账号信息
// @Tags WcStaff
// @Summary 批量删除账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaff/deleteWcStaffByIds [delete]
func (wcStaffApi *WcStaffApi) DeleteWcStaffByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffService.DeleteWcStaffByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaff 更新账号信息
// @Tags WcStaff
// @Summary 更新账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaff true "更新账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaff/updateWcStaff [put]
func (wcStaffApi *WcStaffApi) UpdateWcStaff(c *gin.Context) {
	var wcStaff weChat.WcStaff
	err := c.ShouldBindJSON(&wcStaff)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffService.UpdateWcStaff(wcStaff); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaff 用id查询账号信息
// @Tags WcStaff
// @Summary 用id查询账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaff true "用id查询账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaff/findWcStaff [get]
func (wcStaffApi *WcStaffApi) FindWcStaff(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaff, err := wcStaffService.GetWcStaff(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaff": rewcStaff}, c)
	}
}

// GetWcStaffList 分页获取账号信息列表
// @Tags WcStaff
// @Summary 分页获取账号信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSearch true "分页获取账号信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaff/getWcStaffList [get]
func (wcStaffApi *WcStaffApi) GetWcStaffList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffService.GetWcStaffInfoList(pageInfo); err != nil {
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

// GetWcStaffPublic 不需要鉴权的账号信息接口
// @Tags WcStaff
// @Summary 不需要鉴权的账号信息接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSearch true "分页获取账号信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaff/getWcStaffList [get]
func (wcStaffApi *WcStaffApi) GetWcStaffPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的账号信息接口信息",
    }, "获取成功", c)
}
