package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WcRankApi struct {
}

var wcRankService = service.ServiceGroupApp.WeChatServiceGroup.WcRankService

// CreateWcRank 创建职级管理
// @Tags WcRank
// @Summary 创建职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcRank true "创建职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcRank/createWcRank [post]
func (wcRankApi *WcRankApi) CreateWcRank(c *gin.Context) {
	var wcRank weChat.WcRank
	err := c.ShouldBindJSON(&wcRank)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcRankService.CreateWcRank(&wcRank); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcRank 删除职级管理
// @Tags WcRank
// @Summary 删除职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcRank true "删除职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcRank/deleteWcRank [delete]
func (wcRankApi *WcRankApi) DeleteWcRank(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcRankService.DeleteWcRank(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcRankByIds 批量删除职级管理
// @Tags WcRank
// @Summary 批量删除职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcRank/deleteWcRankByIds [delete]
func (wcRankApi *WcRankApi) DeleteWcRankByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcRankService.DeleteWcRankByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcRank 更新职级管理
// @Tags WcRank
// @Summary 更新职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcRank true "更新职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcRank/updateWcRank [put]
func (wcRankApi *WcRankApi) UpdateWcRank(c *gin.Context) {
	var wcRank weChat.WcRank
	err := c.ShouldBindJSON(&wcRank)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcRankService.UpdateWcRank(wcRank); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcRank 用id查询职级管理
// @Tags WcRank
// @Summary 用id查询职级管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcRank true "用id查询职级管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcRank/findWcRank [get]
func (wcRankApi *WcRankApi) FindWcRank(c *gin.Context) {
	ID := c.Query("ID")
	if rewcRank, err := wcRankService.GetWcRank(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcRank": rewcRank}, c)
	}
}

// GetWcRankList 分页获取职级管理列表
// @Tags WcRank
// @Summary 分页获取职级管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcRankSearch true "分页获取职级管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcRank/getWcRankList [get]
func (wcRankApi *WcRankApi) GetWcRankList(c *gin.Context) {
	var pageInfo weChatReq.WcRankSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcRankService.GetWcRankInfoList(pageInfo); err != nil {
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

// GetRankTypeList 分页获取职级类型列表
// @Tags WcRank
// @Summary 分页获取职级管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcRankSearch true "分页获取职级类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcRank/getRankTypeList [get]
func (wcRankApi *WcRankApi) GetRankTypeList(c *gin.Context) {
	var pageInfo weChatReq.WcRankSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcRankService.GetRankTypeList(pageInfo); err != nil {
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

// GetWcRankPublic 不需要鉴权的职级管理接口
// @Tags WcRank
// @Summary 不需要鉴权的职级管理接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcRankSearch true "分页获取职级管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcRank/getWcRankList [get]
func (wcRankApi *WcRankApi) GetWcRankPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的职级管理接口信息",
	}, "获取成功", c)
}
