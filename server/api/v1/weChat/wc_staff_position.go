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

type WcStaffPositionApi struct {
}

var wcStaffPositionService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffPositionService


// CreateWcStaffPosition 创建员工职位管理
// @Tags WcStaffPosition
// @Summary 创建员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffPosition true "创建员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffPosition/createWcStaffPosition [post]
func (wcStaffPositionApi *WcStaffPositionApi) CreateWcStaffPosition(c *gin.Context) {
	var wcStaffPosition weChat.WcStaffPosition
	err := c.ShouldBindJSON(&wcStaffPosition)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffPositionService.CreateWcStaffPosition(&wcStaffPosition); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffPosition 删除员工职位管理
// @Tags WcStaffPosition
// @Summary 删除员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffPosition true "删除员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffPosition/deleteWcStaffPosition [delete]
func (wcStaffPositionApi *WcStaffPositionApi) DeleteWcStaffPosition(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffPositionService.DeleteWcStaffPosition(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffPositionByIds 批量删除员工职位管理
// @Tags WcStaffPosition
// @Summary 批量删除员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffPosition/deleteWcStaffPositionByIds [delete]
func (wcStaffPositionApi *WcStaffPositionApi) DeleteWcStaffPositionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffPositionService.DeleteWcStaffPositionByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffPosition 更新员工职位管理
// @Tags WcStaffPosition
// @Summary 更新员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffPosition true "更新员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffPosition/updateWcStaffPosition [put]
func (wcStaffPositionApi *WcStaffPositionApi) UpdateWcStaffPosition(c *gin.Context) {
	var wcStaffPosition weChat.WcStaffPosition
	err := c.ShouldBindJSON(&wcStaffPosition)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffPositionService.UpdateWcStaffPosition(wcStaffPosition); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffPosition 用id查询员工职位管理
// @Tags WcStaffPosition
// @Summary 用id查询员工职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffPosition true "用id查询员工职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffPosition/findWcStaffPosition [get]
func (wcStaffPositionApi *WcStaffPositionApi) FindWcStaffPosition(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffPosition, err := wcStaffPositionService.GetWcStaffPosition(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffPosition": rewcStaffPosition}, c)
	}
}

// GetWcStaffPositionList 分页获取员工职位管理列表
// @Tags WcStaffPosition
// @Summary 分页获取员工职位管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffPositionSearch true "分页获取员工职位管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffPosition/getWcStaffPositionList [get]
func (wcStaffPositionApi *WcStaffPositionApi) GetWcStaffPositionList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffPositionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffPositionService.GetWcStaffPositionInfoList(pageInfo); err != nil {
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

// GetWcStaffPositionPublic 不需要鉴权的员工职位管理接口
// @Tags WcStaffPosition
// @Summary 不需要鉴权的员工职位管理接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffPositionSearch true "分页获取员工职位管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffPosition/getWcStaffPositionList [get]
func (wcStaffPositionApi *WcStaffPositionApi) GetWcStaffPositionPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的员工职位管理接口信息",
    }, "获取成功", c)
}
