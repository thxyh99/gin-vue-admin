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

type WcStaffSalaryApi struct {
}

var wcStaffSalaryService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffSalaryService


// CreateWcStaffSalary 创建薪资奖金
// @Tags WcStaffSalary
// @Summary 创建薪资奖金
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffSalary true "创建薪资奖金"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffSalary/createWcStaffSalary [post]
func (wcStaffSalaryApi *WcStaffSalaryApi) CreateWcStaffSalary(c *gin.Context) {
	var wcStaffSalary weChat.WcStaffSalary
	err := c.ShouldBindJSON(&wcStaffSalary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffSalaryService.CreateWcStaffSalary(&wcStaffSalary); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffSalary 删除薪资奖金
// @Tags WcStaffSalary
// @Summary 删除薪资奖金
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffSalary true "删除薪资奖金"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffSalary/deleteWcStaffSalary [delete]
func (wcStaffSalaryApi *WcStaffSalaryApi) DeleteWcStaffSalary(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffSalaryService.DeleteWcStaffSalary(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffSalaryByIds 批量删除薪资奖金
// @Tags WcStaffSalary
// @Summary 批量删除薪资奖金
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffSalary/deleteWcStaffSalaryByIds [delete]
func (wcStaffSalaryApi *WcStaffSalaryApi) DeleteWcStaffSalaryByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffSalaryService.DeleteWcStaffSalaryByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffSalary 更新薪资奖金
// @Tags WcStaffSalary
// @Summary 更新薪资奖金
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffSalary true "更新薪资奖金"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffSalary/updateWcStaffSalary [put]
func (wcStaffSalaryApi *WcStaffSalaryApi) UpdateWcStaffSalary(c *gin.Context) {
	var wcStaffSalary weChat.WcStaffSalary
	err := c.ShouldBindJSON(&wcStaffSalary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffSalaryService.UpdateWcStaffSalary(wcStaffSalary); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffSalary 用id查询薪资奖金
// @Tags WcStaffSalary
// @Summary 用id查询薪资奖金
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffSalary true "用id查询薪资奖金"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffSalary/findWcStaffSalary [get]
func (wcStaffSalaryApi *WcStaffSalaryApi) FindWcStaffSalary(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffSalary, err := wcStaffSalaryService.GetWcStaffSalary(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffSalary": rewcStaffSalary}, c)
	}
}

// GetWcStaffSalaryList 分页获取薪资奖金列表
// @Tags WcStaffSalary
// @Summary 分页获取薪资奖金列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSalarySearch true "分页获取薪资奖金列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffSalary/getWcStaffSalaryList [get]
func (wcStaffSalaryApi *WcStaffSalaryApi) GetWcStaffSalaryList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffSalarySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffSalaryService.GetWcStaffSalaryInfoList(pageInfo); err != nil {
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

// GetWcStaffSalaryPublic 不需要鉴权的薪资奖金接口
// @Tags WcStaffSalary
// @Summary 不需要鉴权的薪资奖金接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSalarySearch true "分页获取薪资奖金列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffSalary/getWcStaffSalaryList [get]
func (wcStaffSalaryApi *WcStaffSalaryApi) GetWcStaffSalaryPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的薪资奖金接口信息",
    }, "获取成功", c)
}
