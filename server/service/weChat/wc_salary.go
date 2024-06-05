package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcSalaryService struct {
}

// CreateWcSalary 创建工资单模版记录
func (wcSalaryService *WcSalaryService) CreateWcSalary(wcSalary *weChat.WcSalary) (err error) {
	err = global.GVA_DB.Create(wcSalary).Error
	return err
}

// DeleteWcSalary 删除工资单模版记录
func (wcSalaryService *WcSalaryService) DeleteWcSalary(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcSalary{}, "id = ?", ID).Error
	return err
}

// DeleteWcSalaryByIds 批量删除工资单模版记录
func (wcSalaryService *WcSalaryService) DeleteWcSalaryByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcSalary{}, "id in ?", IDs).Error
	return err
}

// UpdateWcSalary 更新工资单模版记录
func (wcSalaryService *WcSalaryService) UpdateWcSalary(wcSalary weChat.WcSalary) (err error) {
	err = global.GVA_DB.Save(&wcSalary).Error
	return err
}

// GetWcSalary 根据ID获取工资单模版记录
func (wcSalaryService *WcSalaryService) GetWcSalary(ID string) (wcSalary weChat.WcSalary, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcSalary).Error
	return
}

// GetWcSalaryInfoList 分页获取工资单模版记录
func (wcSalaryService *WcSalaryService) GetWcSalaryInfoList(info weChatReq.WcSalarySearch) (list []weChat.WcSalary, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcSalary{})
	var wcSalarys []weChat.WcSalary
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

	err = db.Find(&wcSalarys).Error
	return wcSalarys, total, err
}

//func (wcSalaryService *WcSalaryService) AssembleSalaryList(salaries []weChat.WcSalary) (newSalaries []weChat2.WcSalaryResponse, err error) {
//	var newSalary weChat2.WcSalaryResponse
//	configInfo := config.GetConfigInfo()
//	rankTypeList, err := GetRankTypeList()
//	if err != nil {
//		return
//	}
//	rankTypeMap := make(map[int]string)
//	for _, rankTypeItem := range rankTypeList {
//		rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
//	}
//
//	for _, salary := range salaries {
//		newSalary.WcSalary = salary
//		typeText, _ := utils.Find(configInfo.SalaryType, *salary.Type)
//		newSalary.TypeText = typeText
//		rankTypeText, _ := utils.GetValueByKey(rankTypeMap, *salary.RankType)
//		newSalary.RankTypeText = rankTypeText
//
//		var wcSalaryFields []weChat.WcSalaryField
//		err = global.GVA_DB.Where("template_id=?", salary.ID).Find(wcSalaryFields).Error
//		if err != nil {
//			return
//		}
//		newSalary.Fields = wcSalaryFields
//
//		newSalaries = append(newSalaries, newSalary)
//	}
//	return
//}
