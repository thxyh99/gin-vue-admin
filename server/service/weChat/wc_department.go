package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	"time"
)

type WcDepartmentService struct {
}

// CreateWcDepartment 创建wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) CreateWcDepartment(wcDepartment *weChat.WcDepartment) (err error) {
	err = global.GVA_DB.Create(wcDepartment).Error
	return err
}

// DeleteWcDepartment 删除wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) DeleteWcDepartment(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcDepartment{}, "id = ?", ID).Error
	return err
}

// DeleteWcDepartmentByIds 批量删除wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) DeleteWcDepartmentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcDepartment{}, "id in ?", IDs).Error
	return err
}

// UpdateWcDepartment 更新wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) UpdateWcDepartment(wcDepartment weChat.WcDepartment) (err error) {
	err = global.GVA_DB.Save(&wcDepartment).Error
	return err
}

// GetWcDepartment 根据ID获取wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) GetWcDepartment(ID string) (wcDepartment weChat.WcDepartment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcDepartment).Error
	return
}

// GetWcDepartmentInfoList 分页获取wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) GetWcDepartmentInfoList(info weChatReq.WcDepartmentSearch) (list []weChat.WcDepartment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcDepartment{})
	var wcDepartments []weChat.WcDepartment
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

	err = db.Find(&wcDepartments).Error
	return wcDepartments, total, err
}

// SyncWcDepartment 同步wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) SyncWcDepartment(wcDepartment *weChat.WcDepartment) (err error) {
	var departmentItem weChat.WcDepartment
	err = global.GVA_DB.Model(&weChat.WcDepartment{}).
		Where("department_id = ? AND delete_time IS NULL", wcDepartment.DepartmentId).
		First(&departmentItem).Error

	if err != nil {
		return
	}

	if departmentItem.ID > 0 {
		err = global.GVA_DB.Model(&weChat.WcDepartment{}).
			Where("id = ?", departmentItem.ID).First(&departmentItem).Updates(map[string]interface{}{
			"parentid":   wcDepartment.Parentid,
			"order":      wcDepartment.Order,
			"updated_at": time.Now(),
		}).Error
	} else {
		err = global.GVA_DB.Create(wcDepartment).Error
	}

	return err
}

// GetAllFullDepartmentList 获取部门列表
func (wcDepartmentService *WcDepartmentService) GetAllFullDepartmentList() (list []weChat.FullDepartment, total int64, err error) {
	list, err = weChat.GetAllFullDepartments()

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("list", list)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	total = int64(len(list))

	return
}
