package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"gorm.io/gorm"
)

type WcStaffAjustlevelApplicationService struct {
}

// CreateWcStaffAjustlevelApplication 创建调级申请记录
func (wcStaffAjustlevelApplicationService *WcStaffAjustlevelApplicationService) CreateWcStaffAjustlevelApplication(wcStaffAjustlevelApplication *employee.WcStaffAjustlevelApplication) (err error) {
	err = global.GVA_DB.Create(wcStaffAjustlevelApplication).Error
	return err
}

// DeleteWcStaffAjustlevelApplication 删除调级申请记录
func (wcStaffAjustlevelApplicationService *WcStaffAjustlevelApplicationService) DeleteWcStaffAjustlevelApplication(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffAjustlevelApplication{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&employee.WcStaffAjustlevelApplication{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWcStaffAjustlevelApplicationByIds 批量删除调级申请记录
func (wcStaffAjustlevelApplicationService *WcStaffAjustlevelApplicationService) DeleteWcStaffAjustlevelApplicationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffAjustlevelApplication{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&employee.WcStaffAjustlevelApplication{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWcStaffAjustlevelApplication 更新调级申请记录
func (wcStaffAjustlevelApplicationService *WcStaffAjustlevelApplicationService) UpdateWcStaffAjustlevelApplication(wcStaffAjustlevelApplication employee.WcStaffAjustlevelApplication) (err error) {
	err = global.GVA_DB.Save(&wcStaffAjustlevelApplication).Error
	return err
}

// GetWcStaffAjustlevelApplication 根据ID获取调级申请记录
func (wcStaffAjustlevelApplicationService *WcStaffAjustlevelApplicationService) GetWcStaffAjustlevelApplication(ID string) (wcStaffAjustlevelApplication employee.WcStaffAjustlevelApplication, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffAjustlevelApplication).Error
	return
}

// GetWcStaffAjustlevelApplicationInfoList 分页获取调级申请记录
func (wcStaffAjustlevelApplicationService *WcStaffAjustlevelApplicationService) GetWcStaffAjustlevelApplicationInfoList(info employeeReq.WcStaffAjustlevelApplicationSearch) (list []employee.WcStaffAjustlevelApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&employee.WcStaffAjustlevelApplication{})
	var wcStaffAjustlevelApplications []employee.WcStaffAjustlevelApplication
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["effective_date"] = true
	orderMap["source_department"] = true
	orderMap["new_department2"] = true
	orderMap["source_position"] = true
	orderMap["new_position"] = true
	orderMap["source_joblevel"] = true
	orderMap["new_joblevel"] = true
	orderMap["source_manager"] = true
	orderMap["new_manager"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcStaffAjustlevelApplications).Error
	return wcStaffAjustlevelApplications, total, err
}
