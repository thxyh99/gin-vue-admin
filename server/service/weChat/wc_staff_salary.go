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
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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
