package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
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

// 更新OA流程状态
func (wcStaffPassApplicationService *WcStaffPassApplicationService) UpdateWcStaffPassApplicationOaStatus(flowId string, oaStatus uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffPassApplication{}).Where("oa_id = ?", flowId).Update("oa_status", oaStatus).Error; err != nil {
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

// 根据OA ID获取转正申请记录
func (wcStaffPassApplicationService *WcStaffPassApplicationService) GetWcStaffPassApplicationByOA(ID string) (wcStaffPassApplication employee.WcStaffPassApplication, err error) {
	err = global.GVA_DB.Where("oa_id = ?", ID).First(&wcStaffPassApplication).Error
	return
}

// 根据OA 更新员工转正信息
func (wcStaffLeaveApplicationService *WcStaffPassApplicationService) UpdateWcStaffPassApplicationByOA(wcStaffPassApplication employee.WcStaffPassApplication) (err error) {
	updateData := map[string]interface{}{"status": "2", "formal_date": wcStaffPassApplication.PassDate}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&weChat.WcStaffJob{}).Where("staff_id = ?", wcStaffPassApplication.StaffId).Updates(updateData).Error; err != nil {
			return err
		}
		return nil
	})
	return err
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

// 创建OA转正申请流程
func (wcStaffPassApplicationService *WcStaffPassApplicationService) CreateOAPassApplication(wcStaffPassApplication *employee.WcStaffPassApplication) (err error) {
	//读取.ini里面的OA配置
	oaConfig, err := ini.Load("./oa-config.ini")
	if err != nil {
		//失败
		global.GVA_LOG.Error("test", zap.Error(err))
		return err
	}

	//oaWeb := oaConfig.Section("oa-web")
	//oaUrl := oaWeb.Key("base-url").MustString("")
	//addUrl := oaWeb.Key("addOaUrl").MustString("")

	eaConfig := oaConfig.Section("pass_application")

	// 配置OA提交
	eaData := make(map[string]interface{})
	eaData[eaConfig.Key("Title").MustString("")] = wcStaffPassApplication.Title
	eaData[eaConfig.Key("EmploymentDate").MustString("")] = wcStaffPassApplication.EmploymentDate
	//eaData[eaConfig.Key("EmploymentDepartmentID").MustString("")] = wcStaffPassApplication.EmploymentDepartmentId
	eaData[eaConfig.Key("EmploymentDepartmentName").MustString("")] = wcStaffPassApplication.JobDepartment
	eaData[eaConfig.Key("JobPosition").MustString("")] = wcStaffPassApplication.SourcePosition
	eaData[eaConfig.Key("JobLevel").MustString("")] = wcStaffPassApplication.SourceLevel
	eaData[eaConfig.Key("TryPeriod").MustString("")] = wcStaffPassApplication.TryPeriod
	eaData[eaConfig.Key("Education").MustString("")] = wcStaffPassApplication.Education

	// 提交OA流程

	err = global.GVA_DB.Create(wcStaffPassApplication).Error
	return err
}
