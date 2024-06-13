package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WcStaffLeaveApplicationApi struct {
}

var wcStaffLeaveApplicationService = service.ServiceGroupApp.EmployeeServiceGroup.WcStaffLeaveApplicationService

// CreateWcStaffLeaveApplication 创建离职申请
// @Tags WcStaffLeaveApplication
// @Summary 创建离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffLeaveApplication true "创建离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffLeaveApplication/createWcStaffLeaveApplication [post]
func (wcStaffLeaveApplicationApi *WcStaffLeaveApplicationApi) CreateWcStaffLeaveApplication(c *gin.Context) {
	var wcStaffLeaveApplication employee.WcStaffLeaveApplication
	err := c.ShouldBindJSON(&wcStaffLeaveApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wcStaffLeaveApplication.CreatedBy = utils.GetUserID(c)

	if err := wcStaffLeaveApplicationService.CreateWcStaffLeaveApplication(&wcStaffLeaveApplication); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffLeaveApplication 删除离职申请
// @Tags WcStaffLeaveApplication
// @Summary 删除离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffLeaveApplication true "删除离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffLeaveApplication/deleteWcStaffLeaveApplication [delete]
func (wcStaffLeaveApplicationApi *WcStaffLeaveApplicationApi) DeleteWcStaffLeaveApplication(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := wcStaffLeaveApplicationService.DeleteWcStaffLeaveApplication(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffLeaveApplicationByIds 批量删除离职申请
// @Tags WcStaffLeaveApplication
// @Summary 批量删除离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffLeaveApplication/deleteWcStaffLeaveApplicationByIds [delete]
func (wcStaffLeaveApplicationApi *WcStaffLeaveApplicationApi) DeleteWcStaffLeaveApplicationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := wcStaffLeaveApplicationService.DeleteWcStaffLeaveApplicationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffLeaveApplication 更新离职申请
// @Tags WcStaffLeaveApplication
// @Summary 更新离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffLeaveApplication true "更新离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffLeaveApplication/updateWcStaffLeaveApplication [put]
func (wcStaffLeaveApplicationApi *WcStaffLeaveApplicationApi) UpdateWcStaffLeaveApplication(c *gin.Context) {
	var wcStaffLeaveApplication employee.WcStaffLeaveApplication
	err := c.ShouldBindJSON(&wcStaffLeaveApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wcStaffLeaveApplication.UpdatedBy = utils.GetUserID(c)

	if err := wcStaffLeaveApplicationService.UpdateWcStaffLeaveApplication(wcStaffLeaveApplication); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffLeaveApplication 用id查询离职申请
// @Tags WcStaffLeaveApplication
// @Summary 用id查询离职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employee.WcStaffLeaveApplication true "用id查询离职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffLeaveApplication/findWcStaffLeaveApplication [get]
func (wcStaffLeaveApplicationApi *WcStaffLeaveApplicationApi) FindWcStaffLeaveApplication(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffLeaveApplication, err := wcStaffLeaveApplicationService.GetWcStaffLeaveApplication(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffLeaveApplication": rewcStaffLeaveApplication}, c)
	}
}

// GetWcStaffLeaveApplicationList 分页获取离职申请列表
// @Tags WcStaffLeaveApplication
// @Summary 分页获取离职申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffLeaveApplicationSearch true "分页获取离职申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffLeaveApplication/getWcStaffLeaveApplicationList [get]
func (wcStaffLeaveApplicationApi *WcStaffLeaveApplicationApi) GetWcStaffLeaveApplicationList(c *gin.Context) {
	var pageInfo employeeReq.WcStaffLeaveApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffLeaveApplicationService.GetWcStaffLeaveApplicationInfoList(pageInfo); err != nil {
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

// GetWcStaffLeaveApplicationPublic 不需要鉴权的离职申请接口
// @Tags WcStaffLeaveApplication
// @Summary 不需要鉴权的离职申请接口
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffLeaveApplicationSearch true "分页获取离职申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffLeaveApplication/getWcStaffLeaveApplicationList [get]
func (wcStaffLeaveApplicationApi *WcStaffLeaveApplicationApi) GetWcStaffLeaveApplicationPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的离职申请接口信息",
	}, "获取成功", c)
}

// CreateOAStaffEmploymentApplication 创建OA入职申请
// @Tags OAStaffEmploymentApplication
// @Summary 创建OA入职申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffEmploymentApplication true "创建OA入职申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffEmploymentApplication/createOAStaffEmploymentApplication [get]
func (wcStaffLeaveApplicationApi *WcStaffLeaveApplicationApi) CreateOAStaffLeaveApplication(c *gin.Context) {
	var wcStaffLeaveApplication employee.WcStaffLeaveApplication

	wcStaffLeaveApplication.StaffId = 12
	wcStaffLeaveApplication.Title = "离职审批" + "-" + "刘永波"
	wcStaffLeaveApplication.LeaveDate = "2024-06-05"
	//wcStaffLeaveApplication.JobDepartment = "1644f7a70f2503bb7bf923a4c479aff1"
	wcStaffLeaveApplication.JobDepartment = "容大集团_集团总部_信息技术部"
	wcStaffLeaveApplication.LeaveType = "协商解除"
	wcStaffLeaveApplication.LeaveResult = "离职测试"

	// 调用外部服务或库之前添加错误处理
	landrayOa := utils.LandrayOa{}
	oaId, err := landrayOa.CreateOALeaveApplication(wcStaffLeaveApplication)
	if err != nil {
		global.GVA_LOG.Error("创建员工离职申请失败", zap.Error(err))
		response.FailWithMessage("创建员工离职申请失败", c)
		return
	}

	wcStaffLeaveApplication.OaId = oaId
	wcStaffLeaveApplication.OaStatus = 10
	/*
		// 增加数据验证逻辑
		if err := validateApplication(&wcStaffPassApplication); err != nil {
			global.GVA_LOG.Error("数据验证失败", zap.Error(err))
			response.FailWithMessage("数据验证失败", c)
			return
		}
			// 绑定JSON数据并处理错误
			if err := c.ShouldBindJSON(wcStaffEmploymentApplication); err != nil {
				global.GVA_LOG.Error("JSON绑定失败", zap.Error(err))
				response.FailWithMessage(err.Error(), c)
				return
			}
	*/
	wcStaffLeaveApplication.CreatedBy = utils.GetUserID(c)
	if err := wcStaffLeaveApplicationService.CreateWcStaffLeaveApplication(&wcStaffLeaveApplication); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
