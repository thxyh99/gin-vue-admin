package employee

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

type WcThirdOaPushApi struct {
}

var wcThirdOaPushService = service.ServiceGroupApp.EmployeeServiceGroup.WcThirdOaPushService

// CreateWcThirdOaPush 创建OA回调日志
// @Tags WcThirdOaPush
// @Summary 创建OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcThirdOaPush true "创建OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcThirdOaPush/createWcThirdOaPush [post]
func (wcThirdOaPushApi *WcThirdOaPushApi) CreateWcThirdOaPush(c *gin.Context) {
	var wcThirdOaPush employee.WcThirdOaPush
	err := c.ShouldBindJSON(&wcThirdOaPush)
	fmt.Println(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	flowType := GetFlowTypeId(wcThirdOaPush.FdTmplId)
	switch flowType {
	case "employment":
		wcStaffEmploymentApplicationService.UpdateWcStaffEmploymentApplicationOaStatus(wcThirdOaPush.FdId, wcThirdOaPush.Status)
		if wcThirdOaPush.Status == 30 {
			wcStaffEmploymentApplication, err := wcStaffEmploymentApplicationService.GetWcStaffEmploymentApplicationByOA(wcThirdOaPush.FdId)
			if err == nil {

				wcStaffEmploymentApplicationService.UpdateStaffEmploymentStatus(wcStaffEmploymentApplication)
			}
		}
		break
	case "transfer":
		wcStaffTransferApplicationService.UpdateWcStaffTransferApplicationOaStatus(wcThirdOaPush.FdId, wcThirdOaPush.Status)
		wcStaffTransferApplication, err := wcStaffTransferApplicationService.GetWcStaffTransferApplicationByOA(wcThirdOaPush.FdId)
		if err == nil {
			wcStaffTransferApplicationService.UpdateWcStaffTransferApplicationByOA(wcStaffTransferApplication)
		}
		break
	case "leave":
		wcStaffLeaveApplicationService.UpdateWcStaffLeaveApplicationOaStatus(wcThirdOaPush.FdId, wcThirdOaPush.Status)
		wcStaffLeaveApplication, err := wcStaffLeaveApplicationService.GetWcStaffLeaveApplicationByOA(wcThirdOaPush.FdId)
		if err == nil {
			wcStaffLeaveApplicationService.UpdateWcStaffLeaveApplicationByOa(wcStaffLeaveApplication)
		}
		break
	case "adjustlevel":
		wcStaffAdjustlevelApplicationService.UpdateStaffAdjustlevelApplicationOaStatus(wcThirdOaPush.FdId, wcThirdOaPush.Status)
		wcStaffAdjustlevelApplication, err := wcStaffAdjustlevelApplicationService.GetWcStaffAdjustlevelApplicationByOA(wcThirdOaPush.FdId)
		if err == nil {
			wcStaffAdjustlevelApplicationService.UpdateWcStaffAdjustlevelApplicationByOA(wcStaffAdjustlevelApplication)
		}
		break
	case "pass":
		wcStaffPassApplicationService.UpdateWcStaffPassApplicationOaStatus(wcThirdOaPush.FdId, wcThirdOaPush.Status)
		wcStaffPassApplication, err := wcStaffPassApplicationService.GetWcStaffPassApplicationByOA(wcThirdOaPush.FdId)
		if err == nil {
			wcStaffPassApplicationService.UpdateWcStaffPassApplicationByOA(wcStaffPassApplication)
		}
		break
	default:
		response.FailWithMessage("未知的流程类型", c)
	}

	wcThirdOaPush.CreateTime = time.Now()
	if err := wcThirdOaPushService.CreateWcThirdOaPush(&wcThirdOaPush); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcThirdOaPush 删除OA回调日志
// @Tags WcThirdOaPush
// @Summary 删除OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcThirdOaPush true "删除OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcThirdOaPush/deleteWcThirdOaPush [delete]
func (wcThirdOaPushApi *WcThirdOaPushApi) DeleteWcThirdOaPush(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcThirdOaPushService.DeleteWcThirdOaPush(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcThirdOaPushByIds 批量删除OA回调日志
// @Tags WcThirdOaPush
// @Summary 批量删除OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcThirdOaPush/deleteWcThirdOaPushByIds [delete]
func (wcThirdOaPushApi *WcThirdOaPushApi) DeleteWcThirdOaPushByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcThirdOaPushService.DeleteWcThirdOaPushByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcThirdOaPush 更新OA回调日志
// @Tags WcThirdOaPush
// @Summary 更新OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body employee.WcThirdOaPush true "更新OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcThirdOaPush/updateWcThirdOaPush [put]
func (wcThirdOaPushApi *WcThirdOaPushApi) UpdateWcThirdOaPush(c *gin.Context) {
	var wcThirdOaPush employee.WcThirdOaPush
	err := c.ShouldBindJSON(&wcThirdOaPush)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcThirdOaPushService.UpdateWcThirdOaPush(wcThirdOaPush); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcThirdOaPush 用id查询OA回调日志
// @Tags WcThirdOaPush
// @Summary 用id查询OA回调日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employee.WcThirdOaPush true "用id查询OA回调日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcThirdOaPush/findWcThirdOaPush [get]
func (wcThirdOaPushApi *WcThirdOaPushApi) FindWcThirdOaPush(c *gin.Context) {
	ID := c.Query("ID")
	if rewcThirdOaPush, err := wcThirdOaPushService.GetWcThirdOaPush(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcThirdOaPush": rewcThirdOaPush}, c)
	}
}

// GetWcThirdOaPushList 分页获取OA回调日志列表
// @Tags WcThirdOaPush
// @Summary 分页获取OA回调日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcThirdOaPushSearch true "分页获取OA回调日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcThirdOaPush/getWcThirdOaPushList [get]
func (wcThirdOaPushApi *WcThirdOaPushApi) GetWcThirdOaPushList(c *gin.Context) {
	var pageInfo employeeReq.WcThirdOaPushSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcThirdOaPushService.GetWcThirdOaPushInfoList(pageInfo); err != nil {
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

// GetWcThirdOaPushPublic 不需要鉴权的OA回调日志接口
// @Tags WcThirdOaPush
// @Summary 不需要鉴权的OA回调日志接口
// @accept application/json
// @Produce application/json
// @Param data query employeeReq.WcThirdOaPushSearch true "分页获取OA回调日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcThirdOaPush/getWcThirdOaPushList [get]
func (wcThirdOaPushApi *WcThirdOaPushApi) GetWcThirdOaPushPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的OA回调日志接口信息",
	}, "获取成功", c)
}

// 通过模板ID获取流程类型ID
func GetFlowTypeId(templateId string) string {
	cfg, err := LoadConfig("./oa-config.ini")
	if err != nil {
		return ""
	}

	eaConfig := cfg.Section("template_list")

	// 获取模块对应配置ID
	return eaConfig.Key(templateId).MustString("")
}

// LoadConfig 加载配置文件
func LoadConfig(path string) (*ini.File, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		global.GVA_LOG.Error("LandrayOa LoadConfig 加载配置异常", zap.Error(err))
		return nil, err
	}
	return cfg, nil
}
