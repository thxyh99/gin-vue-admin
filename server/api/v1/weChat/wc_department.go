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

type WcDepartmentApi struct {
}

var wcDepartmentService = service.ServiceGroupApp.WeChatServiceGroup.WcDepartmentService

// CreateWcDepartment 创建wcDepartment表
// @Tags WcDepartment
// @Summary 创建wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcDepartment true "创建wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcDepartment/createWcDepartment [post]
func (wcDepartmentApi *WcDepartmentApi) CreateWcDepartment(c *gin.Context) {
	var wcDepartment weChat.WcDepartment
	err := c.ShouldBindJSON(&wcDepartment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcDepartmentService.CreateWcDepartment(&wcDepartment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcDepartment 删除wcDepartment表
// @Tags WcDepartment
// @Summary 删除wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcDepartment true "删除wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcDepartment/deleteWcDepartment [delete]
func (wcDepartmentApi *WcDepartmentApi) DeleteWcDepartment(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcDepartmentService.DeleteWcDepartment(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcDepartmentByIds 批量删除wcDepartment表
// @Tags WcDepartment
// @Summary 批量删除wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcDepartment/deleteWcDepartmentByIds [delete]
func (wcDepartmentApi *WcDepartmentApi) DeleteWcDepartmentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcDepartmentService.DeleteWcDepartmentByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcDepartment 更新wcDepartment表
// @Tags WcDepartment
// @Summary 更新wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcDepartment true "更新wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcDepartment/updateWcDepartment [put]
func (wcDepartmentApi *WcDepartmentApi) UpdateWcDepartment(c *gin.Context) {
	var wcDepartment weChat.WcDepartment
	err := c.ShouldBindJSON(&wcDepartment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcDepartmentService.UpdateWcDepartment(wcDepartment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcDepartment 用id查询wcDepartment表
// @Tags WcDepartment
// @Summary 用id查询wcDepartment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcDepartment true "用id查询wcDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcDepartment/findWcDepartment [get]
func (wcDepartmentApi *WcDepartmentApi) FindWcDepartment(c *gin.Context) {
	ID := c.Query("ID")
	if rewcDepartment, err := wcDepartmentService.GetWcDepartment(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcDepartment": rewcDepartment}, c)
	}
}

// GetWcDepartmentList 分页获取wcDepartment表列表
// @Tags WcDepartment
// @Summary 分页获取wcDepartment表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcDepartmentSearch true "分页获取wcDepartment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcDepartment/getWcDepartmentList [get]
func (wcDepartmentApi *WcDepartmentApi) GetWcDepartmentList(c *gin.Context) {
	var pageInfo weChatReq.WcDepartmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcDepartmentService.GetWcDepartmentInfoList(pageInfo); err != nil {
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

// GetAllFullDepartmentList 获取所有全层级部门名称列表
// @Tags WcDepartment
// @Summary 分页获取wcDepartment表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcDepartmentSearch true "分页获取wcDepartment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcDepartment/getWcDepartmentList [get]
func (wcDepartmentApi *WcDepartmentApi) GetAllFullDepartmentList(c *gin.Context) {
	var pageInfo weChatReq.WcDepartmentSearch
	if list, total, err := wcDepartmentService.GetAllFullDepartmentList(); err != nil {
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

// GetWcDepartmentPublic 不需要鉴权的wcDepartment表接口
// @Tags WcDepartment
// @Summary 不需要鉴权的wcDepartment表接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcDepartmentSearch true "分页获取wcDepartment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcDepartment/getWcDepartmentList [get]
func (wcDepartmentApi *WcDepartmentApi) GetWcDepartmentPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的wcDepartment表接口信息",
	}, "获取成功", c)
}
