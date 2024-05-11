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
func (wcStaffService *WcStaffService) CreateWcStaff(wcStaffRequest *weChatReq.WcStaffRequest) (err error) {
	zero := 0
	wcStaff := weChat.WcStaff{
		UserId:    &zero,
		Userid:    wcStaffRequest.Userid,
		JobNum:    wcStaffRequest.JobNum,
		Name:      wcStaffRequest.Name,
		Gender:    wcStaffRequest.Gender,
		IsLeader:  wcStaffRequest.IsLeader,
		Mobile:    wcStaffRequest.Mobile,
		Telephone: wcStaffRequest.Telephone,
		Email:     wcStaffRequest.Email,
		Address:   wcStaffRequest.Address,
		BizMail:   wcStaffRequest.BizMail,
		Status:    wcStaffRequest.Status,
	}

	err = global.GVA_DB.Create(&wcStaff).Error
	if err != nil {
		fmt.Println("err1:", err)
		return err
	}

	fmt.Println("staff_id:", wcStaff.ID)

	// 更新员工职位信息
	pSize := len(wcStaffRequest.PositionIds)
	if pSize > 0 {
		items := make([]map[string]interface{}, 0, pSize)
		for _, pId := range wcStaffRequest.PositionIds {
			var item = make(map[string]interface{})
			item["staff_id"] = wcStaff.ID
			item["position_id"] = pId
			item["created_at"] = time.Now()
			item["updated_at"] = time.Now()
			items = append(items, item)
			fmt.Println("position item", item)
		}
		cErr := global.GVA_DB.Table(weChat.WcStaffPosition{}.TableName()).CreateInBatches(&items, 1000).Error
		fmt.Println("err3:", cErr)
		if cErr != nil {
			return cErr
		}
	}

	// 更新员工部门信息
	dSize := len(wcStaffRequest.DepartmentIds)
	if dSize > 0 {
		items := make([]map[string]interface{}, 0, dSize)
		for _, dId := range wcStaffRequest.DepartmentIds {
			var item = make(map[string]interface{})
			item["staff_id"] = wcStaff.ID
			item["department_id"] = dId
			item["created_at"] = time.Now()
			item["updated_at"] = time.Now()
			items = append(items, item)
			fmt.Println("department item", item)
		}
		cErr := global.GVA_DB.Table(weChat.WcStaffDepartment{}.TableName()).CreateInBatches(&items, 1000).Error
		fmt.Println("err4:", cErr)
		if cErr != nil {
			return cErr
		}
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
func (wcStaffService *WcStaffService) UpdateWcStaff(wcStaffRequest *weChatReq.WcStaffRequest) (err error) {
	wcStaff := weChat.WcStaff{
		Userid:    wcStaffRequest.Userid,
		JobNum:    wcStaffRequest.JobNum,
		Name:      wcStaffRequest.Name,
		Gender:    wcStaffRequest.Gender,
		IsLeader:  wcStaffRequest.IsLeader,
		Mobile:    wcStaffRequest.Mobile,
		Telephone: wcStaffRequest.Telephone,
		Email:     wcStaffRequest.Email,
		Address:   wcStaffRequest.Address,
		BizMail:   wcStaffRequest.BizMail,
		Status:    wcStaffRequest.Status,
	}
	err = global.GVA_DB.Where("id=?", wcStaffRequest.ID).Updates(&wcStaff).Error
	if err != nil {
		return err
	}

	fmt.Println("staff_id:", wcStaffRequest.ID)
	fmt.Println("position_ids:", wcStaffRequest.PositionIds)
	fmt.Println("department_ids:", wcStaffRequest.DepartmentIds)

	// 更新员工职位信息
	pSize := len(wcStaffRequest.PositionIds)
	if pSize > 0 {
		global.GVA_DB.Table(weChat.WcStaffPosition{}.TableName()).Where("staff_id=?", wcStaffRequest.ID).Unscoped().Delete(&weChat.WcStaffPosition{})
		items := make([]map[string]interface{}, 0, pSize)
		for _, pId := range wcStaffRequest.PositionIds {
			var item = make(map[string]interface{})
			item["staff_id"] = wcStaffRequest.ID
			item["position_id"] = pId
			item["created_at"] = time.Now()
			item["updated_at"] = time.Now()
			items = append(items, item)
			fmt.Println("item", item)
		}
		cErr := global.GVA_DB.Table(weChat.WcStaffPosition{}.TableName()).CreateInBatches(&items, 1000).Error
		fmt.Println("err3:", cErr)
		if cErr != nil {
			return cErr
		}
	}

	// 更新员工部门信息
	dSize := len(wcStaffRequest.DepartmentIds)
	if dSize > 0 {
		global.GVA_DB.Table(weChat.WcStaffDepartment{}.TableName()).Where("staff_id=?", wcStaffRequest.ID).Unscoped().Delete(&weChat.WcStaffDepartment{})
		items := make([]map[string]interface{}, 0, dSize)
		for _, dId := range wcStaffRequest.DepartmentIds {
			var item = make(map[string]interface{})
			item["staff_id"] = wcStaffRequest.ID
			item["department_id"] = dId
			item["created_at"] = time.Now()
			item["updated_at"] = time.Now()
			items = append(items, item)
			fmt.Println("item", item)
		}
		cErr := global.GVA_DB.Table(weChat.WcStaffDepartment{}.TableName()).CreateInBatches(&items, 1000).Error
		fmt.Println("err4:", cErr)
		if cErr != nil {
			return cErr
		}
	}

	return
}

// GetWcStaff 根据ID获取账号信息记录
func (wcStaffService *WcStaffService) GetWcStaff(ID string) (wcStaffResponse weChat2.WcStaffResponse, err error) {
	var wcStaff weChat.WcStaff
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaff).Error
	if err != nil {
		return
	}
	wcStaffResponse, err = weChat2.WcStaffResponse{}.AssembleStaff(wcStaff)

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

	wcStaffResponse, err = weChat2.WcStaffResponse{}.AssembleStaffList(wcStaffs)

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

	configInfo := config.GetConfigInfo()
	now := time.Now().Format("2006-01-02 15:04:05")

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]
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

		for _, row := range values {
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				err = wcStaffService.checkImportParam(key, value)
				if err != nil {
					return err
				}
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
							wd.DepartmentId = &zero
							wd.Order = &zero
							wd.Parentid = &parentId
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

		for _, row := range values {
			var item = make(map[string]interface{})
			var name, jobNum string
			var lastDepartments, positions []string
			var staffExist, staff weChat.WcStaff

			// 更新员工信息
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				if key == "job_num" {
					jobNum = value
				}
				if key == "name" {
					name = value
				}
				if key == "gender" {
					if gender, ok := utils.FindStringValueKey(configInfo.StaffGender, value); ok {
						value = strconv.Itoa(gender)
					} else {
						return errors.New("性别值异常:" + value)
					}
				}
				if key == "is_leader" {
					if isLeader, ok := utils.FindStringValueKey(configInfo.StaffIsLeader, value); ok {
						value = strconv.Itoa(isLeader)
					} else {
						return errors.New("是否领导值异常:" + value)
					}
				}
				if key == "status" {
					if status, ok := utils.FindStringValueKey(configInfo.StaffStatus, value); ok {
						value = strconv.Itoa(status)
					} else {
						return errors.New("状态值异常:" + value)
					}
				}

				if key != "department" && key != "position" {
					item[key] = value
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

			item["created_at"] = now
			item["updated_at"] = now

			tx.Table(staff.TableName()).Where("name=?", name).First(&staffExist)
			if staffExist.ID == 0 {
				// 新增员工
				cErr := tx.Table(staff.TableName()).Create(&item).Error
				if cErr != nil {
					return cErr
				}
			} else {
				// 更新员工
				cErr := tx.Table(staff.TableName()).Omit("name,created_at").Where("id=?", staffExist.ID).Updates(item).Error
				if cErr != nil {
					return cErr
				}
			}

			tx.Table(staff.TableName()).Where("job_num=?", jobNum).First(&staff)

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
		return nil
	})
}

// checkImportParam 参数校验
func (wcStaffService *WcStaffService) checkImportParam(key, value string) error {
	if key == "name" && value == "" {
		return errors.New("成员名称不能为空")
	}

	if key == "job_num" && value == "" {
		return errors.New("员工工号不能为空")
	}

	configInfo := config.GetConfigInfo()
	if key == "gender" && !utils.InArray(configInfo.StaffGender, value) {
		return errors.New("性别异常:" + value)
	}

	if key == "is_leader" && !utils.InArray(configInfo.StaffIsLeader, value) {
		return errors.New("是否领导异常:" + value)
	}

	if key == "status" && !utils.InArray(configInfo.StaffStatus, value) {
		return errors.New("状态异常:" + value)
	}

	return nil
}
