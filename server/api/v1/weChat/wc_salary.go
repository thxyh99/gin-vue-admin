package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WcSalaryApi struct {
}

var wcSalaryService = service.ServiceGroupApp.WeChatServiceGroup.WcSalaryService

// CreateWcSalary 创建工资单模版
// @Tags WcSalary
// @Summary 创建工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalary true "创建工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcSalary/createWcSalary [post]
func (wcSalaryApi *WcSalaryApi) CreateWcSalary(c *gin.Context) {
	var wcSalary weChat.WcSalary
	err := c.ShouldBindJSON(&wcSalary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcSalaryService.CreateWcSalary(&wcSalary); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcSalary 删除工资单模版
// @Tags WcSalary
// @Summary 删除工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalary true "删除工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalary/deleteWcSalary [delete]
func (wcSalaryApi *WcSalaryApi) DeleteWcSalary(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcSalaryService.DeleteWcSalary(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcSalaryByIds 批量删除工资单模版
// @Tags WcSalary
// @Summary 批量删除工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcSalary/deleteWcSalaryByIds [delete]
func (wcSalaryApi *WcSalaryApi) DeleteWcSalaryByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcSalaryService.DeleteWcSalaryByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcSalary 更新工资单模版
// @Tags WcSalary
// @Summary 更新工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalary true "更新工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcSalary/updateWcSalary [put]
func (wcSalaryApi *WcSalaryApi) UpdateWcSalary(c *gin.Context) {
	var wcSalary weChat.WcSalary
	err := c.ShouldBindJSON(&wcSalary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcSalaryService.UpdateWcSalary(wcSalary); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcSalary 用id查询工资单模版
// @Tags WcSalary
// @Summary 用id查询工资单模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcSalary true "用id查询工资单模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcSalary/findWcSalary [get]
func (wcSalaryApi *WcSalaryApi) FindWcSalary(c *gin.Context) {
	ID := c.Query("ID")
	if rewcSalary, err := wcSalaryService.GetWcSalary(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcSalary": rewcSalary}, c)
	}
}

// GetWcSalaryList 分页获取工资单模版列表
// @Tags WcSalary
// @Summary 分页获取工资单模版列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcSalarySearch true "分页获取工资单模版列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalary/getWcSalaryList [get]
func (wcSalaryApi *WcSalaryApi) GetWcSalaryList(c *gin.Context) {
	var pageInfo weChatReq.WcSalarySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("pageInfo", pageInfo)
	if list, total, err := wcSalaryService.GetWcSalaryInfoList(pageInfo); err != nil {
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

// GetWcSalaryPublic 不需要鉴权的工资单模版接口
// @Tags WcSalary
// @Summary 不需要鉴权的工资单模版接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcSalarySearch true "分页获取工资单模版列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalary/getWcSalaryList [get]
func (wcSalaryApi *WcSalaryApi) GetWcSalaryPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的工资单模版接口信息",
	}, "获取成功", c)
}
