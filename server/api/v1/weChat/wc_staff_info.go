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

type WcStaffInfoApi struct {
}

var wcStaffInfoService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffInfoService


// CreateWcStaffInfo 创建个人信息
// @Tags WcStaffInfo
// @Summary 创建个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffInfo true "创建个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffInfo/createWcStaffInfo [post]
func (wcStaffInfoApi *WcStaffInfoApi) CreateWcStaffInfo(c *gin.Context) {
	var wcStaffInfo weChat.WcStaffInfo
	err := c.ShouldBindJSON(&wcStaffInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffInfoService.CreateWcStaffInfo(&wcStaffInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffInfo 删除个人信息
// @Tags WcStaffInfo
// @Summary 删除个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffInfo true "删除个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffInfo/deleteWcStaffInfo [delete]
func (wcStaffInfoApi *WcStaffInfoApi) DeleteWcStaffInfo(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffInfoService.DeleteWcStaffInfo(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffInfoByIds 批量删除个人信息
// @Tags WcStaffInfo
// @Summary 批量删除个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffInfo/deleteWcStaffInfoByIds [delete]
func (wcStaffInfoApi *WcStaffInfoApi) DeleteWcStaffInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffInfoService.DeleteWcStaffInfoByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffInfo 更新个人信息
// @Tags WcStaffInfo
// @Summary 更新个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffInfo true "更新个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffInfo/updateWcStaffInfo [put]
func (wcStaffInfoApi *WcStaffInfoApi) UpdateWcStaffInfo(c *gin.Context) {
	var wcStaffInfo weChat.WcStaffInfo
	err := c.ShouldBindJSON(&wcStaffInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffInfoService.UpdateWcStaffInfo(wcStaffInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffInfo 用id查询个人信息
// @Tags WcStaffInfo
// @Summary 用id查询个人信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffInfo true "用id查询个人信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffInfo/findWcStaffInfo [get]
func (wcStaffInfoApi *WcStaffInfoApi) FindWcStaffInfo(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffInfo, err := wcStaffInfoService.GetWcStaffInfo(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffInfo": rewcStaffInfo}, c)
	}
}

// GetWcStaffInfoList 分页获取个人信息列表
// @Tags WcStaffInfo
// @Summary 分页获取个人信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffInfoSearch true "分页获取个人信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffInfo/getWcStaffInfoList [get]
func (wcStaffInfoApi *WcStaffInfoApi) GetWcStaffInfoList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffInfoService.GetWcStaffInfoInfoList(pageInfo); err != nil {
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

// GetWcStaffInfoPublic 不需要鉴权的个人信息接口
// @Tags WcStaffInfo
// @Summary 不需要鉴权的个人信息接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffInfoSearch true "分页获取个人信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffInfo/getWcStaffInfoList [get]
func (wcStaffInfoApi *WcStaffInfoApi) GetWcStaffInfoPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的个人信息接口信息",
    }, "获取成功", c)
}
