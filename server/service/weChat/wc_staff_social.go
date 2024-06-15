package weChat

import (
	"bytes"
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
	"net/url"
	"reflect"
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
	if info.Type != nil {
		db = db.Where("type = ?", *info.Type)
	}
	if info.StaffId != nil {
		db = db.Where("staff_id = ?", *info.StaffId)
	}
	if info.Keyword != "" {
		keyword := "%" + info.Keyword + "%"
		db = db.Where("(account LIKE ? OR credential_number LIKE ?)", keyword, keyword)
	}
	if info.Period != "" {
		db = db.Where("period_start = ?", info.Period)
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

	switch socialType {
	case "1":
		return wcStaffSocialService.importSzSbExcel(db, rows, titleKeyMap, socialType)
	case "2":
		return wcStaffSocialService.importSzGjjExcel(db, rows, titleKeyMap, socialType)
	case "3":
		return wcStaffSocialService.importDgSbExcel(db, rows, titleKeyMap, socialType)
	case "4":
		return wcStaffSocialService.importDgGjjExcel(db, rows, titleKeyMap, socialType)
	default:
		return errors.New("类型异常")
	}
}

// importSzSbExcel 导入深圳社保
func (wcStaffSocialService *WcStaffSocialService) importSzSbExcel(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, socialType string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		//台账年月：2024年04月
		periodStr := rows[2][0]
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("periodStr", periodStr)
		fmt.Println("year", periodStr[15:19])
		fmt.Println("month", periodStr[22:24])
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		period := periodStr[15:19] + periodStr[22:24]
		excelTitle := rows[4]
		values := rows[5:]

		if len(excelTitle) != 20 {
			return errors.New("导入深圳社保Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+6))
			}
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				fmt.Println("title-key-value", excelTitle[ii], key, value)
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

				//结合每个字段是否为空判断(最后一列为空的话这种方式判断不出来)
				if value == "" {
					return errors.New(fmt.Sprintf("%s不能为空", excelTitle[ii]))
				}
			}
		}

		for _, row := range values {
			var item = make(map[string]interface{})
			var name, credentialNumber string
			var staffExist weChat.WcStaff
			var staffSocial weChat.WcStaffSocial

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "id" {
					continue
				}
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				if key == "credential_number" {
					credentialNumber = utils.FilterBreaksSpaces(value)
				}
				item[key] = utils.FilterBreaksSpaces(value)
			}

			item["period_start"] = period
			item["period_end"] = period
			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("id_number=?", credentialNumber).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSocialExist weChat.WcStaffSocial
				tx.Where("staff_id=? AND period_start=? AND type=?", staffExist.ID, period, socialType).First(&staffSocialExist)
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

// importSzGjjExcel 导入深圳公积金
func (wcStaffSocialService *WcStaffSocialService) importSzGjjExcel(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, socialType string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[3]
		values := rows[4:]

		if len(excelTitle) != 8 {
			return errors.New("导入深圳公积金Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+5))
			}
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				fmt.Println("title-key-value", excelTitle[ii], key, value)
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

				//结合每个字段是否为空判断(最后一列为空的话这种方式判断不出来)
				if value == "" {
					return errors.New(fmt.Sprintf("%s不能为空", excelTitle[ii]))
				}
			}
		}

		for _, row := range values {
			var item = make(map[string]interface{})
			var name, credentialNumber, periodStart string
			var housingBase, housingRatioSelf, housingRatioUnit float64
			var staffExist weChat.WcStaff
			var staffSocial weChat.WcStaffSocial

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				if key == "credential_number" {
					credentialNumber = utils.FilterBreaksSpaces(value)
				}

				if key == "housing_base" {
					housingBase, _ = strconv.ParseFloat(value, 64)
				}
				if key == "housing_ratio_self" {
					housingRatioSelf, _ = strconv.ParseFloat(value, 64)
				}
				if key == "housing_ratio_unit" {
					housingRatioUnit, _ = strconv.ParseFloat(value, 64)
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
				return errors.New(fmt.Sprintf("员工%s不存在", name))
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

// importDgSbExcel 导入东莞社保
func (wcStaffSocialService *WcStaffSocialService) importDgSbExcel(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, socialType string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[6][:7]
		excelTitle = append(excelTitle, rows[7][7:21]...)
		excelTitle = append(excelTitle, rows[6][21:24]...)
		values := rows[8:]

		fmt.Println("excelTitle", len(excelTitle), excelTitle)

		if len(excelTitle) != 24 {
			return errors.New("导入东莞社保Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+6))
			}
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				fmt.Println("title-key-value", excelTitle[ii], key, value)
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

				//结合每个字段是否为空判断(最后一列为空的话这种方式判断不出来)
				if value == "" {
					return errors.New(fmt.Sprintf("%s不能为空", excelTitle[ii]))
				}
			}
		}

		for _, row := range values {
			var item = make(map[string]interface{})
			var name, credentialNumber, period string
			var staffExist weChat.WcStaff
			var staffSocial weChat.WcStaffSocial

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "id" || key == "pension_base1" || key == "unemployed_base1" || key == "medical_base1" {
					continue
				}
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				if key == "credential_number" {
					credentialNumber = utils.FilterBreaksSpaces(value)
				}

				if key == "credential_type" {
					if value == "港澳通行证" {
						item["credential_type"] = 2
					} else {
						item["credential_type"] = 1
					}
				} else if key == "period_start" || key == "period_end" {
					period = strings.ReplaceAll(value, "-", "")
					item[key] = utils.FilterBreaksSpaces(period)
				} else {
					item[key] = utils.FilterBreaksSpaces(value)
				}
			}

			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("id_number=?", credentialNumber).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSocialExist weChat.WcStaffSocial
				tx.Where("staff_id=? AND period_start=? AND type=?", staffExist.ID, period, socialType).First(&staffSocialExist)
				item["staff_id"] = staffExist.ID
				item["type"] = socialType
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

// importDgGjjExcel 导入东莞公积金
func (wcStaffSocialService *WcStaffSocialService) importDgGjjExcel(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, socialType string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[1]
		values := rows[2:]

		if len(excelTitle) != 12 {
			return errors.New("导入东莞公积金Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+3))
			}
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				fmt.Println("title-key-value", excelTitle[ii], key, value)
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

				//结合每个字段是否为空判断(最后一列为空的话这种方式判断不出来)
				if value == "" {
					return errors.New(fmt.Sprintf("%s不能为空", excelTitle[ii]))
				}
			}
		}

		for _, row := range values {
			var item = make(map[string]interface{})
			var name, credentialNumber, period string
			var staffExist weChat.WcStaff
			var staffSocial weChat.WcStaffSocial

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "id" {
					continue
				}
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				if key == "credential_number" {
					credentialNumber = utils.FilterBreaksSpaces(value)
				}
				if key == "housing_ratio_self" || key == "housing_ratio_unit" {
					housingRatio := strings.TrimSuffix(value, "%") // 去掉百分号
					percentage, err := strconv.Atoi(housingRatio)
					if err != nil {
						fmt.Println("Error:", err)
						return err
					}
					item[key] = float64(percentage) / 100
				} else if key == "period" {
					period = utils.FilterBreaksSpaces(value)
					period = strings.ReplaceAll(period, "-", "")
					item["period_start"] = period
					item["period_end"] = period
				} else if key == "credential_type" {
					if value == "港澳通行证" {
						item["credential_type"] = 2
					} else {
						item["credential_type"] = 1
					}
				} else {
					item[key] = utils.FilterBreaksSpaces(value)
				}
			}

			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("id_number=?", credentialNumber).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSocialExist weChat.WcStaffSocial
				tx.Where("staff_id=? AND period_start=? AND type=?", staffExist.ID, period, socialType).First(&staffSocialExist)
				item["staff_id"] = staffExist.ID
				item["type"] = socialType
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

// ExportExcel 导出Excel
func (wcStaffSocialService *WcStaffSocialService) ExportExcel(templateID string, values url.Values) (file *bytes.Buffer, name string, err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.Preload("Conditions").Preload("JoinTemplate").First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return nil, "", err
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var templateInfoMap = make(map[string]string)
	columns, err := utils.GetJSONKeys(template.TemplateInfo)
	if err != nil {
		return nil, "", err
	}
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return nil, "", err
	}
	socialType := values.Get("type")
	period := values.Get("period")
	staffId := values.Get("staffId")
	keyword := values.Get("keyword")

	fmt.Println("socialType-period-staffId-keyword", socialType, period, staffId, keyword)

	var tableTitle []string
	for _, key := range columns {
		//东莞社保表头字段转换
		if socialType == "3" {
			switch key {
			case "pension_base":
				templateInfoMap[key] = "企业基本养老保险(单位缴纳)缴费基数"
				break
			case "pension_unit":
				templateInfoMap[key] = "企业基本养老保险(单位缴纳)应缴金额"
				break
			case "pension_base1":
				templateInfoMap[key] = "企业基本养老保险(个人缴纳)缴费基数"
				break
			case "pension_self":
				templateInfoMap[key] = "企业基本养老保险(个人缴纳)应缴金额"
				break
			case "unemployed_base":
				templateInfoMap[key] = "失业保险(单位缴纳)缴费基数"
				break
			case "unemployed_unit":
				templateInfoMap[key] = "失业保险(单位缴纳)应缴金额"
				break
			case "unemployed_base1":
				templateInfoMap[key] = "失业保险(个人缴纳)缴费基数"
				break
			case "unemployed_self":
				templateInfoMap[key] = "失业保险(个人缴纳)应缴金额"
				break
			case "medical_base":
				templateInfoMap[key] = "单建统筹职工医保（含生育）（单位缴纳）缴费基数"
				break
			case "medical_unit":
				templateInfoMap[key] = "单建统筹职工医保（含生育）（单位缴纳）应缴金额"
				break
			case "medical_base1":
				templateInfoMap[key] = "单建统筹职工医保（含生育）（个人缴纳）缴费基数"
				break
			case "medical_self":
				templateInfoMap[key] = "单建统筹职工医保（含生育）（个人缴纳）应缴金额"
				break
			case "injury_insurance_base":
				templateInfoMap[key] = "工伤保险缴费基数"
				break
			case "injury_insurance_unit":
				templateInfoMap[key] = "工伤保险应缴金额"
				break
			}
		}
		tableTitle = append(tableTitle, templateInfoMap[key])
	}

	fmt.Println("tableTitle", tableTitle)
	var tableMap []map[string]interface{}
	db := global.GVA_DB
	var fields string

	switch socialType {
	case "1":
		name = "导出深圳社保"
		fields = `a.id, a.name, a.job_num,b.account,b.credential_number,b.total_social,b.total_social_self,b.total_social_unit,b.pension_base,b.pension_self,b.pension_unit,b.medical_base,b.medical_self,b.medical_unit,b.injury_insurance_base,b.injury_insurance_unit,b.unemployed_base,b.unemployed_self,b.unemployed_unit,b.birth_base,b.birth_unit `
		break
	case "2":
		name = "导出深圳公积金"
		fields = `a.id, a.name, a.job_num,b.credential_number,b.account,b.period_start as period,b.housing_base,b.housing_ratio_unit,b.housing_ratio_self,b.total_housing `
		break
	case "3":
		name = "导出东莞社保"
		fields = `a.id, a.name, a.job_num,b.credential_number,b.credential_type,b.account,b.period_start,b.period_end,b.pension_base,b.pension_unit,b.pension_base AS pension_base1,b.pension_self,b.unemployed_base,b.unemployed_unit,b.unemployed_base AS unemployed_base1,b.unemployed_self,b.medical_base,b.medical_unit,b.medical_base AS medical_base1,b.medical_self,b.injury_insurance_base,b.injury_insurance_unit,b.total_social_unit,b.total_social_self,b.total_social`
		break
	case "4":
		fields = `a.id, a.name, a.job_num,b.account,b.credential_type,b.credential_number,b.housing_base,b.housing_ratio_unit,b.housing_ratio_self,b.total_housing_unit,b.total_housing_self,b.total_housing,b.period_start AS period`
		name = "导出东莞公积金"
		break
	default:
		return nil, "", errors.New(fmt.Sprintf("社保公积金导出类型异常:%s", socialType))
	}
	where := fmt.Sprintf("b.type = %s", socialType)
	if keyword != "" {
		keyword = "%" + keyword + "%"
		where += fmt.Sprintf(" AND (b.account LIKE '%s' OR b.credential_number LIKE '%s' )", keyword, keyword)
	}
	if staffId != "" {
		where += fmt.Sprintf(" AND a.id = %s ", staffId)
	}
	if period != "" {
		where += fmt.Sprintf(" AND b.period_start = %s ", period)
	}
	sql := fmt.Sprintf(`SELECT %s FROM wc_staff AS a 
		LEFT JOIN wc_staff_social AS b ON a.id = b.staff_id AND b.deleted_at IS NULL
		WHERE %s `, fields, where)

	fmt.Println("sql", sql, keyword)

	db.Debug().Raw(sql).Scan(&tableMap)

	configInfo := config.GetConfigInfo()

	var rows [][]string
	rows = append(rows, tableTitle)
	for i, table := range tableMap {
		var row []string
		for _, column := range columns {
			if column == "id" {
				table[column] = i + 1
			}
			if column == "credential_type" {
				fmt.Println(reflect.TypeOf(table["credential_type"]))
				if intVal64, ok := table["credential_type"].(int64); ok {
					intVal := int(intVal64)
					text := configInfo.CredentialType[intVal]
					table[column] = text
				} else {
					table[column] = ""
				}
			}

			row = append(row, fmt.Sprintf("%v", table[column]))
		}
		rows = append(rows, row)
	}
	for i, row := range rows {
		for j, colCell := range row {
			err := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", utils.GetColumnName(j+1), i+1), colCell)
			if err != nil {
				return nil, "", err
			}
		}
	}
	f.SetActiveSheet(index)
	file, err = f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	return file, name, nil
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
