package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WcSalaryTemplateApi struct {
}

var wcSalaryTemplateService = service.ServiceGroupApp.WeChatServiceGroup.WcSalaryTemplateService

// CreateWcSalaryTemplate 创建工资单模版
// @Tags WcSalaryTemplate
// @Summary 创建工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalaryTemplate true "创建工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcSalaryTemplate/createWcSalaryTemplate [post]
func (wcSalaryTemplateApi *WcSalaryTemplateApi) CreateWcSalaryTemplate(c *gin.Context) {
	var wcSalaryTemplateRequest weChatReq.WcSalaryTemplateCreateRequest
	err := c.ShouldBindJSON(&wcSalaryTemplateRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(wcSalaryTemplateRequest.Fields) == 0 {
		response.FailWithMessage("工资单字段异常", c)
		return
	}

	if err := wcSalaryTemplateService.CreateWcSalaryTemplate(&wcSalaryTemplateRequest); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcSalaryTemplate 删除工资单模版
// @Tags WcSalaryTemplate
// @Summary 删除工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalaryTemplate true "删除工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalaryTemplate/deleteWcSalaryTemplate [delete]
func (wcSalaryTemplateApi *WcSalaryTemplateApi) DeleteWcSalaryTemplate(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcSalaryTemplateService.DeleteWcSalaryTemplate(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcSalaryTemplateByIds 批量删除工资单模版
// @Tags WcSalaryTemplate
// @Summary 批量删除工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcSalaryTemplate/deleteWcSalaryTemplateByIds [delete]
func (wcSalaryTemplateApi *WcSalaryTemplateApi) DeleteWcSalaryTemplateByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcSalaryTemplateService.DeleteWcSalaryTemplateByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcSalaryTemplate 更新工资单模版
// @Tags WcSalaryTemplate
// @Summary 更新工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalaryTemplate true "更新工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcSalaryTemplate/updateWcSalaryTemplate [put]
func (wcSalaryTemplateApi *WcSalaryTemplateApi) UpdateWcSalaryTemplate(c *gin.Context) {
	var wcSalaryTemplateRequest weChatReq.WcSalaryTemplateUpdateRequest
	err := c.ShouldBindJSON(&wcSalaryTemplateRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(wcSalaryTemplateRequest.Fields) == 0 {
		response.FailWithMessage("工资单字段异常", c)
		return
	}

	if err := wcSalaryTemplateService.UpdateWcSalaryTemplate(wcSalaryTemplateRequest); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcSalaryTemplate 用id查询工资单模版
// @Tags WcSalaryTemplate
// @Summary 用id查询工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcSalaryTemplate true "用id查询工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcSalaryTemplate/findWcSalaryTemplate [get]
func (wcSalaryTemplateApi *WcSalaryTemplateApi) FindWcSalaryTemplate(c *gin.Context) {
	ID := c.Query("ID")
	if rewcSalaryTemplate, err := wcSalaryTemplateService.GetWcSalaryTemplate(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcSalaryTemplate": rewcSalaryTemplate}, c)
	}
}

// GetWcSalaryTemplateList 分页获取工资单模版列表
// @Tags WcSalaryTemplate
// @Summary 分页获取工资单模版列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcSalaryTemplateSearch true "分页获取工资单模版列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalaryTemplate/getWcSalaryTemplateList [get]
func (wcSalaryTemplateApi *WcSalaryTemplateApi) GetWcSalaryTemplateList(c *gin.Context) {
	var pageInfo weChatReq.WcSalaryTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcSalaryTemplateService.GetWcSalaryTemplateInfoList(pageInfo); err != nil {
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

// GetWcSalaryFieldsByType 通过工资类型获取工资单字段
func (wcSalaryTemplateApi *WcSalaryTemplateApi) GetWcSalaryFieldsByType(c *gin.Context) {
	salaryType := c.Query("type")
	if list, total, err := wcSalaryTemplateService.GetWcSalaryFieldsByType(salaryType); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		fmt.Println("list", list)
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// GetWcSalaryTemplatePublic 不需要鉴权的工资单模版接口
// @Tags WcSalaryTemplate
// @Summary 不需要鉴权的工资单模版接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcSalaryTemplateSearch true "分页获取工资单模版列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalaryTemplate/getWcSalaryTemplateList [get]
func (wcSalaryTemplateApi *WcSalaryTemplateApi) GetWcSalaryTemplatePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的工资单模版接口信息",
	}, "获取成功", c)
}
