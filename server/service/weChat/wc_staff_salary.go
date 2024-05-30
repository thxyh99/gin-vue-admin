package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
    weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcStaffSalaryService struct {
}

// CreateWcStaffSalary 创建工资列表记录
func (wcStaffSalaryService *WcStaffSalaryService) CreateWcStaffSalary(wcStaffSalary *weChat.WcStaffSalary) (err error) {
	err = global.GVA_DB.Create(wcStaffSalary).Error
	return err
}

// DeleteWcStaffSalary 删除工资列表记录
func (wcStaffSalaryService *WcStaffSalaryService)DeleteWcStaffSalary(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffSalary{},"id = ?",ID).Error
	return err
}

// DeleteWcStaffSalaryByIds 批量删除工资列表记录
func (wcStaffSalaryService *WcStaffSalaryService)DeleteWcStaffSalaryByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffSalary{},"id in ?",IDs).Error
	return err
}

// UpdateWcStaffSalary 更新工资列表记录
func (wcStaffSalaryService *WcStaffSalaryService)UpdateWcStaffSalary(wcStaffSalary weChat.WcStaffSalary) (err error) {
	err = global.GVA_DB.Save(&wcStaffSalary).Error
	return err
}

// GetWcStaffSalary 根据ID获取工资列表记录
func (wcStaffSalaryService *WcStaffSalaryService)GetWcStaffSalary(ID string) (wcStaffSalary weChat.WcStaffSalary, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffSalary).Error
	return
}

// GetWcStaffSalaryInfoList 分页获取工资列表记录
func (wcStaffSalaryService *WcStaffSalaryService)GetWcStaffSalaryInfoList(info weChatReq.WcStaffSalarySearch) (list []weChat.WcStaffSalary, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffSalary{})
    var wcStaffSalarys []weChat.WcStaffSalary
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&wcStaffSalarys).Error
	return  wcStaffSalarys, total, err
}
