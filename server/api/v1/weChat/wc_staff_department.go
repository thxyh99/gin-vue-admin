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

type WcStaffDepartmentApi struct {
}

var wcStaffDepartmentService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffDepartmentService


// CreateWcStaffDepartment 创建员工部门管理
// @Tags WcStaffDepartment
// @Summary 创建员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffDepartment true "创建员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffDepartment/createWcStaffDepartment [post]
func (wcStaffDepartmentApi *WcStaffDepartmentApi) CreateWcStaffDepartment(c *gin.Context) {
	var wcStaffDepartment weChat.WcStaffDepartment
	err := c.ShouldBindJSON(&wcStaffDepartment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffDepartmentService.CreateWcStaffDepartment(&wcStaffDepartment); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffDepartment 删除员工部门管理
// @Tags WcStaffDepartment
// @Summary 删除员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffDepartment true "删除员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffDepartment/deleteWcStaffDepartment [delete]
func (wcStaffDepartmentApi *WcStaffDepartmentApi) DeleteWcStaffDepartment(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffDepartmentService.DeleteWcStaffDepartment(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffDepartmentByIds 批量删除员工部门管理
// @Tags WcStaffDepartment
// @Summary 批量删除员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffDepartment/deleteWcStaffDepartmentByIds [delete]
func (wcStaffDepartmentApi *WcStaffDepartmentApi) DeleteWcStaffDepartmentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffDepartmentService.DeleteWcStaffDepartmentByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffDepartment 更新员工部门管理
// @Tags WcStaffDepartment
// @Summary 更新员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffDepartment true "更新员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffDepartment/updateWcStaffDepartment [put]
func (wcStaffDepartmentApi *WcStaffDepartmentApi) UpdateWcStaffDepartment(c *gin.Context) {
	var wcStaffDepartment weChat.WcStaffDepartment
	err := c.ShouldBindJSON(&wcStaffDepartment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffDepartmentService.UpdateWcStaffDepartment(wcStaffDepartment); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffDepartment 用id查询员工部门管理
// @Tags WcStaffDepartment
// @Summary 用id查询员工部门管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffDepartment true "用id查询员工部门管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffDepartment/findWcStaffDepartment [get]
func (wcStaffDepartmentApi *WcStaffDepartmentApi) FindWcStaffDepartment(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffDepartment, err := wcStaffDepartmentService.GetWcStaffDepartment(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffDepartment": rewcStaffDepartment}, c)
	}
}

// GetWcStaffDepartmentList 分页获取员工部门管理列表
// @Tags WcStaffDepartment
// @Summary 分页获取员工部门管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffDepartmentSearch true "分页获取员工部门管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffDepartment/getWcStaffDepartmentList [get]
func (wcStaffDepartmentApi *WcStaffDepartmentApi) GetWcStaffDepartmentList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffDepartmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffDepartmentService.GetWcStaffDepartmentInfoList(pageInfo); err != nil {
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

// GetWcStaffDepartmentPublic 不需要鉴权的员工部门管理接口
// @Tags WcStaffDepartment
// @Summary 不需要鉴权的员工部门管理接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffDepartmentSearch true "分页获取员工部门管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffDepartment/getWcStaffDepartmentList [get]
func (wcStaffDepartmentApi *WcStaffDepartmentApi) GetWcStaffDepartmentPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的员工部门管理接口信息",
    }, "获取成功", c)
}
