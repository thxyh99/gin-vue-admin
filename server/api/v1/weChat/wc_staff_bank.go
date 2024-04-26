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

type WcStaffBankApi struct {
}

var wcStaffBankService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffBankService


// CreateWcStaffBank 创建银行卡信息
// @Tags WcStaffBank
// @Summary 创建银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffBank true "创建银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffBank/createWcStaffBank [post]
func (wcStaffBankApi *WcStaffBankApi) CreateWcStaffBank(c *gin.Context) {
	var wcStaffBank weChat.WcStaffBank
	err := c.ShouldBindJSON(&wcStaffBank)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffBankService.CreateWcStaffBank(&wcStaffBank); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffBank 删除银行卡信息
// @Tags WcStaffBank
// @Summary 删除银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffBank true "删除银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffBank/deleteWcStaffBank [delete]
func (wcStaffBankApi *WcStaffBankApi) DeleteWcStaffBank(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffBankService.DeleteWcStaffBank(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffBankByIds 批量删除银行卡信息
// @Tags WcStaffBank
// @Summary 批量删除银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffBank/deleteWcStaffBankByIds [delete]
func (wcStaffBankApi *WcStaffBankApi) DeleteWcStaffBankByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffBankService.DeleteWcStaffBankByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffBank 更新银行卡信息
// @Tags WcStaffBank
// @Summary 更新银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffBank true "更新银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffBank/updateWcStaffBank [put]
func (wcStaffBankApi *WcStaffBankApi) UpdateWcStaffBank(c *gin.Context) {
	var wcStaffBank weChat.WcStaffBank
	err := c.ShouldBindJSON(&wcStaffBank)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffBankService.UpdateWcStaffBank(wcStaffBank); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffBank 用id查询银行卡信息
// @Tags WcStaffBank
// @Summary 用id查询银行卡信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffBank true "用id查询银行卡信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffBank/findWcStaffBank [get]
func (wcStaffBankApi *WcStaffBankApi) FindWcStaffBank(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffBank, err := wcStaffBankService.GetWcStaffBank(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffBank": rewcStaffBank}, c)
	}
}

// GetWcStaffBankList 分页获取银行卡信息列表
// @Tags WcStaffBank
// @Summary 分页获取银行卡信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffBankSearch true "分页获取银行卡信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffBank/getWcStaffBankList [get]
func (wcStaffBankApi *WcStaffBankApi) GetWcStaffBankList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffBankSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffBankService.GetWcStaffBankInfoList(pageInfo); err != nil {
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

// GetWcStaffBankPublic 不需要鉴权的银行卡信息接口
// @Tags WcStaffBank
// @Summary 不需要鉴权的银行卡信息接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffBankSearch true "分页获取银行卡信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffBank/getWcStaffBankList [get]
func (wcStaffBankApi *WcStaffBankApi) GetWcStaffBankPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的银行卡信息接口信息",
    }, "获取成功", c)
}
