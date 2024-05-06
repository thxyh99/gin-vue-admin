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
// Author [piexlmax](https://github.com/piexlmax)
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

	fmt.Println("total:", total)

	for _, item := range wcStaffs {
		fmt.Println(item)
	}

	return wcStaffsResponse, total, err
}

// ImportExcel 导入Excel
// Author [piexlmax](https://github.com/piexlmax)
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

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]
		items := make([]map[string]interface{}, 0, len(values))
		for _, row := range values {
			var item = make(map[string]interface{})
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
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
				item[key] = value
				fmt.Println("key:" + key + " value:" + value)
			}

			now := time.Now().Format("2006-01-02 15:04:05")
			item["created_at"] = now
			item["updated_at"] = now
			items = append(items, item)
		}
		cErr := tx.Table(template.TableName).CreateInBatches(&items, 1000).Error
		return cErr
	})
}
