package weChat

import (
	"encoding/json"
	"errors"
	"fmt"
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
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService) CreateWcStaff(wcStaff *weChat.WcStaff) (err error) {
	err = global.GVA_DB.Create(wcStaff).Error
	return err
}

// DeleteWcStaff 删除账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService) DeleteWcStaff(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaff{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffByIds 批量删除账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService) DeleteWcStaffByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaff{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaff 更新账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService) UpdateWcStaff(wcStaff weChat.WcStaff) (err error) {
	err = global.GVA_DB.Save(&wcStaff).Error
	return err
}

// GetWcStaff 根据ID获取账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService) GetWcStaff(ID string) (wcStaffResponse weChat2.WcStaffResponse, err error) {
	var wcStaff weChat.WcStaff
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaff).Error

	if err != nil {
		return
	}

	wcStaffResponse = weChat2.WcStaffResponse{}.AssembleItem(wcStaff)

	fmt.Println(wcStaffResponse)

	return
}

// GetWcStaffInfoList 分页获取账号信息记录
func (wcStaffService *WcStaffService) GetWcStaffInfoList(info weChatReq.WcStaffSearch) (list []weChat2.WcStaffResponse, total int64, err error) {

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

	var wcStaffsResponse []weChat2.WcStaffResponse
	wcStaffsResponse = weChat2.WcStaffResponse{}.Assemble(wcStaffs)

	return wcStaffsResponse, total, err
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

	GenderMap := map[int]string{
		0: "未知",
		1: "男",
		2: "女",
	}
	IsLeaderMaps := map[int]string{
		0: "否",
		1: "是",
	}
	StatusMaps := map[int]string{
		1: "已激活",
		2: "已禁用",
		4: "未激活",
		5: "退出企业",
	}
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
					gender := utils.GetKeyByValue(GenderMap, value)
					if gender == -1 {
						return errors.New("性别值异常:" + value)
					}
					value = strconv.Itoa(gender)
				}
				if key == "is_leader" {
					isLeader := utils.GetKeyByValue(IsLeaderMaps, value)
					if isLeader == -1 {
						return errors.New("是否领导值异常:" + value)
					}
					value = strconv.Itoa(isLeader)
				}
				if key == "status" {
					status := utils.GetKeyByValue(StatusMaps, value)
					if status == -1 {
						return errors.New("状态值异常:" + value)
					}
					value = strconv.Itoa(status)
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
				//fmt.Println("key:" + key + " value:" + value)
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

	var genderItems = []string{"未知", "男", "女"}
	var isLeaderItems = []string{"否", "是"}
	var statusItems = []string{"已激活", "已禁用", "未激活", "退出企业"}

	if key == "gender" && !utils.InArray(genderItems, value) {
		return errors.New("性别异常:" + value)
	}

	if key == "is_leader" && !utils.InArray(isLeaderItems, value) {
		return errors.New("是否领导异常:" + value)
	}

	if key == "status" && !utils.InArray(statusItems, value) {
		return errors.New("状态异常:" + value)
	}

	return nil
}
