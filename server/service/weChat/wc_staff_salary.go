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
	"time"
)

type WcStaffSalaryService struct {
}

// CreateWcStaffSalary 创建工资单记录
func (wcStaffSalaryService *WcStaffSalaryService) CreateWcStaffSalary(wcStaffSalary *weChat.WcStaffSalary) (err error) {
	err = global.GVA_DB.Create(wcStaffSalary).Error
	return err
}

// DeleteWcStaffSalary 删除工资单记录
func (wcStaffSalaryService *WcStaffSalaryService) DeleteWcStaffSalary(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffSalary{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffSalaryByIds 批量删除工资单记录
func (wcStaffSalaryService *WcStaffSalaryService) DeleteWcStaffSalaryByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffSalary{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffSalary 更新工资单记录
func (wcStaffSalaryService *WcStaffSalaryService) UpdateWcStaffSalary(wcStaffSalary weChat.WcStaffSalary) (err error) {
	err = global.GVA_DB.Save(&wcStaffSalary).Error
	return err
}

// GetWcStaffSalary 根据ID获取工资单记录
func (wcStaffSalaryService *WcStaffSalaryService) GetWcStaffSalary(ID string) (wcStaffSalaryResponse weChat2.WcStaffSalaryResponse, err error) {
	var wcStaffSalary weChat.WcStaffSalary
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffSalary).Error
	if err != nil {
		return
	}
	wcStaffSalaryResponse, err = wcStaffSalaryService.AssembleStaffSalary(wcStaffSalary)
	return
}

// GetWcStaffSalaryInfoList 分页获取工资单记录
func (wcStaffSalaryService *WcStaffSalaryService) GetWcStaffSalaryInfoList(info weChatReq.WcStaffSalarySearch) (list []weChat2.WcStaffSalaryResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffSalary{})
	var wcStaffSalaries []weChat.WcStaffSalary
	if info.Type != nil && *info.Type > 0 {
		db = db.Where("type = ?", *info.Type)
	}
	if info.Month != "" {
		db = db.Where("month = ?", info.Month)
	}
	if info.StaffId != nil && *info.StaffId > 0 {
		db = db.Where("staff_id = ?", *info.StaffId)
	}
	if info.Keyword != "" {
		keyword := "%" + info.Keyword + "%"
		db = db.Where("(department_first LIKE ? OR department_second LIKE ?)", keyword, keyword)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcStaffSalaries).Error
	if err != nil {
		return
	}
	list, err = wcStaffSalaryService.AssembleStaffSalaryList(wcStaffSalaries)
	return
}

func (wcStaffSalaryService *WcStaffSalaryService) AssembleStaffSalaryList(staffSalaries []weChat.WcStaffSalary) (newStaffSalaries []weChat2.WcStaffSalaryResponse, err error) {
	var newStaffSalary weChat2.WcStaffSalaryResponse
	configInfo := config.GetConfigInfo()

	for _, staffSalary := range staffSalaries {
		newStaffSalary.WcStaffSalary = staffSalary
		typeText, _ := utils.Find(configInfo.SalaryType, *staffSalary.Type)
		newStaffSalary.TypeText = typeText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffSalary.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffSalary Err:", err)
			return
		}
		newStaffSalary.StaffName = staff.Name
		newStaffSalary.JobNum = staff.JobNum

		newStaffSalaries = append(newStaffSalaries, newStaffSalary)
	}
	return
}

func (wcStaffSalaryService *WcStaffSalaryService) AssembleStaffSalary(staffSalary weChat.WcStaffSalary) (newStaffSalary weChat2.WcStaffSalaryResponse, err error) {
	newStaffSalary.WcStaffSalary = staffSalary
	configInfo := config.GetConfigInfo()
	typeText, _ := utils.Find(configInfo.SalaryType, *staffSalary.Type)
	newStaffSalary.TypeText = typeText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffSalary.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffSalary Err:", err)
		return
	}
	newStaffSalary.StaffName = staff.Name
	newStaffSalary.JobNum = staff.JobNum

	return
}

// ImportExcel 导入Excel
func (wcStaffSalaryService *WcStaffSalaryService) ImportExcel(templateID, salaryType, month, rankType string, file *multipart.FileHeader) (err error) {
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

	switch salaryType {
	case "1":
		return wcStaffSalaryService.importExcelA(db, rows, titleKeyMap, salaryType, month)
	case "2":
		return wcStaffSalaryService.importExcelB(db, rows, titleKeyMap, salaryType, month)
	case "3":
		return wcStaffSalaryService.importExcelB(db, rows, titleKeyMap, salaryType, month)
	case "4":
		return wcStaffSalaryService.importExcelB(db, rows, titleKeyMap, salaryType, month)
	case "5":
		return wcStaffSalaryService.importExcelC(db, rows, titleKeyMap, salaryType, month)
	case "6":
		return wcStaffSalaryService.importExcelD(db, rows, titleKeyMap, salaryType, month)
	case "7":
		return wcStaffSalaryService.importExcelE(db, rows, titleKeyMap, salaryType, month)
	case "8":
		return wcStaffSalaryService.importExcelF(db, rows, titleKeyMap, salaryType, month)
	default:
		return errors.New("工资类型异常")
	}
}

// importExcelA 导入基本工资
func (wcStaffSalaryService *WcStaffSalaryService) importExcelA(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, salaryType, month string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0][:13]
		excelTitle = append(excelTitle, rows[1][13:26]...)
		excelTitle = append(excelTitle, rows[2][26:39]...)
		excelTitle = append(excelTitle, rows[1][39:40]...)
		excelTitle = append(excelTitle, rows[2][40:43]...)
		excelTitle = append(excelTitle, rows[1][43:51]...)
		excelTitle = append(excelTitle, rows[0][51:56]...)
		values := rows[3:]

		if len(excelTitle) != 56 {
			return errors.New("导入Excel模版异常")
		}

		fmt.Println(excelTitle)

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+2))
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
			var name string
			var staffExist weChat.WcStaff
			var staffSalary weChat.WcStaffSalary

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "id" {
					continue
				}
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				item[key] = utils.FilterBreaksSpaces(value)
			}

			item["month"] = month
			item["type"] = salaryType
			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("name=?", name).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSalaryExist weChat.WcStaffSalary
				tx.Where("staff_id=? AND month=? AND type=?", staffExist.ID, month, salaryType).First(&staffSalaryExist)
				item["staff_id"] = staffExist.ID
				if staffSalaryExist.ID == 0 {
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					fmt.Println("item", item)
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					cErr := tx.Table(staffSalary.TableName()).Create(&item).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffSalary.TableName()).Omit("name,type,staff_id,created_at").Where("id=?", staffSalaryExist.ID).Updates(item).Error
					if cErr != nil {
						return cErr
					}
				}
			}
		}
		return nil
	})
}

// importExcelB 导入集团经营绩效奖励|节日金|半年奖
func (wcStaffSalaryService *WcStaffSalaryService) importExcelB(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, salaryType, month string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]

		if len(excelTitle) != 6 {
			return errors.New("导入Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+2))
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
			var name string
			var staffExist weChat.WcStaff
			var staffSalary weChat.WcStaffSalary

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				item[key] = utils.FilterBreaksSpaces(value)
			}

			item["month"] = month
			item["type"] = salaryType
			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("name=?", name).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSalaryExist weChat.WcStaffSalary
				tx.Where("staff_id=? AND month=? AND type=?", staffExist.ID, month, salaryType).First(&staffSalaryExist)
				item["staff_id"] = staffExist.ID
				if staffSalaryExist.ID == 0 {
					cErr := tx.Table(staffSalary.TableName()).Create(&item).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffSalary.TableName()).Omit("name,type,staff_id,created_at").Where("id=?", staffSalaryExist.ID).Updates(item).Error
					if cErr != nil {
						return cErr
					}
				}
			}
		}
		return nil
	})
}

// importExcelC 导入年度奖金
func (wcStaffSalaryService *WcStaffSalaryService) importExcelC(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, salaryType, month string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]

		if len(excelTitle) != 9 {
			return errors.New("导入Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+2))
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
			var name string
			var staffExist weChat.WcStaff
			var staffSalary weChat.WcStaffSalary

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				item[key] = utils.FilterBreaksSpaces(value)
			}

			item["month"] = month
			item["type"] = salaryType
			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("name=?", name).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSalaryExist weChat.WcStaffSalary
				tx.Where("staff_id=? AND month=? AND type=?", staffExist.ID, month, salaryType).First(&staffSalaryExist)
				item["staff_id"] = staffExist.ID
				if staffSalaryExist.ID == 0 {
					cErr := tx.Table(staffSalary.TableName()).Create(&item).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffSalary.TableName()).Omit("name,type,staff_id,created_at").Where("id=?", staffSalaryExist.ID).Updates(item).Error
					if cErr != nil {
						return cErr
					}
				}
			}
		}
		return nil
	})
}

// importExcelD 导入总部职能体系月度奖金
func (wcStaffSalaryService *WcStaffSalaryService) importExcelD(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, salaryType, month string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]

		if len(excelTitle) != 16 {
			return errors.New("导入Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+2))
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
			var name string
			var staffExist weChat.WcStaff
			var staffSalary weChat.WcStaffSalary

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "id" {
					continue
				}
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				item[key] = utils.FilterBreaksSpaces(value)
			}

			item["month"] = month
			item["type"] = salaryType
			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("name=?", name).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSalaryExist weChat.WcStaffSalary
				tx.Where("staff_id=? AND month=? AND type=?", staffExist.ID, month, salaryType).First(&staffSalaryExist)
				item["staff_id"] = staffExist.ID
				if staffSalaryExist.ID == 0 {
					cErr := tx.Table(staffSalary.TableName()).Create(&item).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffSalary.TableName()).Omit("name,type,staff_id,created_at").Where("id=?", staffSalaryExist.ID).Updates(item).Error
					if cErr != nil {
						return cErr
					}
				}
			}
		}
		return nil
	})
}

// importExcelE 导入总部金纳斯市场体系月度奖金
func (wcStaffSalaryService *WcStaffSalaryService) importExcelE(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, salaryType, month string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0][:7]
		excelTitle = append(excelTitle, rows[1][7:11]...)
		excelTitle = append(excelTitle, rows[0][11:22]...)
		values := rows[2:]

		if len(excelTitle) != 22 {
			return errors.New("导入Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+2))
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
			var name string
			var staffExist weChat.WcStaff
			var staffSalary weChat.WcStaffSalary

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "id" {
					continue
				}
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				item[key] = utils.FilterBreaksSpaces(value)
			}

			item["month"] = month
			item["type"] = salaryType
			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("name=?", name).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSalaryExist weChat.WcStaffSalary
				tx.Where("staff_id=? AND month=? AND type=?", staffExist.ID, month, salaryType).First(&staffSalaryExist)
				item["staff_id"] = staffExist.ID
				if staffSalaryExist.ID == 0 {
					cErr := tx.Table(staffSalary.TableName()).Create(&item).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffSalary.TableName()).Omit("name,type,staff_id,created_at").Where("id=?", staffSalaryExist.ID).Updates(item).Error
					if cErr != nil {
						return cErr
					}
				}
			}
		}
		return nil
	})
}

// importExcelF 导入总部调理中心体系月度奖金
func (wcStaffSalaryService *WcStaffSalaryService) importExcelF(db *gorm.DB, rows [][]string, titleKeyMap map[string]string, salaryType, month string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]

		if len(excelTitle) != 8 {
			return errors.New("导入Excel模版异常")
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+2))
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
			var name string
			var staffExist weChat.WcStaff
			var staffSalary weChat.WcStaffSalary

			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "name" {
					name = utils.FilterBreaksSpaces(value)
				}
				item[key] = utils.FilterBreaksSpaces(value)
			}

			item["month"] = month
			item["type"] = salaryType
			item["created_at"] = now
			item["updated_at"] = now

			tx.Where("name=?", name).First(&staffExist)
			if staffExist.ID == 0 {
				return errors.New(fmt.Sprintf("员工%s不存在", name))
			} else {
				var staffSalaryExist weChat.WcStaffSalary
				tx.Where("staff_id=? AND month=? AND type=?", staffExist.ID, month, salaryType).First(&staffSalaryExist)
				item["staff_id"] = staffExist.ID
				if staffSalaryExist.ID == 0 {
					cErr := tx.Table(staffSalary.TableName()).Create(&item).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffSalary.TableName()).Omit("name,type,staff_id,created_at").Where("id=?", staffSalaryExist.ID).Updates(item).Error
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
func (wcStaffSalaryService *WcStaffSalaryService) ExportExcel(templateID string, values url.Values) (file *bytes.Buffer, name string, err error) {
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
	salaryType := values.Get("type")
	staffId := values.Get("staffId")
	keyword := values.Get("keyword")
	month := values.Get("month")
	var tableTitle []string
	for _, key := range columns {
		tableTitle = append(tableTitle, templateInfoMap[key])
	}

	fmt.Println("tableTitle", tableTitle)
	var tableMap []map[string]interface{}
	db := global.GVA_DB

	switch salaryType {
	case "1":
		name = "导出工资表"
		break
	case "2":
		name = "导出集团经营绩效奖励"
		break
	case "3":
		name = "导出节日金"
		break
	case "4":
		name = "导出半年奖"
		break
	case "5":
		name = "导出年度奖金"
		break
	case "6":
		name = "导出总部职能体系月度奖金"
		break
	case "7":
		name = "导出总部金纳斯市场体系月度奖金"
		break
	case "8":
		name = "导出总部调理中心体系月度奖金"
		break
	default:
		return nil, "", errors.New(fmt.Sprintf("工资单导出类型异常:%s", salaryType))
	}
	where := fmt.Sprintf("b.type = %s", salaryType)
	if keyword != "" {
		keyword = "%" + keyword + "%"
		where += fmt.Sprintf(" AND (b.department_first LIKE '%s' OR b.department_second LIKE '%s' )", keyword, keyword)
	}
	if staffId != "" {
		where += fmt.Sprintf(" AND a.id = %s ", staffId)
	}
	if month != "" {
		where += fmt.Sprintf(" AND b.month = %s ", month)
	}
	sql := fmt.Sprintf(`SELECT a.name, a.job_num, b.* FROM wc_staff AS a 
		LEFT JOIN wc_staff_salary AS b ON a.id = b.staff_id AND b.deleted_at IS NULL
		WHERE %s `, where)

	fmt.Println("sql", sql)

	db.Debug().Raw(sql).Scan(&tableMap)

	var rows [][]string
	rows = append(rows, tableTitle)
	for i, table := range tableMap {
		var row []string
		for _, column := range columns {
			if column == "id" {
				table[column] = i + 1
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
