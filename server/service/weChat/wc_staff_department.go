package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
    weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcStaffDepartmentService struct {
}

// CreateWcStaffDepartment 创建员工部门管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffDepartmentService *WcStaffDepartmentService) CreateWcStaffDepartment(wcStaffDepartment *weChat.WcStaffDepartment) (err error) {
	err = global.GVA_DB.Create(wcStaffDepartment).Error
	return err
}

// DeleteWcStaffDepartment 删除员工部门管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffDepartmentService *WcStaffDepartmentService)DeleteWcStaffDepartment(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffDepartment{},"id = ?",ID).Error
	return err
}

// DeleteWcStaffDepartmentByIds 批量删除员工部门管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffDepartmentService *WcStaffDepartmentService)DeleteWcStaffDepartmentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffDepartment{},"id in ?",IDs).Error
	return err
}

// UpdateWcStaffDepartment 更新员工部门管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffDepartmentService *WcStaffDepartmentService)UpdateWcStaffDepartment(wcStaffDepartment weChat.WcStaffDepartment) (err error) {
	err = global.GVA_DB.Save(&wcStaffDepartment).Error
	return err
}

// GetWcStaffDepartment 根据ID获取员工部门管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffDepartmentService *WcStaffDepartmentService)GetWcStaffDepartment(ID string) (wcStaffDepartment weChat.WcStaffDepartment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffDepartment).Error
	return
}

// GetWcStaffDepartmentInfoList 分页获取员工部门管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffDepartmentService *WcStaffDepartmentService)GetWcStaffDepartmentInfoList(info weChatReq.WcStaffDepartmentSearch) (list []weChat.WcStaffDepartment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffDepartment{})
    var wcStaffDepartments []weChat.WcStaffDepartment
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
	
	err = db.Find(&wcStaffDepartments).Error
	return  wcStaffDepartments, total, err
}
