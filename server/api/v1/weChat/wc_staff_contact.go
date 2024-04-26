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

type WcStaffContactApi struct {
}

var wcStaffContactService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffContactService


// CreateWcStaffContact 创建紧急联系人
// @Tags WcStaffContact
// @Summary 创建紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffContact true "创建紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaffContact/createWcStaffContact [post]
func (wcStaffContactApi *WcStaffContactApi) CreateWcStaffContact(c *gin.Context) {
	var wcStaffContact weChat.WcStaffContact
	err := c.ShouldBindJSON(&wcStaffContact)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffContactService.CreateWcStaffContact(&wcStaffContact); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaffContact 删除紧急联系人
// @Tags WcStaffContact
// @Summary 删除紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffContact true "删除紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaffContact/deleteWcStaffContact [delete]
func (wcStaffContactApi *WcStaffContactApi) DeleteWcStaffContact(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffContactService.DeleteWcStaffContact(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffContactByIds 批量删除紧急联系人
// @Tags WcStaffContact
// @Summary 批量删除紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaffContact/deleteWcStaffContactByIds [delete]
func (wcStaffContactApi *WcStaffContactApi) DeleteWcStaffContactByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffContactService.DeleteWcStaffContactByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaffContact 更新紧急联系人
// @Tags WcStaffContact
// @Summary 更新紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaffContact true "更新紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaffContact/updateWcStaffContact [put]
func (wcStaffContactApi *WcStaffContactApi) UpdateWcStaffContact(c *gin.Context) {
	var wcStaffContact weChat.WcStaffContact
	err := c.ShouldBindJSON(&wcStaffContact)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffContactService.UpdateWcStaffContact(wcStaffContact); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaffContact 用id查询紧急联系人
// @Tags WcStaffContact
// @Summary 用id查询紧急联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaffContact true "用id查询紧急联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaffContact/findWcStaffContact [get]
func (wcStaffContactApi *WcStaffContactApi) FindWcStaffContact(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaffContact, err := wcStaffContactService.GetWcStaffContact(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaffContact": rewcStaffContact}, c)
	}
}

// GetWcStaffContactList 分页获取紧急联系人列表
// @Tags WcStaffContact
// @Summary 分页获取紧急联系人列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffContactSearch true "分页获取紧急联系人列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffContact/getWcStaffContactList [get]
func (wcStaffContactApi *WcStaffContactApi) GetWcStaffContactList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffContactSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffContactService.GetWcStaffContactInfoList(pageInfo); err != nil {
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

// GetWcStaffContactPublic 不需要鉴权的紧急联系人接口
// @Tags WcStaffContact
// @Summary 不需要鉴权的紧急联系人接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffContactSearch true "分页获取紧急联系人列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaffContact/getWcStaffContactList [get]
func (wcStaffContactApi *WcStaffContactApi) GetWcStaffContactPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的紧急联系人接口信息",
    }, "获取成功", c)
}
