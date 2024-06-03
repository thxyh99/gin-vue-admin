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

type WcSalaryTypeFieldApi struct {
}

var wcSalaryTypeFieldService = service.ServiceGroupApp.WeChatServiceGroup.WcSalaryTypeFieldService


// CreateWcSalaryTypeField 创建工资类型款项
// @Tags WcSalaryTypeField
// @Summary 创建工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalaryTypeField true "创建工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcSalaryTypeField/createWcSalaryTypeField [post]
func (wcSalaryTypeFieldApi *WcSalaryTypeFieldApi) CreateWcSalaryTypeField(c *gin.Context) {
	var wcSalaryTypeField weChat.WcSalaryTypeField
	err := c.ShouldBindJSON(&wcSalaryTypeField)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcSalaryTypeFieldService.CreateWcSalaryTypeField(&wcSalaryTypeField); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcSalaryTypeField 删除工资类型款项
// @Tags WcSalaryTypeField
// @Summary 删除工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalaryTypeField true "删除工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcSalaryTypeField/deleteWcSalaryTypeField [delete]
func (wcSalaryTypeFieldApi *WcSalaryTypeFieldApi) DeleteWcSalaryTypeField(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcSalaryTypeFieldService.DeleteWcSalaryTypeField(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcSalaryTypeFieldByIds 批量删除工资类型款项
// @Tags WcSalaryTypeField
// @Summary 批量删除工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcSalaryTypeField/deleteWcSalaryTypeFieldByIds [delete]
func (wcSalaryTypeFieldApi *WcSalaryTypeFieldApi) DeleteWcSalaryTypeFieldByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcSalaryTypeFieldService.DeleteWcSalaryTypeFieldByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcSalaryTypeField 更新工资类型款项
// @Tags WcSalaryTypeField
// @Summary 更新工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcSalaryTypeField true "更新工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcSalaryTypeField/updateWcSalaryTypeField [put]
func (wcSalaryTypeFieldApi *WcSalaryTypeFieldApi) UpdateWcSalaryTypeField(c *gin.Context) {
	var wcSalaryTypeField weChat.WcSalaryTypeField
	err := c.ShouldBindJSON(&wcSalaryTypeField)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcSalaryTypeFieldService.UpdateWcSalaryTypeField(wcSalaryTypeField); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcSalaryTypeField 用id查询工资类型款项
// @Tags WcSalaryTypeField
// @Summary 用id查询工资类型款项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcSalaryTypeField true "用id查询工资类型款项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcSalaryTypeField/findWcSalaryTypeField [get]
func (wcSalaryTypeFieldApi *WcSalaryTypeFieldApi) FindWcSalaryTypeField(c *gin.Context) {
	ID := c.Query("ID")
	if rewcSalaryTypeField, err := wcSalaryTypeFieldService.GetWcSalaryTypeField(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcSalaryTypeField": rewcSalaryTypeField}, c)
	}
}

// GetWcSalaryTypeFieldList 分页获取工资类型款项列表
// @Tags WcSalaryTypeField
// @Summary 分页获取工资类型款项列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcSalaryTypeFieldSearch true "分页获取工资类型款项列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalaryTypeField/getWcSalaryTypeFieldList [get]
func (wcSalaryTypeFieldApi *WcSalaryTypeFieldApi) GetWcSalaryTypeFieldList(c *gin.Context) {
	var pageInfo weChatReq.WcSalaryTypeFieldSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcSalaryTypeFieldService.GetWcSalaryTypeFieldInfoList(pageInfo); err != nil {
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

// GetWcSalaryTypeFieldPublic 不需要鉴权的工资类型款项接口
// @Tags WcSalaryTypeField
// @Summary 不需要鉴权的工资类型款项接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcSalaryTypeFieldSearch true "分页获取工资类型款项列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcSalaryTypeField/getWcSalaryTypeFieldList [get]
func (wcSalaryTypeFieldApi *WcSalaryTypeFieldApi) GetWcSalaryTypeFieldPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的工资类型款项接口信息",
    }, "获取成功", c)
}
