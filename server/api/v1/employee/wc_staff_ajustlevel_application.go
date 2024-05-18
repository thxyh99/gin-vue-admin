package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/employee"
    employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type WcStaffAjustlevelApplicationApi struct {
}

var wcStaffAjustlevelApplicationService = service.ServiceGroupApp.EmployeeServiceGroup.WcStaffAjustlevelApplicationService


// CreateWcStaffAjustlevelApplication 创建调级申请
// @Tags WcStaffAjustlevelApplication
// @Summary 创建调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffAjustlevelApplication true "创建调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffAjustlevelApplication/createWcStaffAjustlevelApplication [post]
func (wcStaffAjustlevelApplicationApi *WcStaffAjustlevelApplicationApi) CreateWcStaffAjustlevelApplication(c *gin.Context) {
	var wcStaffAjustlevelApplication employee.WcStaffAjustlevelApplication
	err := c.ShouldBindJSON(&wcStaffAjustlevelApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wcStaffAjustlevelApplication.CreatedBy = utils.GetUserID(c)

	if err := wcStaffAjustlevelApplicationService.CreateWcStaffAjustlevelApplication(&wcStaffAjustlevelApplication); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffAjustlevelApplication 删除调级申请
// @Tags WcStaffAjustlevelApplication
// @Summary 删除调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffAjustlevelApplication true "删除调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAjustlevelApplication/deleteWcStaffAjustlevelApplication [delete]
func (wcStaffAjustlevelApplicationApi *WcStaffAjustlevelApplicationApi) DeleteWcStaffAjustlevelApplication(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := wcStaffAjustlevelApplicationService.DeleteWcStaffAjustlevelApplication(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffAjustlevelApplicationByIds 批量删除调级申请
// @Tags WcStaffAjustlevelApplication
// @Summary 批量删除调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffAjustlevelApplication/deleteWcStaffAjustlevelApplicationByIds [delete]
func (wcStaffAjustlevelApplicationApi *WcStaffAjustlevelApplicationApi) DeleteWcStaffAjustlevelApplicationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := wcStaffAjustlevelApplicationService.DeleteWcStaffAjustlevelApplicationByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffAjustlevelApplication 更新调级申请
// @Tags WcStaffAjustlevelApplication
// @Summary 更新调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffAjustlevelApplication true "更新调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffAjustlevelApplication/updateWcStaffAjustlevelApplication [put]
func (wcStaffAjustlevelApplicationApi *WcStaffAjustlevelApplicationApi) UpdateWcStaffAjustlevelApplication(c *gin.Context) {
	var wcStaffAjustlevelApplication employee.WcStaffAjustlevelApplication
	err := c.ShouldBindJSON(&wcStaffAjustlevelApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wcStaffAjustlevelApplication.UpdatedBy = utils.GetUserID(c)

	if err := wcStaffAjustlevelApplicationService.UpdateWcStaffAjustlevelApplication(wcStaffAjustlevelApplication); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffAjustlevelApplication 用id查询调级申请
// @Tags WcStaffAjustlevelApplication
// @Summary 用id查询调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employee.WcStaffAjustlevelApplication true "用id查询调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffAjustlevelApplication/findWcStaffAjustlevelApplication [get]
func (wcStaffAjustlevelApplicationApi *WcStaffAjustlevelApplicationApi) FindWcStaffAjustlevelApplication(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffAjustlevelApplication, err := wcStaffAjustlevelApplicationService.GetWcStaffAjustlevelApplication(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffAjustlevelApplication": rewcStaffAjustlevelApplication}, c)
	}
}

// GetWcStaffAjustlevelApplicationList 分页获取调级申请列表
// @Tags WcStaffAjustlevelApplication
// @Summary 分页获取调级申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffAjustlevelApplicationSearch true "分页获取调级申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAjustlevelApplication/getWcStaffAjustlevelApplicationList [get]
func (wcStaffAjustlevelApplicationApi *WcStaffAjustlevelApplicationApi) GetWcStaffAjustlevelApplicationList(c *gin.Context) {
	var pageInfo employeeReq.WcStaffAjustlevelApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffAjustlevelApplicationService.GetWcStaffAjustlevelApplicationInfoList(pageInfo); err != nil {
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

// GetWcStaffAjustlevelApplicationPublic 不需要鉴权的调级申请接口
// @Tags WcStaffAjustlevelApplication
// @Summary 不需要鉴权的调级申请接口
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffAjustlevelApplicationSearch true "分页获取调级申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAjustlevelApplication/getWcStaffAjustlevelApplicationList [get]
func (wcStaffAjustlevelApplicationApi *WcStaffAjustlevelApplicationApi) GetWcStaffAjustlevelApplicationPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的调级申请接口信息",
    }, "获取成功", c)
}
