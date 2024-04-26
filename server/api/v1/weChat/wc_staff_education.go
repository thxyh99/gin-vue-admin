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

type WcStaffEducationApi struct {
}

var wcStaffEducationService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffEducationService


// CreateWcStaffEducation 创建学历信息
// @Tags WcStaffEducation
// @Summary 创建学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffEducation true "创建学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffEducation/createWcStaffEducation [post]
func (wcStaffEducationApi *WcStaffEducationApi) CreateWcStaffEducation(c *gin.Context) {
	var wcStaffEducation weChat.WcStaffEducation
	err := c.ShouldBindJSON(&wcStaffEducation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffEducationService.CreateWcStaffEducation(&wcStaffEducation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffEducation 删除学历信息
// @Tags WcStaffEducation
// @Summary 删除学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffEducation true "删除学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffEducation/deleteWcStaffEducation [delete]
func (wcStaffEducationApi *WcStaffEducationApi) DeleteWcStaffEducation(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffEducationService.DeleteWcStaffEducation(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffEducationByIds 批量删除学历信息
// @Tags WcStaffEducation
// @Summary 批量删除学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffEducation/deleteWcStaffEducationByIds [delete]
func (wcStaffEducationApi *WcStaffEducationApi) DeleteWcStaffEducationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffEducationService.DeleteWcStaffEducationByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffEducation 更新学历信息
// @Tags WcStaffEducation
// @Summary 更新学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffEducation true "更新学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffEducation/updateWcStaffEducation [put]
func (wcStaffEducationApi *WcStaffEducationApi) UpdateWcStaffEducation(c *gin.Context) {
	var wcStaffEducation weChat.WcStaffEducation
	err := c.ShouldBindJSON(&wcStaffEducation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffEducationService.UpdateWcStaffEducation(wcStaffEducation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffEducation 用id查询学历信息
// @Tags WcStaffEducation
// @Summary 用id查询学历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffEducation true "用id查询学历信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffEducation/findWcStaffEducation [get]
func (wcStaffEducationApi *WcStaffEducationApi) FindWcStaffEducation(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffEducation, err := wcStaffEducationService.GetWcStaffEducation(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffEducation": rewcStaffEducation}, c)
	}
}

// GetWcStaffEducationList 分页获取学历信息列表
// @Tags WcStaffEducation
// @Summary 分页获取学历信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffEducationSearch true "分页获取学历信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffEducation/getWcStaffEducationList [get]
func (wcStaffEducationApi *WcStaffEducationApi) GetWcStaffEducationList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffEducationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffEducationService.GetWcStaffEducationInfoList(pageInfo); err != nil {
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

// GetWcStaffEducationPublic 不需要鉴权的学历信息接口
// @Tags WcStaffEducation
// @Summary 不需要鉴权的学历信息接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffEducationSearch true "分页获取学历信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffEducation/getWcStaffEducationList [get]
func (wcStaffEducationApi *WcStaffEducationApi) GetWcStaffEducationPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的学历信息接口信息",
    }, "获取成功", c)
}
