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

type WcStaffEmploymentApplicationApi struct {
}

var wcStaffEmploymentApplicationService = service.ServiceGroupApp.EmployeeServiceGroup.WcStaffEmploymentApplicationService


// CreateWcStaffEmploymentApplication 创建入职申请
// @Tags WcStaffEmploymentApplication
// @Summary 创建入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffEmploymentApplication true "创建入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffEmploymentApplication/createWcStaffEmploymentApplication [post]
func (wcStaffEmploymentApplicationApi *WcStaffEmploymentApplicationApi) CreateWcStaffEmploymentApplication(c *gin.Context) {
	var wcStaffEmploymentApplication employee.WcStaffEmploymentApplication
	err := c.ShouldBindJSON(&wcStaffEmploymentApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wcStaffEmploymentApplication.CreatedBy = utils.GetUserID(c)

	if err := wcStaffEmploymentApplicationService.CreateWcStaffEmploymentApplication(&wcStaffEmploymentApplication); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffEmploymentApplication 删除入职申请
// @Tags WcStaffEmploymentApplication
// @Summary 删除入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffEmploymentApplication true "删除入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffEmploymentApplication/deleteWcStaffEmploymentApplication [delete]
func (wcStaffEmploymentApplicationApi *WcStaffEmploymentApplicationApi) DeleteWcStaffEmploymentApplication(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := wcStaffEmploymentApplicationService.DeleteWcStaffEmploymentApplication(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffEmploymentApplicationByIds 批量删除入职申请
// @Tags WcStaffEmploymentApplication
// @Summary 批量删除入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffEmploymentApplication/deleteWcStaffEmploymentApplicationByIds [delete]
func (wcStaffEmploymentApplicationApi *WcStaffEmploymentApplicationApi) DeleteWcStaffEmploymentApplicationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := wcStaffEmploymentApplicationService.DeleteWcStaffEmploymentApplicationByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffEmploymentApplication 更新入职申请
// @Tags WcStaffEmploymentApplication
// @Summary 更新入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffEmploymentApplication true "更新入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffEmploymentApplication/updateWcStaffEmploymentApplication [put]
func (wcStaffEmploymentApplicationApi *WcStaffEmploymentApplicationApi) UpdateWcStaffEmploymentApplication(c *gin.Context) {
	var wcStaffEmploymentApplication employee.WcStaffEmploymentApplication
	err := c.ShouldBindJSON(&wcStaffEmploymentApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wcStaffEmploymentApplication.UpdatedBy = utils.GetUserID(c)

	if err := wcStaffEmploymentApplicationService.UpdateWcStaffEmploymentApplication(wcStaffEmploymentApplication); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffEmploymentApplication 用id查询入职申请
// @Tags WcStaffEmploymentApplication
// @Summary 用id查询入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employee.WcStaffEmploymentApplication true "用id查询入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffEmploymentApplication/findWcStaffEmploymentApplication [get]
func (wcStaffEmploymentApplicationApi *WcStaffEmploymentApplicationApi) FindWcStaffEmploymentApplication(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffEmploymentApplication, err := wcStaffEmploymentApplicationService.GetWcStaffEmploymentApplication(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffEmploymentApplication": rewcStaffEmploymentApplication}, c)
	}
}

// GetWcStaffEmploymentApplicationList 分页获取入职申请列表
// @Tags WcStaffEmploymentApplication
// @Summary 分页获取入职申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffEmploymentApplicationSearch true "分页获取入职申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffEmploymentApplication/getWcStaffEmploymentApplicationList [get]
func (wcStaffEmploymentApplicationApi *WcStaffEmploymentApplicationApi) GetWcStaffEmploymentApplicationList(c *gin.Context) {
	var pageInfo employeeReq.WcStaffEmploymentApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffEmploymentApplicationService.GetWcStaffEmploymentApplicationInfoList(pageInfo); err != nil {
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

// GetWcStaffEmploymentApplicationPublic 不需要鉴权的入职申请接口
// @Tags WcStaffEmploymentApplication
// @Summary 不需要鉴权的入职申请接口
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffEmploymentApplicationSearch true "分页获取入职申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffEmploymentApplication/getWcStaffEmploymentApplicationList [get]
func (wcStaffEmploymentApplicationApi *WcStaffEmploymentApplicationApi) GetWcStaffEmploymentApplicationPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的入职申请接口信息",
    }, "获取成功", c)
}
