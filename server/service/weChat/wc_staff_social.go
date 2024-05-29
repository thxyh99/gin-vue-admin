package weChat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type WcStaffSocialService struct {
}

// CreateWcStaffSocial 创建社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) CreateWcStaffSocial(wcStaffSocial *weChat.WcStaffSocial) (err error) {
	err = global.GVA_DB.Create(wcStaffSocial).Error
	return err
}

// DeleteWcStaffSocial 删除社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) DeleteWcStaffSocial(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffSocial{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffSocialByIds 批量删除社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) DeleteWcStaffSocialByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffSocial{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffSocial 更新社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) UpdateWcStaffSocial(wcStaffSocial weChat.WcStaffSocial) (err error) {
	err = global.GVA_DB.Save(&wcStaffSocial).Error
	return err
}

// GetWcStaffSocial 根据ID获取社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) GetWcStaffSocial(ID string) (wcStaffSocial weChat2.WcStaffSocialResponse, err error) {
	var staffSocial weChat.WcStaffSocial
	err = global.GVA_DB.Where("id = ?", ID).First(&staffSocial).Error
	if err != nil {
		return
	}
	wcStaffSocial, err = wcStaffSocialService.AssembleStaffSocial(staffSocial)
	return
}

// GetWcStaffSocialInfoList 分页获取社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) GetWcStaffSocialInfoList(info weChatReq.WcStaffSocialSearch) (list []weChat2.WcStaffSocialResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffSocial{})
	var wcStaffSocials []weChat.WcStaffSocial
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.StaffId != nil {
		db = db.Where("staff_id = ?", info.StaffId)
	}
	if info.Keyword != "" {
		keyword := "%" + info.Keyword + "%"
		db = db.Where("(account LIKE ? OR credential_number LIKE ?)", keyword, keyword)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcStaffSocials).Error
	if err != nil {
		return
	}
	list, err = wcStaffSocialService.AssembleStaffSocialList(wcStaffSocials)
	return
}

// ImportExcel 导入Excel
func (wcStaffSocialService *WcStaffSocialService) ImportExcel(templateID, socialType string, file *multipart.FileHeader) (err error) {
	var template system.SysExportTemplate

	err = global.GVA_DB.First(&template, "template_id = ?", templateID).Error

	fmt.Println("-----------------------")
	fmt.Println(template)
	fmt.Println("-----------------------")

	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	var templateInfoMap = make(map[string]string)
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return err
	}

	var titleKeyMap = make(map[string]string)
	for key, title := range templateInfoMap {
		titleKeyMap[title] = key
	}

	db := global.GVA_DB
	if template.DBName != "" {
		db = global.MustGetGlobalDBByDBName(template.DBName)
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[3]
		values := rows[4:]

		//模版校验
		if socialType == "1" && len(excelTitle) != 20 {
			return errors.New("导入深圳社保Excel模版异常")
		} else if socialType == "2" && len(excelTitle) != 8 {
			return errors.New("导入深圳公积金Excel模版异常")
		} else if socialType == "3" && len(excelTitle) != 24 {
			return errors.New("导入东莞社保Excel模版异常")
		} else if socialType == "4" && len(excelTitle) != 12 {
			return errors.New("导入东莞公积金Excel模版异常")
		}

		//参数校验
		for _, row := range values {
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				err = wcStaffSocialService.checkImportParam(key, value)
				if err != nil {
					return err
				}
			}
		}

		for _, row := range values {
			var item = make(map[string]interface{})
			var name, credentialNumber, periodStart string
			var housingBase, housingRatioSelf, housingRatioUnit float64
			var staffExist weChat.WcStaff
			var staffSocial weChat.WcStaffSocial

			// 更新员工信息
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				if key == "credential_number" {
					credentialNumber = utils.FilterBreaksSpaces(value)
				}
				if key == "housing_base" {
					housingBase, err = strconv.ParseFloat(value, 64)
				}
				if key == "housing_ratio_self" {
					housingRatioSelf, err = strconv.ParseFloat(value, 64)
				}
				if key == "housing_ratio_unit" {
					housingRatioUnit, err = strconv.ParseFloat(value, 64)
				}
				if key == "period" {
					periods := strings.Split(value, "-")
					periodStart = utils.FilterBreaksSpaces(periods[0])
					item["period_start"] = periodStart
					item["period_end"] = utils.FilterBreaksSpaces(periods[1])
				} else {
					item[key] = utils.FilterBreaksSpaces(value)
				}
			}

			item["total_housing_self"] = housingBase * housingRatioSelf
			item["total_housing_unit"] = housingBase * housingRatioUnit
			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("id_number=?", credentialNumber).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New("员工不存在:" + name)
			} else {
				var staffSocialExist weChat.WcStaffSocial
				tx.Where("staff_id=? AND period_start=? AND type=?", staffExist.ID, periodStart, socialType).First(&staffSocialExist)
				item["staff_id"] = staffExist.ID
				item["type"] = socialType
				item["credential_type"] = 1
				if staffSocialExist.ID == 0 {
					cErr := tx.Table(staffSocial.TableName()).Create(&item).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffSocial.TableName()).Omit("name,account,credential_number,period_start,period_end,type,created_at").Where("id=?", staffSocialExist.ID).Updates(item).Error
					if cErr != nil {
						return cErr
					}
				}
			}
		}
		return nil
	})
}

// checkImportParam 参数校验
func (wcStaffSocialService *WcStaffSocialService) checkImportParam(key, value string) error {
	if key == "name" && value == "" {
		return errors.New("姓名不能为空")
	}

	if key == "credential_number" && value == "" {
		return errors.New("证件号码不能为空")
	}

	if key == "account" && value == "" {
		return errors.New("账号不能为空")
	}

	if key == "period" && value == "" {
		return errors.New("缴存时段不能为空")
	}

	if key == "housing_base" && value == "" {
		return errors.New("缴存基数不能为空")
	}

	if key == "housing_ratio_unit" && value == "" {
		return errors.New("单位缴存比例不能为空")
	}

	if key == "housing_ratio_self" && value == "" {
		return errors.New("个人缴存比例不能为空")
	}

	if key == "total_housing" && value == "" {
		return errors.New("金额合计不能为空")
	}

	return nil
}

func (wcStaffSocialService *WcStaffSocialService) AssembleStaffSocialList(staffSocials []weChat.WcStaffSocial) (newStaffSocials []weChat2.WcStaffSocialResponse, err error) {
	var newStaffSocial weChat2.WcStaffSocialResponse
	configInfo := config.GetConfigInfo()

	for _, staffSocial := range staffSocials {
		newStaffSocial.WcStaffSocial = staffSocial
		socialTypeText, _ := utils.Find(configInfo.SocialType, *staffSocial.Type)
		newStaffSocial.TypeText = socialTypeText
		credentialTypeText, _ := utils.Find(configInfo.CredentialType, *staffSocial.CredentialType)
		newStaffSocial.CredentialTypeText = credentialTypeText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffSocial.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffSocialList Err:", err)
			return
		}
		newStaffSocial.StaffName = staff.Name
		newStaffSocial.JobNum = staff.JobNum

		newStaffSocials = append(newStaffSocials, newStaffSocial)
	}
	return
}

func (wcStaffSocialService *WcStaffSocialService) AssembleStaffSocial(staffSocial weChat.WcStaffSocial) (newStaffSocial weChat2.WcStaffSocialResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffSocial.WcStaffSocial = staffSocial
	socialTypeText, _ := utils.Find(configInfo.SocialType, *staffSocial.Type)
	newStaffSocial.TypeText = socialTypeText
	credentialTypeText, _ := utils.Find(configInfo.CredentialType, *staffSocial.CredentialType)
	newStaffSocial.CredentialTypeText = credentialTypeText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffSocial.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffSocial Err:", err)
		return
	}

	newStaffSocial.StaffName = staff.Name
	newStaffSocial.JobNum = staff.JobNum

	return
}