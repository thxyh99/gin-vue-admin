package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"gorm.io/gorm"
)

type WcStaffTransferApplicationService struct {
}

// CreateWcStaffTransferApplication 创建调动申请记录
func (wcStaffTransferApplicationService *WcStaffTransferApplicationService) CreateWcStaffTransferApplication(wcStaffTransferApplication *employee.WcStaffTransferApplication) (err error) {
	err = global.GVA_DB.Create(wcStaffTransferApplication).Error
	return err
}

// DeleteWcStaffTransferApplication 删除调动申请记录
func (wcStaffTransferApplicationService *WcStaffTransferApplicationService) DeleteWcStaffTransferApplication(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffTransferApplication{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&employee.WcStaffTransferApplication{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWcStaffTransferApplicationByIds 批量删除调动申请记录
func (wcStaffTransferApplicationService *WcStaffTransferApplicationService) DeleteWcStaffTransferApplicationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffTransferApplication{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&employee.WcStaffTransferApplication{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWcStaffTransferApplication 更新调动申请记录
func (wcStaffTransferApplicationService *WcStaffTransferApplicationService) UpdateWcStaffTransferApplication(wcStaffTransferApplication employee.WcStaffTransferApplication) (err error) {
	err = global.GVA_DB.Save(&wcStaffTransferApplication).Error
	return err
}

// GetWcStaffTransferApplication 根据ID获取调动申请记录
func (wcStaffTransferApplicationService *WcStaffTransferApplicationService) GetWcStaffTransferApplication(ID string) (wcStaffTransferApplication employee.WcStaffTransferApplication, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffTransferApplication).Error
	return
}

// GetWcStaffTransferApplicationInfoList 分页获取调动申请记录
func (wcStaffTransferApplicationService *WcStaffTransferApplicationService) GetWcStaffTransferApplicationInfoList(info employeeReq.WcStaffTransferApplicationSearch) (list []employee.WcStaffTransferApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&employee.WcStaffTransferApplication{})
	var wcStaffTransferApplications []employee.WcStaffTransferApplication
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
	orderMap["transfer_type"] = true
	orderMap["transfer_staff"] = true
	orderMap["source_department"] = true
	orderMap["source_position"] = true
	orderMap["new_department"] = true
	orderMap["new_position"] = true
	orderMap["to_date"] = true
	orderMap["inspection_perion"] = true
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

	err = db.Find(&wcStaffTransferApplications).Error
	return wcStaffTransferApplications, total, err
}
