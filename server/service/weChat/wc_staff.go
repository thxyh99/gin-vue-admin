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

type WcStaffService struct {
}

type TransferParams struct {
	StaffId     int
	ConfigInfo  config.CommonConfig
	Column      string
	Value       interface{}
	StaffMap    map[int]string
	RankTypeMap map[int]string
	RankMap     map[int]string
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
func (wcStaffService *WcStaffService) GetWcStaffInfoList(info weChatReq.WcStaffSearch) (wcStaffStatisticsResponse weChat2.WcStaffStatisticsResponse, wcStaffResponse []weChat2.WcStaffResponse, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaff{})
	var wcStaffs []weChat.WcStaff
	where := `1`
	//选择员工
	if info.StaffId != nil && *info.StaffId > 0 {
		db = db.Where("id = ?", *info.StaffId)
		where += fmt.Sprintf(" AND a.id = %d ", *info.StaffId)
	}
	//选择部门
	if info.DepartmentIds != nil && len(info.DepartmentIds) > 0 {
		var staffDepartments []weChat.WcStaffDepartment
		global.GVA_DB.Where("department_id IN (?)", info.DepartmentIds).Find(&staffDepartments)
		fmt.Println("staffDepartments", staffDepartments)
		if len(staffDepartments) == 0 {
			db = db.Where("id = -1")
			where += fmt.Sprintf(" AND a.id = -1 ")
		} else {
			where += fmt.Sprintf(" AND a.id IN (")
			var ids []int
			for _, staffDepartment := range staffDepartments {
				ids = append(ids, *staffDepartment.StaffId)
				where += strconv.Itoa(*staffDepartment.StaffId) + ","
			}
			db = db.Where("id IN (?)", ids)
			where += "0)"
		}
	}
	//入职容大日期
	if info.EmploymentDateStart != "" && info.EmploymentDateEnd != "" {
		var staffJobs []weChat.WcStaffJob
		global.GVA_DB.Where("employment_date >= ? AND employment_date <= ?", info.EmploymentDateStart, info.EmploymentDateEnd).Find(&staffJobs)
		if len(staffJobs) == 0 {
			db = db.Where("id = -1")
			where += fmt.Sprintf(" AND a.id = -1 ")
		} else {
			var ids []int
			where += " AND a.id IN("
			for _, staffJob := range staffJobs {
				ids = append(ids, *staffJob.StaffId)
				where += strconv.Itoa(*staffJob.StaffId) + ","
			}
			db = db.Where("id IN (?)", ids)
			where += "0)"
		}
	}
	//历史日期(快照节点):入职容大日期<=DATE && (离职日期为空 || 离职日期>DATE)
	if info.HistoryDate != "" {
		var staffJobs []weChat.WcStaffJob
		global.GVA_DB.Debug().Where("employment_date <= ? AND (leave_date > ? OR leave_date IS NULL)", info.HistoryDate, info.HistoryDate).Find(&staffJobs)
		if len(staffJobs) == 0 {
			db = db.Where("id = -1")
			where += fmt.Sprintf(" AND a.id = -1 ")
		} else {
			var ids []int
			where += " AND a.id IN("
			for _, staffJob := range staffJobs {
				ids = append(ids, *staffJob.StaffId)
				where += strconv.Itoa(*staffJob.StaffId) + ","
			}
			db = db.Where("id IN (?)", ids)
			where += "0)"
		}
	}
	// 添加员工名称、员工工号、手机模糊查询
	if info.Keyword != "" {
		keyword := "%" + info.Keyword + "%"
		db = db.Where("(name LIKE ? OR job_num LIKE ? OR mobile LIKE ?)", keyword, keyword, keyword)
		where += fmt.Sprintf(" AND (a.name LIKE '%s' OR a.job_num LIKE '%s' OR a.mobile LIKE '%s')", keyword, keyword, keyword)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Debug().Find(&wcStaffs).Error

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("GetWcStaffInfoList err", err)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	wcStaffResponse, err = wcStaffService.AssembleStaffList(wcStaffs)
	wcStaffStatisticsResponse = getStaffStatistics(where)

	return
}

// getStaffStatistics 统计数据
func getStaffStatistics(where string) (wcStaffStatisticsResponse weChat2.WcStaffStatisticsResponse) {
	var tableMap []map[string]interface{}
	var onJobCount, fullTimeCount, partTimeCount, followCount, replaceCount, retireCount, outsourcingCount, toBeEmployedCount, probationCount, formalCount, toBeDepartedCount, todoCount int
	fields := `a.id, b.type, b.status`
	sql := fmt.Sprintf(`SELECT %s FROM wc_staff AS a	
          LEFT JOIN wc_staff_job AS b ON a.id = b.staff_id AND b.deleted_at IS NULL 
          WHERE %s `, fields, where)

	global.GVA_DB.Debug().Raw(sql).Scan(&tableMap)
	for _, table := range tableMap {
		if intVal64, ok := table["type"].(int64); ok {
			switch intVal64 {
			case 1:
				fullTimeCount++
				break
			case 2:
				partTimeCount++
				break
			case 3:
				followCount++
				break
			case 4:
				replaceCount++
				break
			case 5:
				retireCount++
				break
			case 6:
				outsourcingCount++
				break
			}
		}
		if intVal64, ok := table["status"].(int64); ok {
			switch intVal64 {
			case 1:
				toBeEmployedCount++
				onJobCount++
				break
			case 2:
				probationCount++
				onJobCount++
				break
			case 3:
				formalCount++
				onJobCount++
				break
			}

		}
	}

	//统计代办事项个数
	var todoList []weChat2.TodoList
	var todoListAll, todoListAfter, todoListNow, todoListBefore []weChat2.TodoItem
	var numAfter, numNow, numBefore int

	//试用期到期前30天(拟转正日期)
	var resultFormal []map[string]interface{}
	fieldsTodoFormal := `a.id,a.name,b.presume_formal_date,CURDATE() AS today`
	whereTodoFormal := where + " AND b.presume_formal_date IS NOT NULL AND DATEDIFF(b.presume_formal_date, CURDATE()) <= 30"
	sqlTodoFormal := fmt.Sprintf(`SELECT %s FROM wc_staff AS a
          LEFT JOIN wc_staff_job AS b ON a.id = b.staff_id AND b.deleted_at IS NULL 
          WHERE %s `, fieldsTodoFormal, whereTodoFormal)
	global.GVA_DB.Debug().Raw(sqlTodoFormal).Scan(&resultFormal)

	for _, row := range resultFormal {
		todoCount++
		var staffId int
		var staffName, value, now string
		if id, ok := row["id"].(uint32); ok {
			staffId = int(id)
		}
		if name, ok := row["name"].(string); ok {
			staffName = name
		}
		if presumeFormalDate, ok := row["presume_formal_date"].(time.Time); ok {
			value = presumeFormalDate.Format("2006-01-02")
		}
		if today, ok := row["today"].(time.Time); ok {
			now = today.Format("2006-01-02")
		}
		if value == now {
			numNow++
			todoListNow = append(todoListNow, weChat2.TodoItem{
				Title:   staffName + "试用期到期",
				Label:   "转正日期",
				Value:   value,
				StaffId: staffId,
			})
		} else if value < now {
			numAfter++
			todoListAfter = append(todoListAfter, weChat2.TodoItem{
				Title:   staffName + "试用期到期",
				Label:   "转正日期",
				Value:   value,
				StaffId: staffId,
			})
		} else {
			numBefore++
			todoListBefore = append(todoListBefore, weChat2.TodoItem{
				Title:   staffName + "试用期到期",
				Label:   "转正日期",
				Value:   value,
				StaffId: staffId,
			})
		}
		todoListAll = append(todoListAll, weChat2.TodoItem{
			Title:   staffName + "试用期到期",
			Label:   "转正日期",
			Value:   value,
			StaffId: staffId,
		})

	}

	//劳动合同到期前60天
	var resultAgreement []map[string]interface{}
	fieldsTodoAgreement := `a.id,a.name,b.end_day,CURDATE() AS today`
	whereTodoAgreement := where + " AND b.type=1 AND b.is_renew = 0 AND b.end_day IS NOT NULL AND DATEDIFF(b.end_day, CURDATE()) <= 60"
	sqlTodoAgreement := fmt.Sprintf(`SELECT %s FROM wc_staff AS a
          LEFT JOIN wc_staff_agreement AS b ON a.id = b.staff_id AND b.deleted_at IS NULL 
          WHERE %s `, fieldsTodoAgreement, whereTodoAgreement)
	global.GVA_DB.Debug().Raw(sqlTodoAgreement).Scan(&resultAgreement)
	for _, row := range resultAgreement {
		todoCount++
		var staffId int
		var staffName, value, now string
		if id, ok := row["id"].(uint32); ok {
			staffId = int(id)
		}
		if name, ok := row["name"].(string); ok {
			staffName = name
		}
		if endDay, ok := row["end_day"].(time.Time); ok {
			value = endDay.Format("2006-01-02")
		}
		if today, ok := row["today"].(time.Time); ok {
			now = today.Format("2006-01-02")
		}
		if value == now {
			numNow++
			todoListNow = append(todoListNow, weChat2.TodoItem{
				Title:   staffName + "合同到期",
				Label:   "合同到期日",
				Value:   value,
				StaffId: staffId,
			})
		} else if value < now {
			numAfter++
			todoListAfter = append(todoListAfter, weChat2.TodoItem{
				Title:   staffName + "合同到期",
				Label:   "合同到期日",
				Value:   value,
				StaffId: staffId,
			})
		} else {
			numBefore++
			todoListBefore = append(todoListBefore, weChat2.TodoItem{
				Title:   staffName + "合同到期",
				Label:   "合同到期日",
				Value:   value,
				StaffId: staffId,
			})
		}
		todoListAll = append(todoListAll, weChat2.TodoItem{
			Title:   staffName + "合同到期",
			Label:   "合同到期日",
			Value:   value,
			StaffId: staffId,
		})

	}

	//健康证到期前30天
	var resultHealth []map[string]interface{}
	fieldsTodoHealth := `a.id,a.name,b.health_end,CURDATE() AS today`
	whereTodoHealth := where + " AND b.health_end IS NOT NULL AND DATEDIFF(b.health_end, CURDATE()) <= 30"
	sqlTodoHealth := fmt.Sprintf(`SELECT %s FROM wc_staff AS a
          LEFT JOIN wc_staff_job AS b ON a.id = b.staff_id AND b.deleted_at IS NULL 
          WHERE %s `, fieldsTodoHealth, whereTodoHealth)
	global.GVA_DB.Debug().Raw(sqlTodoHealth).Scan(&resultHealth)
	for _, row := range resultHealth {
		todoCount++
		var staffId int
		var staffName, value, now string
		if id, ok := row["id"].(uint32); ok {
			staffId = int(id)
		}
		if name, ok := row["name"].(string); ok {
			staffName = name
		}
		if healthEnd, ok := row["health_end"].(time.Time); ok {
			value = healthEnd.Format("2006-01-02")
		}
		if today, ok := row["today"].(time.Time); ok {
			now = today.Format("2006-01-02")
		}
		if value == now {
			numNow++
			todoListNow = append(todoListNow, weChat2.TodoItem{
				Title:   staffName + "健康证到期",
				Label:   "健康证到期日",
				Value:   value,
				StaffId: staffId,
			})
		} else if value < now {
			numAfter++
			todoListAfter = append(todoListAfter, weChat2.TodoItem{
				Title:   staffName + "健康证到期",
				Label:   "健康证到期日",
				Value:   value,
				StaffId: staffId,
			})
		} else {
			numBefore++
			todoListBefore = append(todoListBefore, weChat2.TodoItem{
				Title:   staffName + "健康证到期",
				Label:   "健康证到期日",
				Value:   value,
				StaffId: staffId,
			})
		}
		todoListAll = append(todoListAll, weChat2.TodoItem{
			Title:   staffName + "健康证到期",
			Label:   "健康证到期日",
			Value:   value,
			StaffId: staffId,
		})

	}

	//离职日当天
	var resultLeave []map[string]interface{}
	fieldsTodoLeave := `a.id,a.name,b.leave_date,CURDATE() AS today`
	whereTodoLeave := where + " AND b.leave_date IS NOT NULL AND DATEDIFF(b.leave_date, CURDATE()) = 0"
	sqlTodoLeave := fmt.Sprintf(`SELECT %s FROM wc_staff AS a
          LEFT JOIN wc_staff_job AS b ON a.id = b.staff_id AND b.deleted_at IS NULL 
          WHERE %s `, fieldsTodoLeave, whereTodoLeave)
	global.GVA_DB.Debug().Raw(sqlTodoLeave).Scan(&resultLeave)
	for _, row := range resultLeave {
		todoCount++
		var staffId int
		var staffName, value, now string
		if id, ok := row["id"].(uint32); ok {
			staffId = int(id)
		}
		if name, ok := row["name"].(string); ok {
			staffName = name
		}
		if leaveDate, ok := row["leave_date"].(time.Time); ok {
			value = leaveDate.Format("2006-01-02")
		}
		if today, ok := row["today"].(time.Time); ok {
			now = today.Format("2006-01-02")
		}
		if value == now {
			numNow++
			todoListNow = append(todoListNow, weChat2.TodoItem{
				Title:   staffName + "离职到期",
				Label:   "离职到期日",
				Value:   value,
				StaffId: staffId,
			})
		} else if value < now {
			numAfter++
			todoListAfter = append(todoListAfter, weChat2.TodoItem{
				Title:   staffName + "离职到期",
				Label:   "离职到期日",
				Value:   value,
				StaffId: staffId,
			})
		} else {
			numBefore++
			todoListBefore = append(todoListBefore, weChat2.TodoItem{
				Title:   staffName + "离职到期",
				Label:   "离职到期日",
				Value:   value,
				StaffId: staffId,
			})
		}
		todoListAll = append(todoListAll, weChat2.TodoItem{
			Title:   staffName + "离职到期",
			Label:   "离职到期日",
			Value:   value,
			StaffId: staffId,
		})

	}

	todoList = append(todoList, weChat2.TodoList{
		Type:  0,
		Total: len(todoListAll),
		Label: "全部",
		List:  todoListAll,
	})
	todoList = append(todoList, weChat2.TodoList{
		Type:  1,
		Total: len(todoListAfter),
		Label: "超期未处理",
		List:  todoListAfter,
	})
	todoList = append(todoList, weChat2.TodoList{
		Type:  2,
		Total: len(todoListNow),
		Label: "今日到期",
		List:  todoListNow,
	})
	todoList = append(todoList, weChat2.TodoList{
		Type:  3,
		Total: len(todoListBefore),
		Label: "提前提醒",
		List:  todoListBefore,
	})

	wcStaffStatisticsResponse = weChat2.WcStaffStatisticsResponse{
		OnJobCount:        onJobCount,
		FullTimeCount:     fullTimeCount,
		PartTimeCount:     partTimeCount,
		FollowCount:       followCount,
		ReplaceCount:      replaceCount,
		RetireCount:       retireCount,
		OutsourcingCount:  outsourcingCount,
		ToBeEmployedCount: toBeEmployedCount,
		ProbationCount:    probationCount,
		FormalCount:       formalCount,
		ToBeDepartedCount: toBeDepartedCount,
		TodoCount:         todoCount,
		TodoList:          todoList,
	}
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

	var titleKeyMap = make(map[string]string, len(templateInfoMap))
	for key, title := range templateInfoMap {
		title = utils.FilterBreaksSpaces(title)
		titleKeyMap[title] = key
	}

	db := global.GVA_DB
	if template.DBName != "" {
		db = global.MustGetGlobalDBByDBName(template.DBName)
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	configInfo := config.GetConfigInfo()

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]
		fmt.Println(len(excelTitle))
		//模版校验
		if len(excelTitle) != 58 {
			return errors.New("导入花名册Excel模版异常")
		}

		rankTypeList, err := GetRankTypeList()
		if err != nil {
			return errors.New("职级类型异常:" + err.Error())
		}

		var rankTypes []string
		for _, rankTypeItem := range rankTypeList {
			rankTypes = append(rankTypes, rankTypeItem.Name)
		}

		rankTypeMap := make(map[int]string)
		for _, rankTypeItem := range rankTypeList {
			rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
		}

		//参数校验
		for _, row := range values {
			var rankTypeValue, rankValue string
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]

				//结合每个字段是否为空判断(最后一列为空的话这种方式判断不出来)
				noRequired := []string{"社保电脑号",
					"公积金账号",
					"社保公积金缴纳地",
					"费用科目",
					"试用期",
					"转正日期",
					"离职日期",
					"职称证书",
					"技能证书",
					"联系人常住地址",
					"合同公司",
					"合同类型",
					"续签次数",
				}
				if value == "" && !utils.InArray(noRequired, excelTitle[ii]) {
					return errors.New(fmt.Sprintf("%s不能为空", excelTitle[ii]))
				}

				if key == "rank_type" {
					rankTypeValue = value
				}

				if key == "rank" {
					rankValue = value
				}

				err = checkImportParam(key, value, configInfo, rankTypes)
				if err != nil {
					return err
				}
			}

			_, _, err = checkAndReturnRank(rankTypeValue, rankValue, rankTypeMap)
			if err != nil {
				return err
			}
		}

		departmentFirstMaps := make(map[string]int)
		departmentSecondMaps := make(map[int][]map[string]int)
		positionMaps := make(map[string]int)
		var ds []weChat.WcDepartment
		var ps []weChat.WcPosition

		tx.Table(weChat.WcDepartment{}.TableName()).Select("id,name,parentid").Find(&ds)
		for _, dsItem := range ds {
			fmt.Println("dsItem", dsItem, *dsItem.Parentid)
			if *dsItem.Parentid == 0 {
				departmentFirstMaps[dsItem.Name] = int(dsItem.ID)
			} else {
				departmentItemMaps := make(map[string]int)
				departmentItemMaps[dsItem.Name] = int(dsItem.ID)
				departmentSecondMaps[*dsItem.Parentid] = append(departmentSecondMaps[*dsItem.Parentid], departmentItemMaps)
			}
		}
		fmt.Println("======================================")
		fmt.Println("departmentFirstMaps", departmentFirstMaps)
		fmt.Println("departmentSecondMaps", departmentSecondMaps)
		fmt.Println("======================================")

		tx.Table(weChat.WcPosition{}.TableName()).Select("id,name").Find(&ps)
		for _, psItem := range ps {
			positionMaps[psItem.Name] = int(psItem.ID)
		}

		//更新操作
		for _, row := range values {
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]

				// 更新岗位信息
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
		staffJobFields := []string{"job_type", "status", "employment_date", "employment_headquarter_date", "try_period", "formal_date", "presume_formal_date", "leave_date", "health_start", "health_end", "leader", "level", "io_type", "rank_type", "rank", "rank_salary", "expense_account"}
		staffBankFields := []string{"bank", "card_number"}
		staffEducationFields := []string{"education", "education_pay", "school", "date", "major", "professional_certificate", "skill_certificate", "skill_pay"}
		staffContactFields := []string{"contact_name", "relationship", "contact_mobile", "contact_address"}
		staffAgreementFields := []string{"company", "agreement_type", "start_day", "end_day", "times", "remark"}

		for _, row := range values {
			var itemBase = make(map[string]interface{})
			var itemJob = make(map[string]interface{})
			var itemBank = make(map[string]interface{})
			var itemEducation = make(map[string]interface{})
			var itemContact = make(map[string]interface{})
			var itemAgreement = make(map[string]interface{})
			var name, rankValue, rankTypeValue string
			var positions []string
			var departmentFirstId = 0
			var departmentSecondId = 0
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
					} else if key == "payment_place" {
						paymentPlace, _ := utils.FindKeyByValue(configInfo.PaymentPlace, value)
						itemBase[key] = paymentPlace
					} else {
						itemBase[key] = value
					}
				}
				if utils.InArray(staffJobFields, key) {
					if key == "job_type" {
						jobType, _ := utils.FindKeyByValue(configInfo.StaffJobType, value)
						itemJob["type"] = jobType
					} else if key == "status" {
						jobStatus, _ := utils.FindKeyByValue(configInfo.StaffJobStatus, value)
						itemJob[key] = jobStatus
					} else if key == "try_period" {
						tryPeriod, _ := utils.FindKeyByValue(configInfo.StaffJobTryPeriod, value)
						itemJob[key] = tryPeriod
					} else if key == "level" {
						level, _ := utils.FindKeyByValue(configInfo.Level, value)
						itemJob[key] = level
					} else if key == "io_type" {
						ioType, _ := utils.FindKeyByValue(configInfo.IoType, value)
						itemJob[key] = ioType
					} else if key == "leader" {
						if value != "" {
							var staffLeader weChat.WcStaff
							tx.Table(staffLeader.TableName()).Where("name=?", value).First(&staffLeader)
							if staffLeader.ID > 0 {
								itemJob["leader_id"] = staffLeader.ID
							} else {
								itemJob["leader_id"] = 0 //初始化用户的领导可能还没插入 再次初始化一次就好了
							}
						} else {
							itemJob["leader_id"] = 0
						}
					} else if key == "rank_type" {
						rankTypeValue = value
					} else if key == "rank" {
						rankValue = value
					} else if key == "expense_account" {
						expenseAccount, _ := utils.FindKeyByValue(configInfo.ExpenseAccount, value)
						itemJob[key] = expenseAccount
					} else {
						if value == "" {
							itemJob[key] = nil
						} else {
							itemJob[key] = value
						}
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
						if value == "" {
							itemEducation[key] = nil
						} else {
							itemEducation[key] = value
						}
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
						itemAgreement["type"] = agreementType
					} else if key == "company" {
						agreementCompany, _ := utils.FindKeyByValue(configInfo.AgreementCompany, value)
						itemAgreement[key] = agreementCompany
					} else if key == "times" {
						renewTimes, _ := utils.FindKeyByValue(configInfo.RenewTimes, value)
						itemAgreement[key] = renewTimes
					} else {
						itemAgreement[key] = value
					}
				}

				if key == "department_first" {
					value = utils.FilterBreaksSpaces(value)
					if dfId, ok := departmentFirstMaps[value]; ok {
						departmentFirstId = dfId
					} else {
						var wd weChat.WcDepartment
						zero := 0
						wd.Name = value
						wd.Parentid = &zero
						wd.CreatedAt = time.Now()
						wd.UpdatedAt = time.Now()
						tx.Table(wd.TableName()).Create(&wd)
						departmentFirstMaps[value] = int(wd.ID)
						departmentFirstId = int(wd.ID)
						fmt.Println("departmentFirstId000000", departmentFirstId)
					}
				}
				if key == "department_second" && value != "" {
					value = utils.FilterBreaksSpaces(value)
					if dsMaps, dsOk := departmentSecondMaps[departmentFirstId]; dsOk { //一级部门已存在
						for _, dsMap := range dsMaps {
							if dsId, ok := dsMap[value]; ok { //已存在二级部门
								fmt.Println("******************1", value, dsMap, ok)
								departmentSecondId = dsId
								break
							} else { //不存在二级部门
								fmt.Println("******************2", value, dsMap, ok)
								var wdExist weChat.WcDepartment
								tx.Table(wdExist.TableName()).Where("parentid=? AND name=?", departmentFirstId, value).First(&wdExist)
								if wdExist.ID != 0 { //刚插入进去的直接返回ID
									departmentSecondId = int(wdExist.ID)
									fmt.Println("departmentSecondId11111", int(wdExist.ID))
								} else { //不存在的插入数据
									var wd weChat.WcDepartment
									wd.Name = value
									wd.Parentid = &departmentFirstId
									wd.CreatedAt = time.Now()
									wd.UpdatedAt = time.Now()
									tx.Table(wd.TableName()).Create(&wd)

									departmentItemMaps := make(map[string]int)
									departmentItemMaps[wd.Name] = int(wd.ID)
									departmentSecondId = int(wd.ID)
									departmentSecondMaps[departmentFirstId] = append(departmentSecondMaps[departmentFirstId], departmentItemMaps)
									fmt.Println("departmentSecondId222222", int(wd.ID))
								}
							}
						}
					} else { //一级部门刚添加的场景
						var wdExist weChat.WcDepartment
						tx.Table(wdExist.TableName()).Where("parentid=? AND name=?", departmentFirstId, value).First(&wdExist)
						if wdExist.ID != 0 { //刚插入进去的直接返回ID
							departmentSecondId = int(wdExist.ID)
							fmt.Println("departmentSecondId33333", int(wdExist.ID))
						} else { //不存在的插入数据
							var wd weChat.WcDepartment
							wd.Name = value
							wd.Parentid = &departmentFirstId
							wd.CreatedAt = time.Now()
							wd.UpdatedAt = time.Now()
							tx.Table(wd.TableName()).Create(&wd)

							departmentItemMaps := make(map[string]int)
							departmentItemMaps[wd.Name] = int(wd.ID)
							departmentSecondId = int(wd.ID)
							departmentSecondMaps[departmentFirstId] = append(departmentSecondMaps[departmentFirstId], departmentItemMaps)
							fmt.Println("departmentSecondId44444", int(wd.ID))
						}
					}
				}

				if key == "position" {
					positions = strings.Split(value, ";")
				}
			}

			{
				itemBase["created_at"] = now
				itemBase["updated_at"] = now

				tx.Table(staffExist.TableName()).Where("name=?", name).First(&staffExist)
				if staffExist.ID == 0 {
					// 新增员工
					cErr := tx.Table(staffExist.TableName()).Create(&itemBase).Error
					if cErr != nil {
						return cErr
					}
				} else {
					// 更新员工
					cErr := tx.Table(staffExist.TableName()).Omit("name,created_at").Where("id=?", staffExist.ID).Updates(itemBase).Error
					if cErr != nil {
						return cErr
					}
				}

				tx.Table(staff.TableName()).Where("name=?", name).First(&staff)

				//更新员工部门信息
				if departmentSecondId > 0 {
					tx.Where("staff_id=?", staff.ID).Unscoped().Delete(&weChat.WcStaffDepartment{})
					var sdItem = make(map[string]interface{})
					sdItem["staff_id"] = staff.ID
					sdItem["department_id"] = departmentSecondId
					sdItem["created_at"] = now
					sdItem["updated_at"] = now
					sdErr := tx.Table(weChat.WcStaffDepartment{}.TableName()).Create(&sdItem).Error
					if sdErr != nil {
						return sdErr
					}
				}

				// 更新员工岗位信息
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
				itemJob["staff_id"] = staff.ID

				rankType, rank, err := checkAndReturnRank(rankTypeValue, rankValue, rankTypeMap)
				if err != nil {
					return err
				}
				itemJob["rank_type"] = rankType
				itemJob["rank"] = rank

				tx.Table(staffJob.TableName()).Where("staff_id=?", staff.ID).First(&staffJob)
				if staffJob.ID == 0 {
					cErr := tx.Table(staffJob.TableName()).Create(&itemJob).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Debug().Table(staffJob.TableName()).Omit("created_at").Where("id=?", staffJob.ID).Updates(itemJob).Error
					if cErr != nil {
						return cErr
					}
				}
			}

			{
				itemBank["created_at"] = now
				itemBank["updated_at"] = now
				itemBank["staff_id"] = staff.ID

				tx.Table(staffBank.TableName()).Where("staff_id=?", staff.ID).First(&staffBank)
				if staffBank.ID == 0 {
					cErr := tx.Table(staffBank.TableName()).Create(&itemBank).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffBank.TableName()).Omit("created_at").Where("id=?", staffBank.ID).Updates(itemBank).Error
					if cErr != nil {
						return cErr
					}
				}
			}

			{
				itemEducation["created_at"] = now
				itemEducation["updated_at"] = now
				itemEducation["staff_id"] = staff.ID

				tx.Table(staffEducation.TableName()).Where("staff_id=?", staff.ID).First(&staffEducation)
				if staffEducation.ID == 0 {
					cErr := tx.Table(staffJob.TableName()).Create(&itemEducation).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Debug().Table(staffEducation.TableName()).Omit("created_at").Where("id=?", staffEducation.ID).Updates(itemEducation).Error
					if cErr != nil {
						return cErr
					}
				}
			}

			{
				itemContact["created_at"] = now
				itemContact["updated_at"] = now
				itemContact["staff_id"] = staff.ID

				tx.Table(staffContact.TableName()).Where("staff_id=?", staff.ID).First(&staffContact)
				if staffContact.ID == 0 {
					cErr := tx.Table(staffContact.TableName()).Create(&itemContact).Error
					if cErr != nil {
						return cErr
					}
				} else {
					cErr := tx.Table(staffContact.TableName()).Omit("created_at").Where("id=?", staffContact.ID).Updates(itemContact).Error
					if cErr != nil {
						return cErr
					}
				}
			}

			{
				itemAgreement["created_at"] = now
				itemAgreement["updated_at"] = now
				itemAgreement["staff_id"] = staff.ID

				tx.Debug().Table(staffAgreement.TableName()).Where("staff_id=? AND type =? AND start_day =? AND end_day =?", staff.ID, itemAgreement["type"], itemAgreement["start_day"], itemAgreement["end_day"]).First(&staffAgreement)
				if staffAgreement.ID == 0 {
					cErr := tx.Debug().Table(staffAgreement.TableName()).Create(&itemAgreement).Error
					if cErr != nil {
						return cErr
					}
					// 新增固定续签合同需要把之前的续签合同更新为已续签
					if itemAgreement["type"] == 1 {
						cErr := tx.Debug().Table(staffAgreement.TableName()).Where("staff_id=? AND type=? AND start_day<?", staff.ID, 1, itemAgreement["start_day"]).Updates(map[string]interface{}{
							"is_renew":   1,
							"updated_at": now,
						}).Error
						if cErr != nil {
							return cErr
						}
					}
				} else {
					cErr := tx.Debug().Table(staffAgreement.TableName()).Omit("start_day,end_day,created_at").Where("id=?", staffAgreement.ID).Updates(itemAgreement).Error
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
func checkImportParam(key, value string, configInfo config.CommonConfig, rankTypes []string) error {
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

	if key == "rank_type" && !utils.InArray(rankTypes, value) {
		return errors.New("职级类型异常:" + value)
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

	if key == "times" && !utils.InArray(configInfo.RenewTimes, value) {
		return errors.New("续签次数异常:" + value)
	}

	if key == "company" && !utils.InArray(configInfo.AgreementCompany, value) {
		return errors.New("合同公司异常:" + value)
	}

	if key == "payment_place" && !utils.InArray(configInfo.PaymentPlace, value) {
		return errors.New("社保公积金缴纳地异常:" + value)
	}

	if key == "level" && !utils.InArray(configInfo.Level, value) {
		return errors.New("层级异常:" + value)
	}

	if key == "io_type" && !utils.InArray(configInfo.IoType, value) {
		return errors.New("内外勤异常:" + value)
	}
	return nil
}

// checkAndReturnRank 参数校验
func checkAndReturnRank(rankTypeValue, rankValue string, rankTypeMap map[int]string) (rankType, rank int, err error) {
	rankType = utils.GetKeyByValue(rankTypeMap, rankTypeValue)
	rankList, _, err := GetRankListByRankTypeCommon(strconv.Itoa(rankType))
	if err != nil {
		return 0, 0, errors.New("职级异常:" + err.Error())
	}

	for _, rankItem := range rankList {
		if rankItem.Name == rankValue {
			return rankType, rankItem.ID, nil
		}
	}

	return 0, 0, errors.New("职级与职级类型不匹配:" + rankValue)
}

// ExportExcel 导出Excel
func (wcStaffService *WcStaffService) ExportExcel(templateID string, values url.Values) (file *bytes.Buffer, name string, err error) {
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
	var tableTitle []string
	for _, key := range columns {
		tableTitle = append(tableTitle, templateInfoMap[key])
	}

	// 员工名称map  直属领导
	var staffAll []weChat.WcStaff
	staffMap := make(map[int]string)
	err = global.GVA_DB.Debug().Find(&staffAll).Error
	if err != nil {
		return nil, "", err
	}
	for _, staffItem := range staffAll {
		staffMap[int(staffItem.ID)] = staffItem.Name
	}
	fmt.Println("staffMap", staffMap)
	// 配置文件
	configInfo := config.GetConfigInfo()
	// 职级类型
	rankTypeMap := make(map[int]string)
	rankMap := make(map[int]string)
	rankTypeList, err := GetRankTypeList()
	if err != nil {
		return nil, "", err
	}
	for _, item := range rankTypeList {
		rankTypeMap[item.ID] = item.Name
	}

	var tableMap []map[string]interface{}
	db := global.GVA_DB

	fields := `a.id, a.name, a.job_num, a.userid, a.mobile, a.gender, a.height, a.weight, a.birthday, a.native_place, a.nation, marriage, 
a.political_outlook, a.id_number, a.id_address, a.household_type, a.address, a.social_number, a.account_number, a.payment_place, 
b.type AS job_type, b.status, b.employment_date,b.employment_headquarter_date, b.try_period, b.formal_date, b.leader_id as leader, b.rank_type, b.rank, 
b.rank_salary, b.expense_account, c.bank, c.card_number, d.education, d.education_pay, d.school, d.date, d.major, d.professional_certificate,d.skill_certificate, 
d.skill_pay, e.name AS contact_name, e.relationship, e.mobile AS contact_mobile, e.address AS contact_address, 
f.company, f.type AS agreement_type, f.start_day, f.end_day, f.times `
	where := `1`
	keyword := values.Get("keyword")
	staffId := values.Get("staffId")
	historyDate := values.Get("historyDate")
	employmentDateRange := values.Get("employmentDateRange")
	departmentIds := values.Get("departmentIds")
	//关键字
	if keyword != "" {
		keyword = "%" + keyword + "%"
		where += fmt.Sprintf(" AND (a.name LIKE '%s' OR a.job_num LIKE '%s' OR a.mobile LIKE '%s')", keyword, keyword, keyword)
	}
	//选择员工
	if staffId != "" {
		where += fmt.Sprintf(" AND a.id = %s", staffId)
	}
	//历史日期(快照节点):入职容大日期<=DATE && (离职日期为空 || 离职日期>DATE)
	if historyDate != "" {
		var staffJobs []weChat.WcStaffJob
		global.GVA_DB.Debug().Where("employment_date <= ? AND (leave_date > ? OR leave_date IS NULL)", historyDate, historyDate).Find(&staffJobs)
		if len(staffJobs) == 0 {
			where += fmt.Sprintf(" AND a.id = -1 ")
		} else {
			var ids []int
			where += " AND a.id IN("
			for _, staffJob := range staffJobs {
				ids = append(ids, *staffJob.StaffId)
				where += strconv.Itoa(*staffJob.StaffId) + ","
			}
			db = db.Where("id IN (?)", ids)
			where += "0)"
		}
	}
	//入职容大日期
	if employmentDateRange != "" {
		employmentDateRangeArr := strings.Split(employmentDateRange, ",")
		if len(employmentDateRangeArr) == 2 {
			var staffJobs []weChat.WcStaffJob
			global.GVA_DB.Where("employment_date >= ? AND employment_date <= ?", employmentDateRangeArr[0], employmentDateRangeArr[1]).Find(&staffJobs)
			if len(staffJobs) == 0 {
				where += fmt.Sprintf(" AND a.id = -1 ")
			} else {
				where += " AND a.id IN("
				for _, staffJob := range staffJobs {
					where += strconv.Itoa(*staffJob.StaffId) + ","
				}
				where += "0)"
			}
		}
	}
	//选择部门
	if departmentIds != "" {
		var staffDepartments []weChat.WcStaffDepartment
		global.GVA_DB.Where("department_id IN (?)", departmentIds).Find(&staffDepartments)
		fmt.Println("staffDepartments", staffDepartments)
		if len(staffDepartments) == 0 {
			where += fmt.Sprintf(" AND a.id = -1 ")
		} else {
			where += fmt.Sprintf(" AND a.id IN (")
			for _, staffDepartment := range staffDepartments {
				where += strconv.Itoa(*staffDepartment.StaffId) + ","
			}
			where += "0)"
		}
	}

	sql := fmt.Sprintf(`SELECT %s FROM wc_staff AS a 
LEFT JOIN wc_staff_job AS b ON a.id = b.staff_id AND b.deleted_at IS NULL
LEFT JOIN wc_staff_bank AS c ON a.id = c.staff_id AND c.deleted_at IS NULL
LEFT JOIN wc_staff_education AS d ON a.id = d.staff_id AND d.deleted_at IS NULL
LEFT JOIN wc_staff_contact AS e ON a.id = e.staff_id AND e.deleted_at IS NULL
LEFT JOIN wc_staff_agreement AS f ON a.id = f.staff_id AND f.deleted_at IS NULL
WHERE %s ORDER BY a.id ASC,f.id ASC`, fields, where)

	fmt.Println("sql", sql)

	db.Debug().Raw(sql).Scan(&tableMap)

	var rows [][]string
	rows = append(rows, tableTitle)
	for _, table := range tableMap {

		if intVal64, ok := table["rank_type"].(int64); ok {
			rankType := strconv.Itoa(int(intVal64))
			rankList, _, err := GetRankListByRankTypeCommon(rankType)
			if err != nil {
				return nil, "", err
			}
			for _, item := range rankList {
				rankMap[item.ID] = item.Name
			}
		}

		staffId := 0
		fmt.Println("type-value", reflect.TypeOf(table["id"]), table["id"])
		if staffIdUInt32, ok := table["id"].(uint32); ok {
			fmt.Println("staffIdUInt32, ok", staffIdUInt32, ok)
			staffId = int(staffIdUInt32)
		} else {
			fmt.Println("staffIdUInt32, ok", staffIdUInt32, ok)
			return nil, "", err

		}

		var row []string
		for _, column := range columns {
			if column == "id" {
				continue
			}
			transferParams := &TransferParams{
				StaffId:     staffId,
				ConfigInfo:  configInfo,
				Column:      column,
				Value:       table[column],
				StaffMap:    staffMap,
				RankTypeMap: rankTypeMap,
				RankMap:     rankMap,
			}
			if tableColumnValue, ok := transferText(transferParams); ok {
				table[column] = tableColumnValue
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

	return file, template.Name, nil
}

func transferText(t *TransferParams) (text string, ok bool) {
	if t.Value != nil && (t.Column == "rank" || t.Column == "rank_type" || t.Column == "leader") {
		fmt.Println("column-type-value", t.Column, reflect.TypeOf(t.Value), t.Value)
	}

	switch t.Column {
	case "gender":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			text = t.ConfigInfo.StaffGender[intVal]
			return text, true
		} else {
			return "", true
		}
	case "nation":
		if intVal8, ok := t.Value.(int8); ok {
			intVal := int(intVal8)
			text = t.ConfigInfo.Nation[intVal]
			return text, true
		} else {
			return "", true
		}
	case "marriage":
		if intVal8, ok := t.Value.(int8); ok {
			intVal := int(intVal8)
			text = t.ConfigInfo.Marriage[intVal]
			return text, true
		} else {
			return "", true
		}
	case "political_outlook":
		if intVal8, ok := t.Value.(int8); ok {
			intVal := int(intVal8)
			text = t.ConfigInfo.PoliticalOutlook[intVal]
			return text, true
		} else {
			return "", true
		}
	case "household_type":
		if intVal8, ok := t.Value.(int8); ok {
			intVal := int(intVal8)
			text = t.ConfigInfo.HouseholdType[intVal]
			return text, true
		} else {
			return "", true
		}
	case "job_type":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			text = t.ConfigInfo.StaffJobType[intVal]
			return text, true
		} else {
			return "", true
		}
	case "status":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			text = t.ConfigInfo.StaffJobStatus[intVal]
			return text, true
		} else {
			return "", true
		}
	case "try_period":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			text = t.ConfigInfo.StaffJobTryPeriod[intVal]
			return text, true
		} else {
			return "", true
		}
	case "expense_account":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			text = t.ConfigInfo.ExpenseAccount[intVal]
			return text, true
		} else {
			return "", true
		}
	case "education":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			text = t.ConfigInfo.Education[intVal]
			return text, true
		} else {
			return "", true
		}
	case "relationship":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			text = t.ConfigInfo.Relationship[intVal]
			return text, true
		} else {
			return "", true
		}
	case "agreement_type":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			text = t.ConfigInfo.AgreementType[intVal]
			return text, true
		} else {
			return "", true
		}
	case "leader":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			if text, exist := t.StaffMap[intVal]; exist {
				return text, true
			} else {
				return "", true
			}
		} else {
			return "", true
		}
	case "rank_type":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			if text, exist := t.RankTypeMap[intVal]; exist {
				return text, true
			} else {
				return "", true
			}
		} else {
			return "", true
		}
	case "rank":
		if intVal64, ok := t.Value.(int64); ok {
			intVal := int(intVal64)
			if text, exist := t.RankMap[intVal]; exist {
				return text, true
			} else {
				return "", true
			}
		} else {
			return "", true
		}
	case "department":
		text = GetStaffDepartment(t.StaffId)
		return text, true
	case "position":
		text = GetStaffPosition(t.StaffId)
		return text, true
	default:
		if t.Value == nil {
			return "", true
		}
	}
	return "", false
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
		paymentPlaceText, _ := utils.Find(configInfo.PaymentPlace, *item.PaymentPlace)
		newStaff.PaymentPlaceText = paymentPlaceText

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
	paymentPlaceText, _ := utils.Find(configInfo.PaymentPlace, *staff.PaymentPlace)
	newStaff.PaymentPlaceText = paymentPlaceText
	newStaff.WcStaff = staff

	return
}
