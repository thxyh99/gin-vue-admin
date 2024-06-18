package weChat

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type WcStaffApi struct {
}

var wcStaffService = service.ServiceGroupApp.WeChatServiceGroup.WcStaffService

// CreateWcStaff 创建账号信息
// @Tags WcStaff
// @Summary 创建账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaff true "创建账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wcStaff/createWcStaff [post]
func (wcStaffApi *WcStaffApi) CreateWcStaff(c *gin.Context) {
	var wcStaff weChat.WcStaff
	err := c.ShouldBindJSON(&wcStaff)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffService.CreateWcStaff(&wcStaff); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWcStaff 删除账号信息
// @Tags WcStaff
// @Summary 删除账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaff true "删除账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wcStaff/deleteWcStaff [delete]
func (wcStaffApi *WcStaffApi) DeleteWcStaff(c *gin.Context) {
	ID := c.Query("ID")
	if err := wcStaffService.DeleteWcStaff(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWcStaffByIds 批量删除账号信息
// @Tags WcStaff
// @Summary 批量删除账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wcStaff/deleteWcStaffByIds [delete]
func (wcStaffApi *WcStaffApi) DeleteWcStaffByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wcStaffService.DeleteWcStaffByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWcStaff 更新账号信息
// @Tags WcStaff
// @Summary 更新账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weChat.WcStaff true "更新账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wcStaff/updateWcStaff [put]
func (wcStaffApi *WcStaffApi) UpdateWcStaff(c *gin.Context) {
	var wcStaff weChat.WcStaff
	err := c.ShouldBindJSON(&wcStaff)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wcStaffService.UpdateWcStaff(&wcStaff); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWcStaff 用id查询账号信息
// @Tags WcStaff
// @Summary 用id查询账号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaff true "用id查询账号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaff/findWcStaff [get]
func (wcStaffApi *WcStaffApi) FindWcStaff(c *gin.Context) {
	ID := c.Query("ID")
	if rewcStaff, err := wcStaffService.GetWcStaff(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewcStaff": rewcStaff}, c)
	}
}

// ObtainEmployeeRoster 获取员工花名册
// @Tags WcStaff
// @Summary 获取员工花名册
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChat.WcStaff true "获取员工花名册"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wcStaff/obtainEmployeeRoster [get]
func (wcStaffApi *WcStaffApi) ObtainEmployeeRoster(c *gin.Context) {
	ID := c.Query("id")
	staffId, _ := strconv.Atoi(ID)
	rewcStaff, err := wcStaffService.GetWcStaff(ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询员工信息失败!", zap.Error(err))
		response.FailWithMessage("查询员工信息失败", c)
	}
	rewcStaffJob, err := wcStaffJobService.GetWcStaffJob(ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询工作信息失败!", zap.Error(err))
		response.FailWithMessage("查询工作信息失败", c)
	}
	rewcStaffEducation, err := wcStaffEducationService.GetWcStaffEducation(ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询学历信息失败!", zap.Error(err))
		response.FailWithMessage("查询学历信息失败", c)
	}
	rewcStaffContact, err := wcStaffContactService.GetWcStaffContact(ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询紧急联系人失败!", zap.Error(err))
		response.FailWithMessage("查询紧急联系人失败", c)
	}
	rewcStaffBank, err := wcStaffBankService.GetWcStaffBank(ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询银行卡信息失败!", zap.Error(err))
		response.FailWithMessage("查询银行卡信息失败", c)
	}

	var pageInfo weChatReq.WcStaffAgreementSearch
	pageInfo.Page = 1
	pageInfo.PageSize = 10
	rewcStaffAgreementList, rewcStaffAgreementTotal, err := wcStaffAgreementService.GetWcStaffAgreementInfoList(pageInfo, staffId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询合同信息失败1!", zap.Error(err))
		response.FailWithMessage("查询合同信息失败", c)
	}
	rewcStaffAgreement := make([]weChat2.WcStaffAgreementResponse, 0, rewcStaffAgreementTotal)
	for _, item := range rewcStaffAgreementList {
		var attachment weChat2.WcFileResponse
		if *item.FileId != 0 {
			attachment, _ = wcFileService.FindFile(uint(*item.FileId))
		}
		item.Attachment = &attachment
		rewcStaffAgreement = append(rewcStaffAgreement, item)
	}

	rewcStaffMaterials := new(weChat2.WcStaffMaterialsResponse)
	fileList, err := wcFileService.GetFileRecordInfoListByStaffId(staffId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询证件材料失败!", zap.Error(err))
		response.FailWithMessage("查询证件材料失败", c)
	}

	for _, item := range fileList {
		clonedItem := item
		if *item.Type == 2 {
			rewcStaffMaterials.IdCardPortrait = append(rewcStaffMaterials.IdCardPortrait, &clonedItem)
		}
		if *item.Type == 3 {
			rewcStaffMaterials.IdCardNational = append(rewcStaffMaterials.IdCardNational, &clonedItem)
		}
		if *item.Type == 4 {
			rewcStaffMaterials.EducationCertificate = append(rewcStaffMaterials.EducationCertificate, &clonedItem)
		}
		if *item.Type == 5 {
			rewcStaffMaterials.DegreeCertificate = append(rewcStaffMaterials.DegreeCertificate, &clonedItem)
		}
		if *item.Type == 6 {
			rewcStaffMaterials.ResignationCertificate = append(rewcStaffMaterials.ResignationCertificate, &clonedItem)
		}
		if *item.Type == 7 {
			rewcStaffMaterials.OnboardingForm = append(rewcStaffMaterials.OnboardingForm, &clonedItem)
		}
		if *item.Type == 8 {
			rewcStaffMaterials.TrialProvide = append(rewcStaffMaterials.TrialProvide, &clonedItem)
		}
		if *item.Type == 9 {
			rewcStaffMaterials.PersonalResume = append(rewcStaffMaterials.PersonalResume, &clonedItem)
		}
		if *item.Type == 10 {
			rewcStaffMaterials.SkillCertificate = append(rewcStaffMaterials.SkillCertificate, &clonedItem)
		}
		if *item.Type == 11 {
			rewcStaffMaterials.Health = append(rewcStaffMaterials.Health, &clonedItem)
		}
	}

	response.OkWithData(gin.H{
		"rewcStaff":          rewcStaff,
		"rewcStaffJob":       rewcStaffJob,
		"rewcStaffEducation": rewcStaffEducation,
		"rewcStaffContact":   rewcStaffContact,
		"rewcStaffBank":      rewcStaffBank,
		"rewcStaffAgreement": rewcStaffAgreement,
		"rewcStaffMaterials": rewcStaffMaterials,
	}, c)
}

// GetWcStaffList 分页获取账号信息列表
// @Tags WcStaff
// @Summary 分页获取账号信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSearch true "分页获取账号信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaff/getWcStaffList [get]
func (wcStaffApi *WcStaffApi) GetWcStaffList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if statistics, list, total, err := wcStaffService.GetWcStaffInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResultCustom{
			Statistics: statistics,
			List:       list,
			Total:      total,
			Page:       pageInfo.Page,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetSimpleStaffList 分页获取账号信息列表
// @Tags WcStaff
// @Summary 分页获取账号信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSearch true "分页获取账号信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaff/getSimpleStaffList [get]
func (wcStaffApi *WcStaffApi) GetSimpleStaffList(c *gin.Context) {
	var pageInfo weChatReq.WcStaffSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wcStaffService.GetSimpleStaffInfoList(pageInfo); err != nil {
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

// GetWcStaffPublic 不需要鉴权的账号信息接口
// @Tags WcStaff
// @Summary 不需要鉴权的账号信息接口
// @accept application/json
// @Produce application/json
// @Param data query weChatReq.WcStaffSearch true "分页获取账号信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wcStaff/getWcStaffList [get]
func (wcStaffApi *WcStaffApi) GetWcStaffPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的账号信息接口信息",
	}, "获取成功", c)
}

// ImportExcel 导入表格
// @Tags SysImportTemplate
// @Summary 导入表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/importExcel [post]
func (wcStaffApi *WcStaffApi) ImportExcel(c *gin.Context) {
	templateID := c.Query("templateID")
	if templateID == "" {
		response.FailWithMessage("模板ID不能为空", c)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("文件获取失败!", zap.Error(err))
		response.FailWithMessage("文件获取失败", c)
		return
	}
	fmt.Println("=======ImportExcel 导入表格=========")
	if err := wcStaffService.ImportExcel(templateID, file); err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("导入成功", c)
	}
}

// ExportExcel 导出表格
// @Tags SysExportTemplate
// @Summary 导出表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /wcStaff/exportExcel [get]
func (wcStaffApi *WcStaffApi) ExportExcel(c *gin.Context) {
	templateID := c.Query("templateID")
	queryParams := c.Request.URL.Query()
	fmt.Println("===============================")
	fmt.Println("queryParams", queryParams)
	fmt.Println("===============================")
	if templateID == "" {
		response.FailWithMessage("模板ID不能为空", c)
		return
	}
	if file, name, err := wcStaffService.ExportExcel(templateID, queryParams); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+utils.RandomString(6)+".xlsx")) // 对下载的文件重命名
		c.Header("success", "true")
		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", file.Bytes())
	}
}
