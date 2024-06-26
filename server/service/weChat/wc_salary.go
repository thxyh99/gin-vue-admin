package weChat

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
	"time"
)

type WcSalaryService struct {
}

// CreateWcSalary 创建工资单发放记录
func (wcSalaryService *WcSalaryService) CreateWcSalary(wcSalary *weChat.WcSalary) (err error) {
	var salary weChat.WcSalary
	err = global.GVA_DB.Where("month = ? AND template_id = ?", wcSalary.Month, wcSalary.TemplateId).First(&salary).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if salary.ID > 0 {
		return errors.New("该工资发放已经创建，请勿重复操作！")
	}

	wcSalary.CreatedAt = time.Now()
	wcSalary.UpdatedAt = time.Now()
	err = global.GVA_DB.Create(wcSalary).Error
	return err
}

// DeleteWcSalary 删除工资单发放记录
func (wcSalaryService *WcSalaryService) DeleteWcSalary(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcSalary{}, "id = ?", ID).Error
	return err
}

// DeleteWcSalaryByIds 批量删除工资单发放记录
func (wcSalaryService *WcSalaryService) DeleteWcSalaryByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcSalary{}, "id in ?", IDs).Error
	return err
}

// UpdateWcSalary 更新工资单发放记录
func (wcSalaryService *WcSalaryService) UpdateWcSalary(wcSalary weChat.WcSalary) (err error) {
	err = global.GVA_DB.Save(&wcSalary).Error
	return err
}

// GetWcSalary 根据ID获取工资单发放记录
func (wcSalaryService *WcSalaryService) GetWcSalary(ID string) (wcSalary weChat2.WcSalaryResponse, err error) {
	var salary weChat.WcSalary
	err = global.GVA_DB.Where("id = ?", ID).First(&salary).Error
	if err != nil {
		return
	}
	wcSalary, err = wcSalaryService.AssembleSalary(salary)
	return
}

// GetWcSalaryInfoList 分页获取工资单发放记录
func (wcSalaryService *WcSalaryService) GetWcSalaryInfoList(info weChatReq.WcSalarySearch) (list []weChat2.WcSalaryResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcSalary{})
	var wcSalaries []weChat.WcSalary
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MonthStart != "" && info.MonthEnd != "" {
		db = db.Where("month BETWEEN ? AND ?", info.MonthStart, info.MonthEnd)
	}
	if info.TemplateId != nil && *info.TemplateId > 0 {
		db = db.Where("template_id = ?", *info.TemplateId)
	}
	err = db.Count(&total).Error
	if err != nil {
		fmt.Println("GetWcSalaryInfoList Err", err.Error())
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("month DESC,updated_at DESC").Find(&wcSalaries).Error
	if err != nil {
		return
	}

	fmt.Println("wcSalaries", wcSalaries)

	list, err = wcSalaryService.AssembleSalaryList(wcSalaries)
	return
}

func (wcSalaryService *WcSalaryService) AssembleSalary(salary weChat.WcSalary) (newSalary weChat2.WcSalaryResponse, err error) {
	configInfo := config.GetConfigInfo()
	rankTypeList, err := GetRankTypeList()
	if err != nil {
		return
	}
	rankTypeMap := make(map[int]string)
	for _, rankTypeItem := range rankTypeList {
		rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
	}
	newSalary.WcSalary = salary
	var wcSalaryTemplate weChat.WcSalaryTemplate
	err = global.GVA_DB.Where("id=?", salary.TemplateId).Find(&wcSalaryTemplate).Error
	if err != nil {
		return
	}
	newSalary.Name = wcSalaryTemplate.Name
	typeText, _ := utils.Find(configInfo.SalaryType, *wcSalaryTemplate.Type)
	newSalary.TypeText = typeText
	newSalary.Type = wcSalaryTemplate.Type
	rankTypeText, _ := utils.GetValueByKey(rankTypeMap, *wcSalaryTemplate.RankType)
	newSalary.RankTypeText = rankTypeText
	newSalary.RankType = wcSalaryTemplate.RankType
	switch *wcSalaryTemplate.Type {
	case 1:
		newSalary.TemplateId = "staffSalary1"
		break
	case 2:
		newSalary.TemplateId = "staffSalary234"
		break
	case 3:
		newSalary.TemplateId = "staffSalary234"
		break
	case 4:
		newSalary.TemplateId = "staffSalary234"
		break
	case 5:
		newSalary.TemplateId = "staffSalary5"
		break
	case 6:
		newSalary.TemplateId = "staffSalary6"
		break
	case 7:
		newSalary.TemplateId = "staffSalary7"
		break
	case 8:
		newSalary.TemplateId = "staffSalary8"
		break
	default:
		return newSalary, errors.New(fmt.Sprintf("工资类型%d异常", *wcSalaryTemplate.Type))
	}

	return
}

func (wcSalaryService *WcSalaryService) AssembleSalaryList(salaries []weChat.WcSalary) (newSalaries []weChat2.WcSalaryResponse, err error) {
	configInfo := config.GetConfigInfo()
	rankTypeList, err := GetRankTypeList()
	if err != nil {
		fmt.Println("AssembleSalaryList Err1", err.Error())
		return
	}
	rankTypeMap := make(map[int]string)
	for _, rankTypeItem := range rankTypeList {
		rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
	}
	for _, salary := range salaries {
		var newSalary weChat2.WcSalaryResponse
		newSalary.WcSalary = salary
		var wcSalaryTemplate weChat.WcSalaryTemplate
		err = global.GVA_DB.Where("id=?", salary.TemplateId).Find(&wcSalaryTemplate).Error
		if err != nil {
			fmt.Println("AssembleSalaryList Err2", err.Error())
			return
		}
		newSalary.Name = wcSalaryTemplate.Name
		typeText, _ := utils.Find(configInfo.SalaryType, *wcSalaryTemplate.Type)
		newSalary.TypeText = typeText
		newSalary.Type = wcSalaryTemplate.Type
		rankTypeText, _ := utils.GetValueByKey(rankTypeMap, *wcSalaryTemplate.RankType)
		newSalary.RankTypeText = rankTypeText
		newSalary.RankType = wcSalaryTemplate.RankType
		switch *wcSalaryTemplate.Type {
		case 1:
			newSalary.TemplateId = "staffSalary1"
			break
		case 2:
			newSalary.TemplateId = "staffSalary234"
			break
		case 3:
			newSalary.TemplateId = "staffSalary234"
			break
		case 4:
			newSalary.TemplateId = "staffSalary234"
			break
		case 5:
			newSalary.TemplateId = "staffSalary5"
			break
		case 6:
			newSalary.TemplateId = "staffSalary6"
			break
		case 7:
			newSalary.TemplateId = "staffSalary7"
			break
		case 8:
			newSalary.TemplateId = "staffSalary8"
			break
		default:
			return newSalaries, errors.New(fmt.Sprintf("工资类型%d异常", *wcSalaryTemplate.Type))
		}
		newSalaries = append(newSalaries, newSalary)
	}
	return newSalaries, err
}
