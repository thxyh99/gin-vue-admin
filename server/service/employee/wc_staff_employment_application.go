package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"gorm.io/gorm"
)

type WcStaffEmploymentApplicationService struct {
}

// CreateWcStaffEmploymentApplication 创建入职申请记录
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) CreateWcStaffEmploymentApplication(wcStaffEmploymentApplication *employee.WcStaffEmploymentApplication) (err error) {
	err = global.GVA_DB.Create(wcStaffEmploymentApplication).Error
	return err
}

// DeleteWcStaffEmploymentApplication 删除入职申请记录
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) DeleteWcStaffEmploymentApplication(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffEmploymentApplication{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&employee.WcStaffEmploymentApplication{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWcStaffEmploymentApplicationByIds 批量删除入职申请记录
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) DeleteWcStaffEmploymentApplicationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffEmploymentApplication{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&employee.WcStaffEmploymentApplication{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWcStaffEmploymentApplication 更新入职申请记录
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) UpdateWcStaffEmploymentApplication(wcStaffEmploymentApplication employee.WcStaffEmploymentApplication) (err error) {
	err = global.GVA_DB.Save(&wcStaffEmploymentApplication).Error
	return err
}

// GetWcStaffEmploymentApplication 根据ID获取入职申请记录
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) GetWcStaffEmploymentApplication(ID string) (wcStaffEmploymentApplication employee.WcStaffEmploymentApplication, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffEmploymentApplication).Error
	return
}

// GetWcStaffEmploymentApplicationInfoList 分页获取入职申请记录
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) GetWcStaffEmploymentApplicationInfoList(info employeeReq.WcStaffEmploymentApplicationSearch) (list []employee.WcStaffEmploymentApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&employee.WcStaffEmploymentApplication{})
	var wcStaffEmploymentApplications []employee.WcStaffEmploymentApplication
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
	orderMap["employment_department"] = true
	orderMap["job_position"] = true
	orderMap["payment_place"] = true
	orderMap["gender"] = true
	orderMap["birthday"] = true
	orderMap["native_place"] = true
	orderMap["nationality"] = true
	orderMap["marriage"] = true
	orderMap["political_outlook"] = true
	orderMap["education"] = true
	orderMap["job_level"] = true
	orderMap["try_period"] = true
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

	err = db.Find(&wcStaffEmploymentApplications).Error
	return wcStaffEmploymentApplications, total, err
}
