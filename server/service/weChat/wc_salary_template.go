package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcSalaryTemplateService struct {
}

// CreateWcSalaryTemplate 创建工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) CreateWcSalaryTemplate(wcSalaryTemplateRequest *weChatReq.WcSalaryTemplateRequest) (err error) {
	var wcSalaryTemplate weChat.WcSalaryTemplate
	wcSalaryTemplate = wcSalaryTemplateRequest.WcSalaryTemplate
	err = global.GVA_DB.Create(wcSalaryTemplate).Error
	if err != nil {
		return err
	}
	templateId := wcSalaryTemplate.ID

	//绑定模版对应的工资类型的工资单字段并且设置其可见性
	var wcSalaryField weChat.WcSalaryField
	items := make([]map[string]interface{}, 0, len(wcSalaryTemplateRequest.Fields))
	for _, field := range wcSalaryTemplateRequest.Fields {
		var item = make(map[string]interface{})
		item["template_id"] = templateId
		item["field"] = field.Field
		item["name"] = field.Name
		item["is_visible"] = field.IsVisible
		items = append(items, item)
	}

	err = global.GVA_DB.Table(wcSalaryField.TableName()).CreateInBatches(&items, 1000).Error
	return
}

// DeleteWcSalaryTemplate 删除工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) DeleteWcSalaryTemplate(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcSalaryTemplate{}, "id = ?", ID).Error
	return err
}

// DeleteWcSalaryTemplateByIds 批量删除工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) DeleteWcSalaryTemplateByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcSalaryTemplate{}, "id in ?", IDs).Error
	return err
}

// UpdateWcSalaryTemplate 更新工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) UpdateWcSalaryTemplate(wcSalaryTemplate weChat.WcSalaryTemplate) (err error) {
	err = global.GVA_DB.Save(&wcSalaryTemplate).Error
	return err
}

// GetWcSalaryTemplate 根据ID获取工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) GetWcSalaryTemplate(ID string) (wcSalaryTemplate weChat.WcSalaryTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcSalaryTemplate).Error
	return
}

// GetWcSalaryTemplateInfoList 分页获取工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) GetWcSalaryTemplateInfoList(info weChatReq.WcSalaryTemplateSearch) (list []weChat.WcSalaryTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcSalaryTemplate{})
	var wcSalaryTemplates []weChat.WcSalaryTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcSalaryTemplates).Error
	return wcSalaryTemplates, total, err
}
