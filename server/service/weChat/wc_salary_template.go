package weChat

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type WcSalaryTemplateService struct {
}

// CreateWcSalaryTemplate 创建工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) CreateWcSalaryTemplate(wcSalaryTemplateRequest *weChatReq.WcSalaryTemplateCreateRequest) (err error) {

	var wcSalaryTemplateExist weChat.WcSalaryTemplate
	err = global.GVA_DB.Where("type = ? AND rank_type = ?", wcSalaryTemplateRequest.Type, wcSalaryTemplateRequest.RankType).First(&wcSalaryTemplateExist).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if wcSalaryTemplateExist.ID > 0 {
		return errors.New("该类型模版已经存在，请勿重复创建！")
	}

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
func (wcSalaryTemplateService *WcSalaryTemplateService) UpdateWcSalaryTemplate(wcSalaryTemplateRequest weChatReq.WcSalaryTemplateUpdateRequest) (err error) {
	var wcSalaryTemplate weChat.WcSalaryTemplate
	wcSalaryTemplate = wcSalaryTemplateRequest.WcSalaryTemplate
	err = global.GVA_DB.Save(&wcSalaryTemplate).Error
	if err != nil {
		return err
	}
	for _, field := range wcSalaryTemplateRequest.Fields {
		var wcSalaryTemplateItem weChat.WcSalaryTemplate
		err = global.GVA_DB.Where("id = ?", field.ID).First(&wcSalaryTemplateItem).Update("is_visible", field.IsVisible).Error
		if err != nil {
			return err
		}
	}
	return
}

// GetWcSalaryTemplate 根据ID获取工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) GetWcSalaryTemplate(ID string) (wcSalaryTemplateResponse weChat2.WcSalaryTemplateResponse, err error) {
	var wcSalaryTemplate weChat.WcSalaryTemplate
	err = global.GVA_DB.Where("id = ?", ID).First(&wcSalaryTemplate).Error
	if err != nil {
		return
	}

	wcSalaryTemplateResponse, err = wcSalaryTemplateService.AssembleSalaryTemplate(wcSalaryTemplate)
	return
}

// GetWcSalaryTemplateInfoList 分页获取工资单模版记录
func (wcSalaryTemplateService *WcSalaryTemplateService) GetWcSalaryTemplateInfoList(info weChatReq.WcSalaryTemplateSearch) (list []weChat2.WcSalaryTemplateResponse, total int64, err error) {
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
	if err != nil {
		return
	}

	list, err = wcSalaryTemplateService.AssembleSalaryTemplateList(wcSalaryTemplates)
	return
}

func (wcSalaryTemplateService *WcSalaryTemplateService) AssembleSalaryTemplateList(salaryTemplates []weChat.WcSalaryTemplate) (newSalaryTemplates []weChat2.WcSalaryTemplateResponse, err error) {
	var newSalaryTemplate weChat2.WcSalaryTemplateResponse
	configInfo := config.GetConfigInfo()
	rankTypeList, err := GetRankTypeList()
	if err != nil {
		return
	}
	rankTypeMap := make(map[int]string)
	for _, rankTypeItem := range rankTypeList {
		rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
	}

	for _, salaryTemplate := range salaryTemplates {
		newSalaryTemplate.WcSalaryTemplate = salaryTemplate
		typeText, _ := utils.Find(configInfo.SalaryType, *salaryTemplate.Type)
		newSalaryTemplate.TypeText = typeText
		rankTypeText, _ := utils.GetValueByKey(rankTypeMap, *salaryTemplate.RankType)
		newSalaryTemplate.RankTypeText = rankTypeText

		var wcSalaryFields []weChat.WcSalaryField
		err = global.GVA_DB.Where("template_id=?", salaryTemplate.ID).Find(wcSalaryFields).Error
		if err != nil {
			return
		}
		newSalaryTemplate.Fields = wcSalaryFields

		newSalaryTemplates = append(newSalaryTemplates, newSalaryTemplate)
	}
	return
}

func (wcSalaryTemplateService *WcSalaryTemplateService) AssembleSalaryTemplate(salaryTemplate weChat.WcSalaryTemplate) (newSalaryTemplate weChat2.WcSalaryTemplateResponse, err error) {
	newSalaryTemplate.WcSalaryTemplate = salaryTemplate
	configInfo := config.GetConfigInfo()
	typeText, _ := utils.Find(configInfo.SalaryType, *salaryTemplate.Type)
	newSalaryTemplate.TypeText = typeText
	rankTypeList, err := GetRankTypeList()
	if err != nil {
		return
	}
	rankTypeMap := make(map[int]string)
	for _, rankTypeItem := range rankTypeList {
		rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
	}
	rankTypeText, _ := utils.GetValueByKey(rankTypeMap, *salaryTemplate.RankType)
	newSalaryTemplate.RankTypeText = rankTypeText

	var wcSalaryFields []weChat.WcSalaryField
	err = global.GVA_DB.Where("template_id=?", salaryTemplate.ID).Find(wcSalaryFields).Error
	if err != nil {
		return
	}
	newSalaryTemplate.Fields = wcSalaryFields

	return
}
