package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"gorm.io/gorm"
)

type WcStaffAdjustlevelApplicationService struct {
}

// CreateWcStaffAdjustlevelApplication 创建调级申请记录
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) CreateWcStaffAdjustlevelApplication(wcStaffAdjustlevelApplication *employee.WcStaffAdjustlevelApplication) (err error) {
	err = global.GVA_DB.Create(wcStaffAdjustlevelApplication).Error
	return err
}

// DeleteWcStaffAdjustlevelApplication 删除调级申请记录
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) DeleteWcStaffAdjustlevelApplication(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffAdjustlevelApplication{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&employee.WcStaffAdjustlevelApplication{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// 更新OA流程状态
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) UpdateStaffAdjustlevelApplicationOaStatus(flowId string, oaStatus uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffAdjustlevelApplication{}).Where("oa_id = ?", flowId).Update("oa_status", oaStatus).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWcStaffAdjustlevelApplicationByIds 批量删除调级申请记录
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) DeleteWcStaffAdjustlevelApplicationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffAdjustlevelApplication{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&employee.WcStaffAdjustlevelApplication{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWcStaffAdjustlevelApplication 更新调级申请记录
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) UpdateWcStaffAdjustlevelApplication(wcStaffAdjustlevelApplication employee.WcStaffAdjustlevelApplication) (err error) {
	err = global.GVA_DB.Save(&wcStaffAdjustlevelApplication).Error
	return err
}

// GetWcStaffAdjustlevelApplication 根据ID获取调级申请记录
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) GetWcStaffAdjustlevelApplication(ID string) (wcStaffAdjustlevelApplication employee.WcStaffAdjustlevelApplication, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffAdjustlevelApplication).Error
	return
}

// 根据OAID获取调级申请记录
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) GetWcStaffAdjustlevelApplicationByOA(oaID string) (wcStaffAdjustlevelApplication employee.WcStaffAdjustlevelApplication, err error) {
	err = global.GVA_DB.Where("oa_id = ?", oaID).First(&wcStaffAdjustlevelApplication).Error
	return
}

// 更新员工调级信息
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) UpdateWcStaffAdjustlevelApplicationByOA(wcStaffAdjustlevelApplication employee.WcStaffAdjustlevelApplication) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&weChat.WcPosition{}).Where("staff_id = ?", wcStaffAdjustlevelApplication.StaffId).Update("position_id", wcStaffAdjustlevelApplication.NewPosition).Error; err != nil {
			return err
		}
		if err := tx.Model(&weChat.WcStaffDepartment{}).Where("staff_id = ?", wcStaffAdjustlevelApplication.StaffId).Update("department_id", wcStaffAdjustlevelApplication.NewDepartment).Error; err != nil {
			return err
		}
		if err := tx.Model(&weChat.WcStaffJob{}).Where("staff_id = ?", wcStaffAdjustlevelApplication.StaffId).Update("rank", wcStaffAdjustlevelApplication.NewJoblevel).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// GetWcStaffAdjustlevelApplicationInfoList 分页获取调级申请记录
func (wcStaffAdjustlevelApplicationService *WcStaffAdjustlevelApplicationService) GetWcStaffAdjustlevelApplicationInfoList(info employeeReq.WcStaffAdjustlevelApplicationSearch) (list []employee.WcStaffAdjustlevelApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&employee.WcStaffAdjustlevelApplication{})
	var wcStaffAdjustlevelApplications []employee.WcStaffAdjustlevelApplication
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

	err = db.Find(&wcStaffAdjustlevelApplications).Error
	return wcStaffAdjustlevelApplications, total, err
}
