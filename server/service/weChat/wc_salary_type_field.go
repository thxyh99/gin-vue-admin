package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type WcSalaryTypeFieldService struct {
}

// CreateWcSalaryTypeField 创建工资类型款项记录
func (wcSalaryTypeFieldService *WcSalaryTypeFieldService) CreateWcSalaryTypeField(wcSalaryTypeField *weChat.WcSalaryTypeField) (err error) {
	err = global.GVA_DB.Create(wcSalaryTypeField).Error
	return err
}

// DeleteWcSalaryTypeField 删除工资类型款项记录
func (wcSalaryTypeFieldService *WcSalaryTypeFieldService) DeleteWcSalaryTypeField(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcSalaryTypeField{}, "id = ?", ID).Error
	return err
}

// DeleteWcSalaryTypeFieldByIds 批量删除工资类型款项记录
func (wcSalaryTypeFieldService *WcSalaryTypeFieldService) DeleteWcSalaryTypeFieldByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcSalaryTypeField{}, "id in ?", IDs).Error
	return err
}

// UpdateWcSalaryTypeField 更新工资类型款项记录
func (wcSalaryTypeFieldService *WcSalaryTypeFieldService) UpdateWcSalaryTypeField(wcSalaryTypeField weChat.WcSalaryTypeField) (err error) {
	err = global.GVA_DB.Save(&wcSalaryTypeField).Error
	return err
}

// GetWcSalaryTypeField 根据ID获取工资类型款项记录
func (wcSalaryTypeFieldService *WcSalaryTypeFieldService) GetWcSalaryTypeField(ID string) (newWcSalaryTypeField weChat2.WcSalaryTypeFieldResponse, err error) {
	var wcSalaryTypeField weChat.WcSalaryTypeField
	err = global.GVA_DB.Where("id = ?", ID).First(&wcSalaryTypeField).Error
	if err != nil {
		return
	}
	newWcSalaryTypeField, err = wcSalaryTypeFieldService.AssembleSalaryTypeField(wcSalaryTypeField)
	return
}

// GetWcSalaryTypeFieldInfoList 分页获取工资类型款项记录
func (wcSalaryTypeFieldService *WcSalaryTypeFieldService) GetWcSalaryTypeFieldInfoList(info weChatReq.WcSalaryTypeFieldSearch) (list []weChat2.WcSalaryTypeFieldResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcSalaryTypeField{})
	var wcSalaryTypeFields []weChat.WcSalaryTypeField
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcSalaryTypeFields).Error
	if err != nil {
		return
	}
	list, err = wcSalaryTypeFieldService.AssembleSalaryTypeFieldList(wcSalaryTypeFields)
	return
}

func (wcSalaryTypeFieldService *WcSalaryTypeFieldService) AssembleSalaryTypeFieldList(salaryTypeFields []weChat.WcSalaryTypeField) (newSalaryTypeFields []weChat2.WcSalaryTypeFieldResponse, err error) {
	var newSalaryTypeField weChat2.WcSalaryTypeFieldResponse
	configInfo := config.GetConfigInfo()
	for _, salaryTypeField := range salaryTypeFields {
		newSalaryTypeField.WcSalaryTypeField = salaryTypeField
		typeText, _ := utils.Find(configInfo.SalaryType, *salaryTypeField.Type)
		newSalaryTypeField.TypeText = typeText
		isRequiredText, _ := utils.Find(configInfo.IsRequired, *salaryTypeField.IsRequired)
		newSalaryTypeField.IsRequiredText = isRequiredText
		newSalaryTypeFields = append(newSalaryTypeFields, newSalaryTypeField)
	}
	return
}

func (wcSalaryTypeFieldService *WcSalaryTypeFieldService) AssembleSalaryTypeField(salaryTypeField weChat.WcSalaryTypeField) (newSalaryTypeField weChat2.WcSalaryTypeFieldResponse, err error) {
	newSalaryTypeField.WcSalaryTypeField = salaryTypeField
	configInfo := config.GetConfigInfo()
	typeText, _ := utils.Find(configInfo.SalaryType, *salaryTypeField.Type)
	newSalaryTypeField.TypeText = typeText
	isRequiredText, _ := utils.Find(configInfo.IsRequired, *salaryTypeField.IsRequired)
	newSalaryTypeField.IsRequiredText = isRequiredText
	return
}
