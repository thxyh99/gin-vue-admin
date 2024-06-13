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

type WcStaffTransferApplicationApi struct {
}

var wcStaffTransferApplicationService = service.ServiceGroupApp.EmployeeServiceGroup.WcStaffTransferApplicationService

// CreateWcStaffTransferApplication 创建调动申请
// @Tags WcStaffTransferApplication
// @Summary 创建调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffTransferApplication true "创建调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffTransferApplication/createWcStaffTransferApplication [post]
func (wcStaffTransferApplicationApi *WcStaffTransferApplicationApi) CreateWcStaffTransferApplication(c *gin.Context) {
	var wcStaffTransferApplication employee.WcStaffTransferApplication
	err := c.ShouldBindJSON(&wcStaffTransferApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wcStaffTransferApplication.CreatedBy = utils.GetUserID(c)

	if err := wcStaffTransferApplicationService.CreateWcStaffTransferApplication(&wcStaffTransferApplication); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffTransferApplication 删除调动申请
// @Tags WcStaffTransferApplication
// @Summary 删除调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffTransferApplication true "删除调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffTransferApplication/deleteWcStaffTransferApplication [delete]
func (wcStaffTransferApplicationApi *WcStaffTransferApplicationApi) DeleteWcStaffTransferApplication(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := wcStaffTransferApplicationService.DeleteWcStaffTransferApplication(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffTransferApplicationByIds 批量删除调动申请
// @Tags WcStaffTransferApplication
// @Summary 批量删除调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffTransferApplication/deleteWcStaffTransferApplicationByIds [delete]
func (wcStaffTransferApplicationApi *WcStaffTransferApplicationApi) DeleteWcStaffTransferApplicationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := wcStaffTransferApplicationService.DeleteWcStaffTransferApplicationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffTransferApplication 更新调动申请
// @Tags WcStaffTransferApplication
// @Summary 更新调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcStaffTransferApplication true "更新调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffTransferApplication/updateWcStaffTransferApplication [put]
func (wcStaffTransferApplicationApi *WcStaffTransferApplicationApi) UpdateWcStaffTransferApplication(c *gin.Context) {
	var wcStaffTransferApplication employee.WcStaffTransferApplication
	err := c.ShouldBindJSON(&wcStaffTransferApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wcStaffTransferApplication.UpdatedBy = utils.GetUserID(c)

	if err := wcStaffTransferApplicationService.UpdateWcStaffTransferApplication(wcStaffTransferApplication); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffTransferApplication 用id查询调动申请
// @Tags WcStaffTransferApplication
// @Summary 用id查询调动申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employee.WcStaffTransferApplication true "用id查询调动申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffTransferApplication/findWcStaffTransferApplication [get]
func (wcStaffTransferApplicationApi *WcStaffTransferApplicationApi) FindWcStaffTransferApplication(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffTransferApplication, err := wcStaffTransferApplicationService.GetWcStaffTransferApplication(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffTransferApplication": rewcStaffTransferApplication}, c)
	}
}

// GetWcStaffTransferApplicationList 分页获取调动申请列表
// @Tags WcStaffTransferApplication
// @Summary 分页获取调动申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffTransferApplicationSearch true "分页获取调动申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffTransferApplication/getWcStaffTransferApplicationList [get]
func (wcStaffTransferApplicationApi *WcStaffTransferApplicationApi) GetWcStaffTransferApplicationList(c *gin.Context) {
	var pageInfo employeeReq.WcStaffTransferApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffTransferApplicationService.GetWcStaffTransferApplicationInfoList(pageInfo); err != nil {
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

// GetWcStaffTransferApplicationPublic 不需要鉴权的调动申请接口
// @Tags WcStaffTransferApplication
// @Summary 不需要鉴权的调动申请接口
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcStaffTransferApplicationSearch true "分页获取调动申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffTransferApplication/getWcStaffTransferApplicationList [get]
func (wcStaffTransferApplicationApi *WcStaffTransferApplicationApi) GetWcStaffTransferApplicationPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的调动申请接口信息",
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
func (wcStaffTransferApplicationApi *WcStaffTransferApplicationApi) CreateOAStaffTransferApplication(c *gin.Context) {
	var wcStaffTransferApplication employee.WcStaffTransferApplication

	//wcStaffTransferApplication.TransferStaff = "12"
	wcStaffTransferApplication.Title = "调动申请" + "-" + "刘永波"
	wcStaffTransferApplication.EmploymentDate = "2024-06-03"
	wcStaffTransferApplication.TransferType = "2"
	wcStaffTransferApplication.TransferStaff = "{'Id': '165859c24b5e99309ab2c7349e7bb6ac', 'name': '刘永波'}"
	wcStaffTransferApplication.SourceDepartment = "{'Id': '1644f7a70f2503bb7bf923a4c479aff1', 'name': '容大集团_集团总部_信息技术部'}"
	wcStaffTransferApplication.SourcePosition = "软件工程师"
	wcStaffTransferApplication.NewDepartment = "{'Id': '1644f7a70a2554c24bcefe44be68c55e', 'name': '容大集团_集团总部_综合行政部'}"
	wcStaffTransferApplication.NewPosition = "行政助理"
	wcStaffTransferApplication.TransferResult = "调动事由"
	wcStaffTransferApplication.SourceResult = "调出部门意见"
	wcStaffTransferApplication.ToResult = "调入部门意见"
	wcStaffTransferApplication.ToDate = "2024-06-11"
	wcStaffTransferApplication.InspectionPerion = "1"

	// 调用外部服务或库之前添加错误处理
	landrayOa := utils.LandrayOa{}
	wcStaffTransferApplication.OaStatus = 10
	oaId, err := landrayOa.CreateOATransferApplication(wcStaffTransferApplication)
	if err != nil {
		global.GVA_LOG.Error("创建员工调动申请失败", zap.Error(err))
		response.FailWithMessage("创建员工调动申请失败", c)
		return
	}

	wcStaffTransferApplication.OaId = oaId
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
	wcStaffTransferApplication.CreatedBy = utils.GetUserID(c)
	if err := wcStaffTransferApplicationService.CreateWcStaffTransferApplication(&wcStaffTransferApplication); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
