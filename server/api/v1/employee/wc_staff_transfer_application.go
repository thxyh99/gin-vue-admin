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
	if err := wcStaffTransferApplicationService.DeleteWcStaffTransferApplication(ID,userID); err != nil {
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
	if err := wcStaffTransferApplicationService.DeleteWcStaffTransferApplicationByIds(IDs,userID); err != nil {
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
