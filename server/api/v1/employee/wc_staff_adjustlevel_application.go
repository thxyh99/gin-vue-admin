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

type WcStaffAdjustlevelApplicationApi struct {
}

var wcStaffAdjustlevelApplicationService = service.ServiceGroupApp.EmployeeServiceGroup.WcStaffAdjustlevelApplicationService

// CreateWcStaffAdjustlevelApplication 创建调级申请
// @Tags WcStaffAdjustlevelApplication
// @Summary 创建调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffAdjustlevelApplication true "创建调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffAdjustlevelApplication/createWcStaffAdjustlevelApplication [post]
func (wcStaffAdjustlevelApplicationApi *WcStaffAdjustlevelApplicationApi) CreateWcStaffAdjustlevelApplication(c *gin.Context) {
	var wcStaffAdjustlevelApplication employee.WcStaffAdjustlevelApplication
	err := c.ShouldBindJSON(&wcStaffAdjustlevelApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wcStaffAdjustlevelApplication.CreatedBy = utils.GetUserID(c)

	if err := wcStaffAdjustlevelApplicationService.CreateWcStaffAdjustlevelApplication(&wcStaffAdjustlevelApplication); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffAdjustlevelApplication 删除调级申请
// @Tags WcStaffAdjustlevelApplication
// @Summary 删除调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffAdjustlevelApplication true "删除调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffAdjustlevelApplication/deleteWcStaffAdjustlevelApplication [delete]
func (wcStaffAdjustlevelApplicationApi *WcStaffAdjustlevelApplicationApi) DeleteWcStaffAdjustlevelApplication(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := wcStaffAdjustlevelApplicationService.DeleteWcStaffAdjustlevelApplication(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffAdjustlevelApplicationByIds 批量删除调级申请
// @Tags WcStaffAdjustlevelApplication
// @Summary 批量删除调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffAdjustlevelApplication/deleteWcStaffAdjustlevelApplicationByIds [delete]
func (wcStaffAdjustlevelApplicationApi *WcStaffAdjustlevelApplicationApi) DeleteWcStaffAdjustlevelApplicationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := wcStaffAdjustlevelApplicationService.DeleteWcStaffAdjustlevelApplicationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffAdjustlevelApplication 更新调级申请
// @Tags WcStaffAdjustlevelApplication
// @Summary 更新调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffAdjustlevelApplication true "更新调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffAdjustlevelApplication/updateWcStaffAdjustlevelApplication [put]
func (wcStaffAdjustlevelApplicationApi *WcStaffAdjustlevelApplicationApi) UpdateWcStaffAdjustlevelApplication(c *gin.Context) {
	var wcStaffAdjustlevelApplication employee.WcStaffAdjustlevelApplication
	err := c.ShouldBindJSON(&wcStaffAdjustlevelApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wcStaffAdjustlevelApplication.UpdatedBy = utils.GetUserID(c)

	if err := wcStaffAdjustlevelApplicationService.UpdateWcStaffAdjustlevelApplication(wcStaffAdjustlevelApplication); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffAdjustlevelApplication 用id查询调级申请
// @Tags WcStaffAdjustlevelApplication
// @Summary 用id查询调级申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employee.WcStaffAdjustlevelApplication true "用id查询调级申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffAdjustlevelApplication/findWcStaffAdjustlevelApplication [get]
func (wcStaffAdjustlevelApplicationApi *WcStaffAdjustlevelApplicationApi) FindWcStaffAdjustlevelApplication(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffAdjustlevelApplication, err := wcStaffAdjustlevelApplicationService.GetWcStaffAdjustlevelApplication(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffAdjustlevelApplication": rewcStaffAdjustlevelApplication}, c)
	}
}

// GetWcStaffAdjustlevelApplicationList 分页获取调级申请列表
// @Tags WcStaffAdjustlevelApplication
// @Summary 分页获取调级申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffAdjustlevelApplicationSearch true "分页获取调级申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAdjustlevelApplication/getWcStaffAdjustlevelApplicationList [get]
func (wcStaffAdjustlevelApplicationApi *WcStaffAdjustlevelApplicationApi) GetWcStaffAdjustlevelApplicationList(c *gin.Context) {
	var pageInfo employeeReq.WcStaffAdjustlevelApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffAdjustlevelApplicationService.GetWcStaffAdjustlevelApplicationInfoList(pageInfo); err != nil {
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

// GetWcStaffAdjustlevelApplicationPublic 不需要鉴权的调级申请接口
// @Tags WcStaffAdjustlevelApplication
// @Summary 不需要鉴权的调级申请接口
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffAdjustlevelApplicationSearch true "分页获取调级申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffAdjustlevelApplication/getWcStaffAdjustlevelApplicationList [get]
func (wcStaffAdjustlevelApplicationApi *WcStaffAdjustlevelApplicationApi) GetWcStaffAdjustlevelApplicationPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的调级申请接口信息",
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
func (wcStaffAdjustlevelApplicationApi *WcStaffAdjustlevelApplicationApi) CreateOAStaffAdjustlevelApplication(c *gin.Context) {
	var wcStaffAdjustlevelApplication employee.WcStaffAdjustlevelApplication

	wcStaffAdjustlevelApplication.StaffId = 12
	wcStaffAdjustlevelApplication.Title = "调级审批" + "-" + "刘永波"
	//wcStaffAdjustlevelApplication.EmploymentDate = "2024-05-05"
	//wcStaffAdjustlevelApplication.JobDepartment = "1644f7a70f2503bb7bf923a4c479aff1"
	wcStaffAdjustlevelApplication.SourceDepartment = "容大集团_集团总部_信息技术部"
	wcStaffAdjustlevelApplication.SourcePosition = "软件工程师"
	wcStaffAdjustlevelApplication.Memo = "调级测试"

	// 调用外部服务或库之前添加错误处理
	landrayOa := utils.LandrayOa{}
	oaId, err := landrayOa.CreateOAAdjustlevelApplication(wcStaffAdjustlevelApplication)
	if err != nil {
		global.GVA_LOG.Error("创建员工转正申请失败", zap.Error(err))
		response.FailWithMessage("创建员工转正申请失败", c)
		return
	}

	wcStaffAdjustlevelApplication.OaId = oaId
	wcStaffAdjustlevelApplication.OaStatus = 10
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
	wcStaffAdjustlevelApplication.CreatedBy = utils.GetUserID(c)
	if err := wcStaffAdjustlevelApplicationService.CreateWcStaffAdjustlevelApplication(&wcStaffAdjustlevelApplication); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
