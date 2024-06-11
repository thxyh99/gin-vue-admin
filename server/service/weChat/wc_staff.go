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

type WcStaffService struct {
}

// CreateWcStaff 创建账号信息记录
func (wcStaffService *WcStaffService) CreateWcStaff(wcStaff *weChat.WcStaff) (err error) {
	err = global.GVA_DB.Create(&wcStaff).Error
	if err != nil {
		fmt.Println("err1:", err)
		return err
	}

	return
}

// DeleteWcStaff 删除账号信息记录
func (wcStaffService *WcStaffService) DeleteWcStaff(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaff{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffByIds 批量删除账号信息记录
func (wcStaffService *WcStaffService) DeleteWcStaffByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaff{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaff 更新账号信息记录
func (wcStaffService *WcStaffService) UpdateWcStaff(wcStaff *weChat.WcStaff) (err error) {
	err = global.GVA_DB.Save(&wcStaff).Error
	return err
}

// GetWcStaff 根据ID获取账号信息记录
func (wcStaffService *WcStaffService) GetWcStaff(ID string) (wcStaffResponse weChat2.WcStaffResponse, err error) {
	var wcStaff weChat.WcStaff
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaff).Error
	if err != nil {
		return
	}
	wcStaffResponse, err = wcStaffService.AssembleStaff(wcStaff)

	return
}

// GetWcStaffInfoList 分页获取账号信息记录
func (wcStaffService *WcStaffService) GetWcStaffInfoList(info weChatReq.WcStaffSearch) (wcStaffResponse []weChat2.WcStaffResponse, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaff{})
	var wcStaffs []weChat.WcStaff
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 添加员工名称、员工工号、手机模糊查询
	if info.Keyword != "" {
		keyword := "%" + info.Keyword + "%"
		db = db.Where("(name LIKE ? OR job_num LIKE ? OR mobile LIKE ?)", keyword, keyword, keyword)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcStaffs).Error

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("GetWcStaffInfoList err", err)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	wcStaffResponse, err = wcStaffService.AssembleStaffList(wcStaffs)

	return
}

// GetSimpleStaffInfoList 分页获取账号信息记录
func (wcStaffService *WcStaffService) GetSimpleStaffInfoList(info weChatReq.WcStaffSearch) (list []weChat.WcStaff, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaff{})
	var wcStaffs []weChat.WcStaff
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

	err = db.Find(&wcStaffs).Error

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("GetSimpleStaffInfoList err", err)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	return wcStaffs, total, err
}

// ImportExcel 导入Excel
func (wcStaffService *WcStaffService) ImportExcel(templateID string, file *multipart.FileHeader) (err error) {
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
		excelTitle := rows[0]
		values := rows[1:]

		//模版校验
		if len(excelTitle) != 49 {
			return errors.New("导入花名册Excel模版异常")
		}

		rankTypeList, err := GetRankTypeList()
		if err != nil {
			return errors.New("职级类型异常:" + err.Error())
		}

		rankTypeMap := make(map[int]string)
		for _, rankTypeItem := range rankTypeList {
			rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
		}

		//参数校验
		for i, row := range values {
			//每一行最后一列为空要这样判空
			if len(titleKeyMap) != len(row) && len(titleKeyMap) != len(row)-1 {
				fmt.Println("length", len(titleKeyMap), len(row))
				return errors.New(fmt.Sprintf("第%d行有数据缺失", i+2))
			}
			var rankTypeValue, rankValue string
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				fmt.Println("title-key-value", excelTitle[ii], key, value)
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

				//结合每个字段是否为空判断(最后一列为空的话这种方式判断不出来)
				noRequired := []string{"社保电脑号",
					"公积金账号",
					"社保公积金缴纳地",
					"费用科目",
					"试用期",
					"转正日期",
					"职称/技能证书",
					"联系人常住地址",
					"合同公司",
					"合同类型",
					"续签次数",
					"合同附件"}
				if value == "" && utils.InArray(noRequired, excelTitle[ii]) {
					return errors.New(fmt.Sprintf("%s不能为空", excelTitle[ii]))
				}

				if key == "rank_type" {
					rankTypeValue = value
				}

				if key == "rank" {
					rankValue = value
				}

				err = checkImportParam(key, value, rankTypeList)
				if err != nil {
					return err
				}
			}

			err = checkImportParamRank(rankTypeValue, rankValue, rankTypeMap)
			if err != nil {
				return err
			}
		}

		//获取部门 职位信息
		departmentMaps := make(map[string]int)
		positionMaps := make(map[string]int)
		var ds []weChat.WcDepartment
		var ps []weChat.WcPosition

		tx.Table(weChat.WcDepartment{}.TableName()).Select("id,name").Find(&ds)
		for _, dsItem := range ds {
			departmentMaps[dsItem.Name] = int(dsItem.ID)
		}

		tx.Table(weChat.WcPosition{}.TableName()).Select("id,name").Find(&ps)
		for _, psItem := range ps {
			positionMaps[psItem.Name] = int(psItem.ID)
		}

		//更新操作
		for _, row := range values {
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				// 更新部门信息
				if key == "department" && value != "" {
					departmentsMultiple := strings.Split(value, ";")
					departmentsParents := make(map[string]string)
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					fmt.Println("departmentsMultiple", departmentsMultiple)
					for _, departmentsSingle := range departmentsMultiple {
						departments := strings.Split(departmentsSingle, "/")
						fmt.Println("departmentsSingle", departmentsSingle)
						fmt.Println("departments", departments)
						fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
						for level, name := range departments {
							if level > 0 {
								departmentsParents[name] = departments[level-1]
							}
							_, ok := departmentMaps[name]
							if ok {
								continue
							}
							var parentId int
							if level > 0 {
								fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
								fmt.Println("level", level)
								fmt.Println("departments", departments)
								fmt.Println("name", departments[level-1])
								fmt.Println("departmentMaps", departmentMaps)
								fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
								pId, mapOk := departmentMaps[departments[level-1]]
								if mapOk {
									parentId = pId
								}
							}
							var wd weChat.WcDepartment
							zero := 0
							wd.Name = name
							wd.Parentid = &parentId
							wd.Order = &zero
							wd.CreatedAt = time.Now()
							wd.UpdatedAt = time.Now()
							tx.Table(wd.TableName()).Create(&wd)
							departmentMaps[name] = int(wd.ID)
						}
					}
				}

				// 更新职务信息
				if key == "position" && value != "" {
					positions := strings.Split(value, ";")
					for _, name := range positions {
						_, ok := positionMaps[name]
						if ok {
							continue
						}
						var wp weChat.WcPosition
						wp.Name = name
						wp.CreatedAt = time.Now()
						wp.UpdatedAt = time.Now()
						tx.Table(wp.TableName()).Create(&wp)
						positionMaps[name] = int(wp.ID)
					}
				}
			}
		}

		staffFields := []string{"name", "job_num", "userid", "mobile", "gender", "height", "weight", "birthday", "native_place", "nation", "marriage", "political_outlook", "id_number", "id_address", "household_type", "address", "social_number", "account_number", "payment_place"}
		staffJobFields := []string{"job_type", "status", "employment_date", "try_period", "formal_date", "leader" /**"department", "position",**/, "rank_type", "rank", "rank_salary", "expense_account"}
		staffBankFields := []string{"bank", "card_number"}
		staffEducationFields := []string{"education", "education_pay", "school", "date", "major", "certificate", "skill_pay"}
		staffContactFields := []string{"contact_name", "relationship", "contact_mobile", "contact_address"}
		staffAgreementFields := []string{"company", "agreement_type", "start_day", "end_day", "times"}
		configInfo := config.GetConfigInfo()

		for _, row := range values {
			var itemBase = make(map[string]interface{})
			var itemJob = make(map[string]interface{})
			var itemBank = make(map[string]interface{})
			var itemEducation = make(map[string]interface{})
			var itemContact = make(map[string]interface{})
			var itemAgreement = make(map[string]interface{})
			var name string
			var lastDepartments, positions []string
			var staffExist, staff weChat.WcStaff
			var staffJob weChat.WcStaffJob
			var staffBank weChat.WcStaffBank
			var staffEducation weChat.WcStaffEducation
			var staffContact weChat.WcStaffContact
			var staffAgreement weChat.WcStaffAgreement

			// 更新员工信息
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "name" {
					name = value
				}
				if utils.InArray(staffFields, key) {
					if key == "gender" {
						gender, _ := utils.FindKeyByValue(configInfo.StaffGender, value)
						itemBase[key] = gender
					} else if key == "nation" {
						nation, _ := utils.FindKeyByValue(configInfo.Nation, value)
						itemBase[key] = nation
					} else if key == "marriage" {
						marriage, _ := utils.FindKeyByValue(configInfo.Marriage, value)
						itemBase[key] = marriage
					} else if key == "political_outlook" {
						politicalOutlook, _ := utils.FindKeyByValue(configInfo.PoliticalOutlook, value)
						itemBase[key] = politicalOutlook
					} else if key == "household_type" {
						householdType, _ := utils.FindKeyByValue(configInfo.HouseholdType, value)
						itemBase[key] = householdType
					} else {
						itemBase[key] = value
					}
				}
				if utils.InArray(staffJobFields, key) {
					if key == "job_type" {
						jobType, _ := utils.FindKeyByValue(configInfo.StaffJobType, value)
						itemJob[key] = jobType
					} else if key == "status" {
						jobStatus, _ := utils.FindKeyByValue(configInfo.StaffJobStatus, value)
						itemJob[key] = jobStatus
					} else if key == "try_period" {
						tryPeriod, _ := utils.FindKeyByValue(configInfo.StaffJobTryPeriod, value)
						itemJob[key] = tryPeriod
					} else if key == "leader" {
						if value != "" {
							itemJob["leader_id"] = 1 //todo
						} else {
							itemJob["leader_id"] = 0
						}
					} else if key == "rank_type" {
						itemJob[key] = 0 //todo
					} else if key == "rank" {
						itemJob[key] = 0 //todo
					} else if key == "expense_account" {
						expenseAccount, _ := utils.FindKeyByValue(configInfo.ExpenseAccount, value)
						itemJob[key] = expenseAccount
					} else {
						itemJob[key] = value
					}
				}
				if utils.InArray(staffBankFields, key) {
					itemBank[key] = value
				}
				if utils.InArray(staffEducationFields, key) {
					if key == "education" {
						education, _ := utils.FindKeyByValue(configInfo.Education, value)
						itemEducation[key] = education
					} else {
						itemEducation[key] = value
					}
				}
				if utils.InArray(staffContactFields, key) {
					if key == "contact_name" {
						itemContact["name"] = value
					} else if key == "relationship" {
						relationship, _ := utils.FindKeyByValue(configInfo.Relationship, value)
						itemContact[key] = relationship
					} else if key == "contact_mobile" {
						itemContact["mobile"] = value
					} else if key == "contact_address" {
						itemContact["address"] = value
					} else {
						itemContact[key] = value
					}
				}
				if utils.InArray(staffAgreementFields, key) {
					if key == "agreement_type" {
						agreementType, _ := utils.FindKeyByValue(configInfo.AgreementType, value)
						itemAgreement[key] = agreementType
					} else {
						itemAgreement[key] = value
					}
				}

				if key == "department" {
					departmentMany := strings.Split(value, ";")
					for _, departmentOne := range departmentMany {
						departmentOneItems := strings.Split(departmentOne, "/")
						lastDepartment := departmentOneItems[len(departmentOneItems)-1]
						lastDepartments = append(lastDepartments, lastDepartment)
					}
				}
				if key == "position" {
					positions = strings.Split(value, ";")
				}
			}

			{
				itemBase["created_at"] = now
				itemBase["updated_at"] = now

				tx.Table(staff.TableName()).Where("name=?", name).First(&staffExist)
				if staffExist.ID == 0 {
					// 新增员工
					cErr := tx.Table(staff.TableName()).Create(&itemBase).Error
					if cErr != nil {
						return cErr
					}
				} else {
					// 更新员工
					cErr := tx.Table(staff.TableName()).Omit("name,created_at").Where("id=?", staffExist.ID).Updates(itemBase).Error
					if cErr != nil {
						return cErr
					}
				}

				tx.Table(staff.TableName()).Where("name=?", name).First(&staff)

				// 更新员工部门信息
				if len(lastDepartments) > 0 {
					tx.Where("staff_id=?", staff.ID).Unscoped().Delete(&weChat.WcStaffDepartment{})
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					fmt.Println("lastDepartments", lastDepartments)
					for _, lastDepartment := range lastDepartments {
						var sdItem = make(map[string]interface{})
						departmentId, ok := departmentMaps[lastDepartment]
						fmt.Println("lastDepartment", lastDepartment)
						fmt.Println("departmentId", departmentId)
						fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
						if ok {
							sdItem["staff_id"] = staff.ID
							sdItem["department_id"] = departmentId
							sdItem["created_at"] = now
							sdItem["updated_at"] = now
							sdErr := tx.Table(weChat.WcStaffDepartment{}.TableName()).Create(&sdItem).Error
							if sdErr != nil {
								return sdErr
							}
						}
					}
				}

				// 更新员工职务信息
				if len(positions) > 0 {
					tx.Where("staff_id=?", staff.ID).Unscoped().Delete(&weChat.WcStaffPosition{})
					for _, pItem := range positions {
						positionId, ok := positionMaps[pItem]
						var spItem = make(map[string]interface{})
						if ok {
							spItem["staff_id"] = staff.ID
							spItem["position_id"] = positionId
							spItem["created_at"] = now
							spItem["updated_at"] = now
							spErr := tx.Table(weChat.WcStaffPosition{}.TableName()).Create(&spItem).Error
							if spErr != nil {
								return spErr
							}
						}
					}
				}
			}

			{
				itemJob["created_at"] = now
				itemJob["updated_at"] = now

				tx.Table(staffJob.TableName()).Where("staff_id=?", staff.ID).First(&staffJob)
				if staffJob.ID == 0 {
					cErr := tx.Table(staffJob.TableName()).Create(&itemJob).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffJob.TableName()).Omit("name,created_at").Where("id=?", staffJob.ID).Updates(itemJob).Error
					if cErr != nil {
						return cErr
					}
				}
			}

			{
				itemBank["created_at"] = now
				itemBank["updated_at"] = now

				tx.Table(staffBank.TableName()).Where("staff_id=?", staff.ID).First(&staffBank)
				if staffBank.ID == 0 {
					cErr := tx.Table(staffBank.TableName()).Create(&itemBank).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffBank.TableName()).Omit("name,created_at").Where("id=?", staffBank.ID).Updates(itemBank).Error
					if cErr != nil {
						return cErr
					}
				}
			}

			{
				itemEducation["created_at"] = now
				itemEducation["updated_at"] = now

				tx.Table(staffEducation.TableName()).Where("staff_id=?", staff.ID).First(&staffEducation)
				if staffEducation.ID == 0 {
					cErr := tx.Table(staffJob.TableName()).Create(&itemEducation).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffEducation.TableName()).Omit("name,created_at").Where("id=?", staffEducation.ID).Updates(itemEducation).Error
					if cErr != nil {
						return cErr
					}
				}
			}

			{
				itemContact["created_at"] = now
				itemContact["updated_at"] = now

				tx.Table(staffContact.TableName()).Where("staff_id=?", staff.ID).First(&staffContact)
				if staffContact.ID == 0 {
					cErr := tx.Table(staffContact.TableName()).Create(&itemContact).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffContact.TableName()).Omit("name,created_at").Where("id=?", staffContact.ID).Updates(itemContact).Error
					if cErr != nil {
						return cErr
					}
				}
			}

			{
				itemAgreement["created_at"] = now
				itemAgreement["updated_at"] = now

				tx.Table(staffAgreement.TableName()).Where("staff_id=?", staff.ID).First(&staffAgreement)
				if staffJob.ID == 0 {
					cErr := tx.Table(staffAgreement.TableName()).Create(&itemAgreement).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffAgreement.TableName()).Omit("name,created_at").Where("id=?", staffAgreement.ID).Updates(itemAgreement).Error
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
func checkImportParam(key, value string, rankTypeList []weChat2.WcRankTypeResponse) error {
	configInfo := config.GetConfigInfo()
	if key == "gender" && !utils.InArray(configInfo.StaffGender, value) {
		return errors.New("性别异常:" + value)
	}

	if key == "nation" && !utils.InArray(configInfo.Nation, value) {
		return errors.New("民族异常:" + value)
	}

	if key == "marriage" && !utils.InArray(configInfo.Marriage, value) {
		return errors.New("婚姻状况异常:" + value)
	}

	if key == "political_outlook" && !utils.InArray(configInfo.PoliticalOutlook, value) {
		return errors.New("政治面貌异常:" + value)
	}

	if key == "household_type" && !utils.InArray(configInfo.HouseholdType, value) {
		return errors.New("户口类型异常:" + value)
	}

	if key == "job_type" && !utils.InArray(configInfo.StaffJobType, value) {
		return errors.New("员工类型异常:" + value)
	}

	if key == "status" && !utils.InArray(configInfo.StaffJobStatus, value) {
		return errors.New("员工状态异常:" + value)
	}

	if key == "try_period" && !utils.InArray(configInfo.StaffJobTryPeriod, value) {
		return errors.New("试用期异常:" + value)
	}

	if key == "rank_type" {
		var rankTypes []string
		for _, rankTypeItem := range rankTypeList {
			rankTypes = append(rankTypes, rankTypeItem.Name)
		}
		if !utils.InArray(rankTypes, value) {
			return errors.New("职级类型异常:" + value)
		}
	}

	if key == "education" && !utils.InArray(configInfo.Education, value) {
		return errors.New("学历异常:" + value)
	}

	if key == "relationship" && !utils.InArray(configInfo.Relationship, value) {
		return errors.New("紧急联系人关系异常:" + value)
	}

	if key == "agreement_type" && !utils.InArray(configInfo.AgreementType, value) {
		return errors.New("合同类型异常:" + value)
	}
	return nil
}

// checkImportParamRank 参数校验
func checkImportParamRank(rankTypeValue, rankValue string, rankTypeMap map[int]string) error {
	rankType := utils.GetKeyByValue(rankTypeMap, rankTypeValue)
	rankList, _, err := GetRankListByRankTypeCommon(strconv.Itoa(rankType))
	if err != nil {
		return errors.New("职级异常:" + err.Error())
	}
	var rankArray []string
	for _, rankTypeItem := range rankList {
		rankArray = append(rankArray, rankTypeItem.Name)
	}

	fmt.Println("rankArray-rankType-value", rankArray, rankType, rankValue)

	if !utils.InArray(rankArray, rankValue) {
		return errors.New("职级异常:" + rankValue)
	}

	return nil
}

func (wcStaffService *WcStaffService) AssembleStaffList(staffs []weChat.WcStaff) (newStaffs []weChat2.WcStaffResponse, err error) {
	var newStaff weChat2.WcStaffResponse
	configInfo := config.GetConfigInfo()

	for _, item := range staffs {
		newStaff.WcStaff = item
		gender, _ := utils.Find(configInfo.StaffGender, *item.Gender)
		newStaff.GenderText = gender
		householdTypeText, _ := utils.Find(configInfo.HouseholdType, *item.HouseholdType)
		newStaff.HouseholdTypeText = householdTypeText
		nationText, _ := utils.Find(configInfo.Nation, *item.Nation)
		newStaff.NationText = nationText
		marriageText, _ := utils.Find(configInfo.Marriage, *item.Marriage)
		newStaff.MarriageText = marriageText
		politicalOutlookText, _ := utils.Find(configInfo.PoliticalOutlook, *item.PoliticalOutlook)
		newStaff.PoliticalOutlookText = politicalOutlookText

		newStaffs = append(newStaffs, newStaff)
	}
	return
}

func (wcStaffService *WcStaffService) AssembleStaff(staff weChat.WcStaff) (newStaff weChat2.WcStaffResponse, err error) {
	configInfo := config.GetConfigInfo()
	gender, _ := utils.Find(configInfo.StaffGender, *staff.Gender)
	newStaff.GenderText = gender
	householdTypeText, _ := utils.Find(configInfo.HouseholdType, *staff.HouseholdType)
	newStaff.HouseholdTypeText = householdTypeText
	nationText, _ := utils.Find(configInfo.Nation, *staff.Nation)
	newStaff.NationText = nationText
	marriageText, _ := utils.Find(configInfo.Marriage, *staff.Marriage)
	newStaff.MarriageText = marriageText
	politicalOutlookText, _ := utils.Find(configInfo.PoliticalOutlook, *staff.PoliticalOutlook)
	newStaff.PoliticalOutlookText = politicalOutlookText
	newStaff.WcStaff = staff

	return
}
