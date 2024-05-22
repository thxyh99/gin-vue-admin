package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type WcFileApi struct {
}

var wcFileService = service.ServiceGroupApp.WeChatServiceGroup.WcFileService

// CreateWcFile 创建文件管理
// @Tags WcFile
// @Summary 创建文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcFile true "创建文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcFile/createWcFile [post]
func (wcFileApi *WcFileApi) CreateWcFile(c *gin.Context) {
	var wcFile weChat.WcFile
	err := c.ShouldBindJSON(&wcFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcFileService.CreateWcFile(&wcFile); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcFile 删除文件管理
// @Tags WcFile
// @Summary 删除文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcFile true "删除文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcFile/deleteWcFile [delete]
func (wcFileApi *WcFileApi) DeleteWcFile(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcFileService.DeleteWcFile(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcFileByIds 批量删除文件管理
// @Tags WcFile
// @Summary 批量删除文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcFile/deleteWcFileByIds [delete]
func (wcFileApi *WcFileApi) DeleteWcFileByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcFileService.DeleteWcFileByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcFile 更新文件管理
// @Tags WcFile
// @Summary 更新文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcFile true "更新文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcFile/updateWcFile [put]
func (wcFileApi *WcFileApi) UpdateWcFile(c *gin.Context) {
	var wcFile weChat.WcFile
	err := c.ShouldBindJSON(&wcFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcFileService.UpdateWcFile(wcFile); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcFile 用id查询文件管理
// @Tags WcFile
// @Summary 用id查询文件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcFile true "用id查询文件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcFile/findWcFile [get]
func (wcFileApi *WcFileApi) FindWcFile(c *gin.Context) {
	ID := c.Query("ID")
	if rewcFile, err := wcFileService.GetWcFile(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcFile": rewcFile}, c)
	}
}

// GetWcFileList 分页获取文件管理列表
// @Tags WcFile
// @Summary 分页获取文件管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcFileSearch true "分页获取文件管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcFile/getWcFileList [get]
func (wcFileApi *WcFileApi) GetWcFileList(c *gin.Context) {
	var pageInfo weChatReq.WcFileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcFileService.GetWcFileInfoList(pageInfo); err != nil {
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

// GetWcFilePublic 不需要鉴权的文件管理接口
// @Tags WcFile
// @Summary 不需要鉴权的文件管理接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcFileSearch true "分页获取文件管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcFile/getWcFileList [get]
func (wcFileApi *WcFileApi) GetWcFilePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的文件管理接口信息",
	}, "获取成功", c)
}

// UploadFile 上传文件
func (wcFileApi *WcFileApi) UploadFile(c *gin.Context) {
	var file weChat2.WcFileResponse
	//noSave := c.DefaultQuery("noSave", "0")
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	configInfo := config.GetConfigInfo()
	if _, ok := utils.Find(configInfo.FileType, fileType); !ok {
		response.FailWithMessage("文件类型异常", c)
		return
	}
	staffId, _ := strconv.Atoi(c.PostForm("staffId"))
	if staffId == 0 {
		response.FailWithMessage("员工ID不能为空", c)
		return
	}
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = wcFileService.UploadFile(header, fileType, staffId) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}

	response.OkWithDetailed(file, "上传成功", c)
}

// EditFileName 编辑文件名或者备注
func (wcFileApi *WcFileApi) EditFileName(c *gin.Context) {
	var file weChat.WcFile
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wcFileService.EditFileName(file)
	if err != nil {
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// DeleteFile 删除文件
func (wcFileApi *WcFileApi) DeleteFile(c *gin.Context) {
	var file request.DeleteById
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wcFileService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetFileList 获取文件列表
func (wcFileApi *WcFileApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	configInfo := config.GetConfigInfo()
	if _, ok := utils.Find(configInfo.FileType, fileType); !ok {
		response.FailWithMessage("文件类型异常", c)
		return
	}
	staffId, _ := strconv.Atoi(c.PostForm("staffId"))
	if staffId == 0 {
		response.FailWithMessage("员工ID不能为空", c)
		return
	}
	list, total, err := wcFileService.GetFileRecordInfoList(pageInfo, fileType, staffId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
