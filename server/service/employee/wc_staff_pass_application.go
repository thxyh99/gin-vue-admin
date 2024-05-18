package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"gorm.io/gorm"
)

type WcStaffPassApplicationService struct {
}

// CreateWcStaffPassApplication 创建转正申请记录
func (wcStaffPassApplicationService *WcStaffPassApplicationService) CreateWcStaffPassApplication(wcStaffPassApplication *employee.WcStaffPassApplication) (err error) {
	err = global.GVA_DB.Create(wcStaffPassApplication).Error
	return err
}

// DeleteWcStaffPassApplication 删除转正申请记录
func (wcStaffPassApplicationService *WcStaffPassApplicationService) DeleteWcStaffPassApplication(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffPassApplication{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&employee.WcStaffPassApplication{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWcStaffPassApplicationByIds 批量删除转正申请记录
func (wcStaffPassApplicationService *WcStaffPassApplicationService) DeleteWcStaffPassApplicationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffPassApplication{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&employee.WcStaffPassApplication{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWcStaffPassApplication 更新转正申请记录
func (wcStaffPassApplicationService *WcStaffPassApplicationService) UpdateWcStaffPassApplication(wcStaffPassApplication employee.WcStaffPassApplication) (err error) {
	err = global.GVA_DB.Save(&wcStaffPassApplication).Error
	return err
}

// GetWcStaffPassApplication 根据ID获取转正申请记录
func (wcStaffPassApplicationService *WcStaffPassApplicationService) GetWcStaffPassApplication(ID string) (wcStaffPassApplication employee.WcStaffPassApplication, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffPassApplication).Error
	return
}

// GetWcStaffPassApplicationInfoList 分页获取转正申请记录
func (wcStaffPassApplicationService *WcStaffPassApplicationService) GetWcStaffPassApplicationInfoList(info employeeReq.WcStaffPassApplicationSearch) (list []employee.WcStaffPassApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&employee.WcStaffPassApplication{})
	var wcStaffPassApplications []employee.WcStaffPassApplication
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
	orderMap["employment_date"] = true
	orderMap["job_department"] = true
	orderMap["source_position"] = true
	orderMap["source_level"] = true
	orderMap["try_period"] = true
	orderMap["pass_date"] = true
	orderMap["new_position"] = true
	orderMap["new_joblevel"] = true
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

	err = db.Find(&wcStaffPassApplications).Error
	return wcStaffPassApplications, total, err
}
