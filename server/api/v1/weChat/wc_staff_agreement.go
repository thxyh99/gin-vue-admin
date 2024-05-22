package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type WcStaffAgreementApi struct {
}

var wcStaffAgreementService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffAgreementService

// CreateWcStaffAgreement 创建合同信息
// @Tags WcStaffAgreement
// @Summary 创建合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffAgreement true "创建合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffAgreement/createWcStaffAgreement [post]
func (wcStaffAgreementApi *WcStaffAgreementApi) CreateWcStaffAgreement(c *gin.Context) {
	var wcStaffAgreement weChat.WcStaffAgreement
	err := c.ShouldBindJSON(&wcStaffAgreement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffAgreementService.CreateWcStaffAgreement(&wcStaffAgreement); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffAgreement 删除合同信息
// @Tags WcStaffAgreement
// @Summary 删除合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffAgreement true "删除合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAgreement/deleteWcStaffAgreement [delete]
func (wcStaffAgreementApi *WcStaffAgreementApi) DeleteWcStaffAgreement(c *gin.Context) {
	ID := c.Query("ID")
	agreement, err := wcStaffAgreementService.GetWcStaffAgreement(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败，合同不存在!", zap.Error(err))
		response.FailWithMessage("删除失败，合同不存在!", c)
	}
	if *agreement.FileId != 0 {
		fileId := strconv.Itoa(*agreement.FileId)
		if err := wcFileService.DeleteWcFile(fileId); err != nil {
			global.GVA_LOG.Error("删除合同附件失败!", zap.Error(err))
			response.FailWithMessage("删除合同附件失败!", c)
		}
	}
	if err := wcStaffAgreementService.DeleteWcStaffAgreement(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffAgreementByIds 批量删除合同信息
// @Tags WcStaffAgreement
// @Summary 批量删除合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffAgreement/deleteWcStaffAgreementByIds [delete]
func (wcStaffAgreementApi *WcStaffAgreementApi) DeleteWcStaffAgreementByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffAgreementService.DeleteWcStaffAgreementByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffAgreement 更新合同信息
// @Tags WcStaffAgreement
// @Summary 更新合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffAgreement true "更新合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffAgreement/updateWcStaffAgreement [put]
func (wcStaffAgreementApi *WcStaffAgreementApi) UpdateWcStaffAgreement(c *gin.Context) {
	var wcStaffAgreement weChat.WcStaffAgreement
	err := c.ShouldBindJSON(&wcStaffAgreement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffAgreementService.UpdateWcStaffAgreement(wcStaffAgreement); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffAgreement 用id查询合同信息
// @Tags WcStaffAgreement
// @Summary 用id查询合同信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffAgreement true "用id查询合同信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffAgreement/findWcStaffAgreement [get]
func (wcStaffAgreementApi *WcStaffAgreementApi) FindWcStaffAgreement(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffAgreement, err := wcStaffAgreementService.GetWcStaffAgreement(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffAgreement": rewcStaffAgreement}, c)
	}
}

// GetWcStaffAgreementList 分页获取合同信息列表
// @Tags WcStaffAgreement
// @Summary 分页获取合同信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffAgreementSearch true "分页获取合同信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAgreement/getWcStaffAgreementList [get]
func (wcStaffAgreementApi *WcStaffAgreementApi) GetWcStaffAgreementList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffAgreementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sId := c.Query("staffId")
	if sId == "" {
		response.FailWithMessage("员工ID不能为空", c)
		return
	}
	staffId, _ := strconv.Atoi(sId)
	if list, total, err := wcStaffAgreementService.GetWcStaffAgreementInfoList(pageInfo, staffId); err != nil {
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

// GetWcStaffAgreementPublic 不需要鉴权的合同信息接口
// @Tags WcStaffAgreement
// @Summary 不需要鉴权的合同信息接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffAgreementSearch true "分页获取合同信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAgreement/getWcStaffAgreementList [get]
func (wcStaffAgreementApi *WcStaffAgreementApi) GetWcStaffAgreementPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的合同信息接口信息",
	}, "获取成功", c)
}
