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

// 更新OA流程状态
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) UpdateWcStaffEmploymentApplicationOaStatus(flowId string, oaStatus uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffEmploymentApplication{}).Where("oa_id = ?", flowId).Update("oa_status", oaStatus).Error; err != nil {
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

func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) GetWcStaffEmploymentApplicationByOA(oaID string) (wcStaffEmploymentApplication employee.WcStaffEmploymentApplication, err error) {
	err = global.GVA_DB.Where("oa_id = ?", oaID).First(&wcStaffEmploymentApplication).Error
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

// 创建OA入职申请流程
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) CreateOAStaffEmploymentApplication(wcStaffEmploymentApplication *employee.WcStaffEmploymentApplication) (err error) {
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

	eaConfig := oaConfig.Section("employment_application")

	// 配置OA提交
	eaData := make(map[string]interface{})
	eaData[eaConfig.Key("Title").MustString("")] = wcStaffEmploymentApplication.Title
	eaData[eaConfig.Key("EmploymentDate").MustString("")] = wcStaffEmploymentApplication.EmploymentDate
	eaData[eaConfig.Key("EmploymentDepartmentID").MustString("")] = wcStaffEmploymentApplication.EmploymentDepartmentId
	eaData[eaConfig.Key("EmploymentDepartmentName").MustString("")] = wcStaffEmploymentApplication.EmploymentDepartmentName
	eaData[eaConfig.Key("JobPosition").MustString("")] = wcStaffEmploymentApplication.JobPosition
	eaData[eaConfig.Key("Gender").MustString("")] = wcStaffEmploymentApplication.Gender
	eaData[eaConfig.Key("Birthday").MustString("")] = wcStaffEmploymentApplication.Birthday
	eaData[eaConfig.Key("NativePlace").MustString("")] = wcStaffEmploymentApplication.NativePlace
	eaData[eaConfig.Key("Nationality").MustString("")] = wcStaffEmploymentApplication.Nationality
	eaData[eaConfig.Key("Height").MustString("")] = wcStaffEmploymentApplication.Height
	eaData[eaConfig.Key("Weight").MustString("")] = wcStaffEmploymentApplication.Weight
	eaData[eaConfig.Key("Marriage").MustString("")] = wcStaffEmploymentApplication.Marriage
	eaData[eaConfig.Key("PoliticalOutlook").MustString("")] = wcStaffEmploymentApplication.PoliticalOutlook
	eaData[eaConfig.Key("Major").MustString("")] = wcStaffEmploymentApplication.Major
	eaData[eaConfig.Key("School").MustString("")] = wcStaffEmploymentApplication.School
	eaData[eaConfig.Key("GraduationDate").MustString("")] = wcStaffEmploymentApplication.GraduationDate
	eaData[eaConfig.Key("Certificate").MustString("")] = wcStaffEmploymentApplication.Certificate
	eaData[eaConfig.Key("IdNumber").MustString("")] = wcStaffEmploymentApplication.IdNumber
	eaData[eaConfig.Key("IdAddress").MustString("")] = wcStaffEmploymentApplication.IdAddress
	eaData[eaConfig.Key("BankAccount").MustString("")] = wcStaffEmploymentApplication.BankAccount
	eaData[eaConfig.Key("BankName").MustString("")] = wcStaffEmploymentApplication.BankName
	eaData[eaConfig.Key("ContactWechat").MustString("")] = wcStaffEmploymentApplication.ContactWechat
	eaData[eaConfig.Key("Mobile").MustString("")] = wcStaffEmploymentApplication.Mobile
	eaData[eaConfig.Key("HomeAddress").MustString("")] = wcStaffEmploymentApplication.HomeAddress
	eaData[eaConfig.Key("LicenseAttachment").MustString("")] = wcStaffEmploymentApplication.LicenseAttachment
	eaData[eaConfig.Key("RelationShip").MustString("")] = wcStaffEmploymentApplication.RelationShip
	eaData[eaConfig.Key("RelationName").MustString("")] = wcStaffEmploymentApplication.RelationName
	eaData[eaConfig.Key("RelationMobile").MustString("")] = wcStaffEmploymentApplication.RelationMobile
	eaData[eaConfig.Key("RelationAddress").MustString("")] = wcStaffEmploymentApplication.RelationAddress
	eaData[eaConfig.Key("IsCeopass").MustString("")] = wcStaffEmploymentApplication.IsCeopass
	eaData[eaConfig.Key("IsBodychecknormal").MustString("")] = wcStaffEmploymentApplication.IsBodychecknormal
	eaData[eaConfig.Key("JobLevel").MustString("")] = wcStaffEmploymentApplication.JobLevel
	eaData[eaConfig.Key("TryPeriod").MustString("")] = wcStaffEmploymentApplication.TryPeriod
	eaData[eaConfig.Key("UrgencyNotification").MustString("")] = wcStaffEmploymentApplication.UrgencyNotification
	eaData[eaConfig.Key("OnboardingOpinion").MustString("")] = wcStaffEmploymentApplication.OnboardingOpinion

	// 提交OA流程

	err = global.GVA_DB.Create(wcStaffEmploymentApplication).Error
	return err
}

// 更新员工入职状态与入职日期
func (wcStaffEmploymentApplicationService *WcStaffEmploymentApplicationService) UpdateStaffEmploymentStatus(wcStaffEmploymentApplication employee.WcStaffEmploymentApplication) (err error) {
	updateData := map[string]interface{}{"status": "1", "employment_date": wcStaffEmploymentApplication.EmploymentDate}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&weChat.WcStaffJob{}).Where("staff_id = ?", wcStaffEmploymentApplication.StaffId).Updates(updateData).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
