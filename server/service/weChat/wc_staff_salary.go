package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
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
