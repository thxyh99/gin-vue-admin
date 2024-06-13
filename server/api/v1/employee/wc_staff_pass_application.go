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

type WcStaffPassApplicationApi struct {
}

var wcStaffPassApplicationService = service.ServiceGroupApp.EmployeeServiceGroup.WcStaffPassApplicationService

// CreateWcStaffPassApplication 创建转正申请
// @Tags WcStaffPassApplication
// @Summary 创建转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffPassApplication true "创建转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffPassApplication/createWcStaffPassApplication [post]
func (wcStaffPassApplicationApi *WcStaffPassApplicationApi) CreateWcStaffPassApplication(c *gin.Context) {
	var wcStaffPassApplication employee.WcStaffPassApplication
	err := c.ShouldBindJSON(&wcStaffPassApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wcStaffPassApplication.CreatedBy = utils.GetUserID(c)

	if err := wcStaffPassApplicationService.CreateWcStaffPassApplication(&wcStaffPassApplication); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffPassApplication 删除转正申请
// @Tags WcStaffPassApplication
// @Summary 删除转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffPassApplication true "删除转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffPassApplication/deleteWcStaffPassApplication [delete]
func (wcStaffPassApplicationApi *WcStaffPassApplicationApi) DeleteWcStaffPassApplication(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := wcStaffPassApplicationService.DeleteWcStaffPassApplication(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffPassApplicationByIds 批量删除转正申请
// @Tags WcStaffPassApplication
// @Summary 批量删除转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffPassApplication/deleteWcStaffPassApplicationByIds [delete]
func (wcStaffPassApplicationApi *WcStaffPassApplicationApi) DeleteWcStaffPassApplicationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := wcStaffPassApplicationService.DeleteWcStaffPassApplicationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffPassApplication 更新转正申请
// @Tags WcStaffPassApplication
// @Summary 更新转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffPassApplication true "更新转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffPassApplication/updateWcStaffPassApplication [put]
func (wcStaffPassApplicationApi *WcStaffPassApplicationApi) UpdateWcStaffPassApplication(c *gin.Context) {
	var wcStaffPassApplication employee.WcStaffPassApplication
	err := c.ShouldBindJSON(&wcStaffPassApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wcStaffPassApplication.UpdatedBy = utils.GetUserID(c)

	if err := wcStaffPassApplicationService.UpdateWcStaffPassApplication(wcStaffPassApplication); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffPassApplication 用id查询转正申请
// @Tags WcStaffPassApplication
// @Summary 用id查询转正申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employee.WcStaffPassApplication true "用id查询转正申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffPassApplication/findWcStaffPassApplication [get]
func (wcStaffPassApplicationApi *WcStaffPassApplicationApi) FindWcStaffPassApplication(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffPassApplication, err := wcStaffPassApplicationService.GetWcStaffPassApplication(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffPassApplication": rewcStaffPassApplication}, c)
	}
}

// GetWcStaffPassApplicationList 分页获取转正申请列表
// @Tags WcStaffPassApplication
// @Summary 分页获取转正申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffPassApplicationSearch true "分页获取转正申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffPassApplication/getWcStaffPassApplicationList [get]
func (wcStaffPassApplicationApi *WcStaffPassApplicationApi) GetWcStaffPassApplicationList(c *gin.Context) {
	var pageInfo employeeReq.WcStaffPassApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffPassApplicationService.GetWcStaffPassApplicationInfoList(pageInfo); err != nil {
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

// GetWcStaffPassApplicationPublic 不需要鉴权的转正申请接口
// @Tags WcStaffPassApplication
// @Summary 不需要鉴权的转正申请接口
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffPassApplicationSearch true "分页获取转正申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffPassApplication/getWcStaffPassApplicationList [get]
func (wcStaffPassApplicationApi *WcStaffPassApplicationApi) GetWcStaffPassApplicationPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的转正申请接口信息",
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
func (wcStaffPassApplicationApi *WcStaffPassApplicationApi) CreateOAStaffPassApplication(c *gin.Context) {
	var wcStaffPassApplication employee.WcStaffPassApplication

	wcStaffPassApplication.StaffId = 12
	wcStaffPassApplication.Title = "员工转正申请" + "-" + "刘永波"
	wcStaffPassApplication.EmploymentDate = "2024-05-05"
	wcStaffPassApplication.JobDepartment = "1644f7a70f2503bb7bf923a4c479aff1"
	wcStaffPassApplication.JobDepartment = "容大集团_集团总部_信息技术部"
	wcStaffPassApplication.SourcePosition = "软件工程师"
	wcStaffPassApplication.Education = "大专"
	//wcStaffEmploymentApplication.IsCeopass =
	//wcStaffEmploymentApplication.IsBodychecknormal =
	wcStaffPassApplication.SourceLevel = "职级"
	wcStaffPassApplication.TryPeriod = "1"

	// 调用外部服务或库之前添加错误处理
	landrayOa := utils.LandrayOa{}
	oaId, err := landrayOa.CreateOAPassApplication(wcStaffPassApplication)
	if err != nil {
		global.GVA_LOG.Error("创建员工转正申请失败", zap.Error(err))
		response.FailWithMessage("创建员工转正申请失败", c)
		return
	}

	wcStaffPassApplication.OaId = oaId
	wcStaffPassApplication.OaStatus = 10
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
	wcStaffPassApplication.CreatedBy = utils.GetUserID(c)
	if err := wcStaffPassApplicationService.CreateOAPassApplication(&wcStaffPassApplication); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
