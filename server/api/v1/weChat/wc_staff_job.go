package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WcStaffJobApi struct {
}

var wcStaffJobService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffJobService

// CreateWcStaffJob 创建工作信息
// @Tags WcStaffJob
// @Summary 创建工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffJob true "创建工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffJob/createWcStaffJob [post]
func (wcStaffJobApi *WcStaffJobApi) CreateWcStaffJob(c *gin.Context) {
	var wcStaffJobRequest weChatReq.WcStaffJobRequest
	err := c.ShouldBindJSON(&wcStaffJobRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffJobService.CreateWcStaffJob(&wcStaffJobRequest); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffJob 删除工作信息
// @Tags WcStaffJob
// @Summary 删除工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffJob true "删除工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffJob/deleteWcStaffJob [delete]
func (wcStaffJobApi *WcStaffJobApi) DeleteWcStaffJob(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffJobService.DeleteWcStaffJob(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffJobByIds 批量删除工作信息
// @Tags WcStaffJob
// @Summary 批量删除工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffJob/deleteWcStaffJobByIds [delete]
func (wcStaffJobApi *WcStaffJobApi) DeleteWcStaffJobByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffJobService.DeleteWcStaffJobByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffJob 更新工作信息
// @Tags WcStaffJob
// @Summary 更新工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffJob true "更新工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffJob/updateWcStaffJob [put]
func (wcStaffJobApi *WcStaffJobApi) UpdateWcStaffJob(c *gin.Context) {
	var wcStaffJobRequest weChatReq.WcStaffJobRequest
	err := c.ShouldBindJSON(&wcStaffJobRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffJobService.UpdateWcStaffJob(&wcStaffJobRequest); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffJob 用id查询工作信息
// @Tags WcStaffJob
// @Summary 用id查询工作信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffJob true "用id查询工作信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffJob/findWcStaffJob [get]
func (wcStaffJobApi *WcStaffJobApi) FindWcStaffJob(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffJob, err := wcStaffJobService.GetWcStaffJob(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffJob": rewcStaffJob}, c)
	}
}

// GetWcStaffJobList 分页获取工作信息列表
// @Tags WcStaffJob
// @Summary 分页获取工作信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffJobSearch true "分页获取工作信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffJob/getWcStaffJobList [get]
func (wcStaffJobApi *WcStaffJobApi) GetWcStaffJobList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffJobSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffJobService.GetWcStaffJobInfoList(pageInfo); err != nil {
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

// GetWcStaffJobPublic 不需要鉴权的工作信息接口
// @Tags WcStaffJob
// @Summary 不需要鉴权的工作信息接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffJobSearch true "分页获取工作信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffJob/getWcStaffJobList [get]
func (wcStaffJobApi *WcStaffJobApi) GetWcStaffJobPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的工作信息接口信息",
	}, "获取成功", c)
}
