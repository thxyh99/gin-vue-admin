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

type WcPositionApi struct {
}

var wcPositionService = service.ServiceGroupApp.WeChatServiceGroup.WcPositionService


// CreateWcPosition 创建职位管理
// @Tags WcPosition
// @Summary 创建职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcPosition true "创建职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcPosition/createWcPosition [post]
func (wcPositionApi *WcPositionApi) CreateWcPosition(c *gin.Context) {
	var wcPosition weChat.WcPosition
	err := c.ShouldBindJSON(&wcPosition)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcPositionService.CreateWcPosition(&wcPosition); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcPosition 删除职位管理
// @Tags WcPosition
// @Summary 删除职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcPosition true "删除职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcPosition/deleteWcPosition [delete]
func (wcPositionApi *WcPositionApi) DeleteWcPosition(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcPositionService.DeleteWcPosition(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcPositionByIds 批量删除职位管理
// @Tags WcPosition
// @Summary 批量删除职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcPosition/deleteWcPositionByIds [delete]
func (wcPositionApi *WcPositionApi) DeleteWcPositionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcPositionService.DeleteWcPositionByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcPosition 更新职位管理
// @Tags WcPosition
// @Summary 更新职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcPosition true "更新职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcPosition/updateWcPosition [put]
func (wcPositionApi *WcPositionApi) UpdateWcPosition(c *gin.Context) {
	var wcPosition weChat.WcPosition
	err := c.ShouldBindJSON(&wcPosition)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcPositionService.UpdateWcPosition(wcPosition); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcPosition 用id查询职位管理
// @Tags WcPosition
// @Summary 用id查询职位管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcPosition true "用id查询职位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcPosition/findWcPosition [get]
func (wcPositionApi *WcPositionApi) FindWcPosition(c *gin.Context) {
	ID := c.Query("ID")
	if rewcPosition, err := wcPositionService.GetWcPosition(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcPosition": rewcPosition}, c)
	}
}

// GetWcPositionList 分页获取职位管理列表
// @Tags WcPosition
// @Summary 分页获取职位管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcPositionSearch true "分页获取职位管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcPosition/getWcPositionList [get]
func (wcPositionApi *WcPositionApi) GetWcPositionList(c *gin.Context) {
	var pageInfo weChatReq.WcPositionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcPositionService.GetWcPositionInfoList(pageInfo); err != nil {
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

// GetWcPositionPublic 不需要鉴权的职位管理接口
// @Tags WcPosition
// @Summary 不需要鉴权的职位管理接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcPositionSearch true "分页获取职位管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcPosition/getWcPositionList [get]
func (wcPositionApi *WcPositionApi) GetWcPositionPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的职位管理接口信息",
    }, "获取成功", c)
}
