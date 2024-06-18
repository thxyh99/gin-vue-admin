package employee

import (
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	wcChatService "github.com/flipped-aurora/gin-vue-admin/server/service/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type WcStaffLeaveApplicationService struct {
}

// CreateWcStaffLeaveApplication 创建离职申请记录
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) CreateWcStaffLeaveApplication(wcStaffLeaveApplication *employee.WcStaffLeaveApplication) (err error) {
	err = global.GVA_DB.Create(wcStaffLeaveApplication).Error
	return err
}

// DeleteWcStaffLeaveApplication 删除离职申请记录
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) DeleteWcStaffLeaveApplication(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffLeaveApplication{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&employee.WcStaffLeaveApplication{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// 更新流程状态
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) UpdateWcStaffLeaveApplicationOaStatus(flowId string, oaStatus uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffLeaveApplication{}).Where("oa_id = ?", flowId).Update("oa_status", oaStatus).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// 更新员工离职状态
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) UpdateWcStaffLeaveApplicationByOa(wcStaffLeaveApplication employee.WcStaffLeaveApplication) (err error) {
	updateData := map[string]interface{}{"status": "3", "leave_date": wcStaffLeaveApplication.LeaveDate}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&weChat.WcStaffJob{}).Where("staff_id = ?", wcStaffLeaveApplication).Updates(updateData).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWcStaffLeaveApplicationByIds 批量删除离职申请记录
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) DeleteWcStaffLeaveApplicationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&employee.WcStaffLeaveApplication{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&employee.WcStaffLeaveApplication{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWcStaffLeaveApplication 更新离职申请记录
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) UpdateWcStaffLeaveApplication(wcStaffLeaveApplication employee.WcStaffLeaveApplication) (err error) {
	err = global.GVA_DB.Save(&wcStaffLeaveApplication).Error
	return err
}

// GetWcStaffLeaveApplication 根据ID获取离职申请记录
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) GetWcStaffLeaveApplication(ID string) (wcStaffLeaveApplication employee.WcStaffLeaveApplication, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffLeaveApplication).Error
	return
}

// 根据oaid获取离职申请记录
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) GetWcStaffLeaveApplicationByOA(ID string) (wcStaffLeaveApplication employee.WcStaffLeaveApplication, err error) {
	err = global.GVA_DB.Where("oa_id = ?", ID).First(&wcStaffLeaveApplication).Error
	return
}

// GetWcStaffLeaveApplicationInfoList 分页获取离职申请记录
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) GetWcStaffLeaveApplicationInfoList(info employeeReq.WcStaffLeaveApplicationSearch) (list []employee.WcStaffLeaveApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&employee.WcStaffLeaveApplication{})
	var wcStaffLeaveApplications []employee.WcStaffLeaveApplication
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
	orderMap["leave_date"] = true
	orderMap["job_department"] = true
	orderMap["leave_type"] = true
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

	err = db.Find(&wcStaffLeaveApplications).Error
	return wcStaffLeaveApplications, total, err
}

// 创建OA离职申请流程
func (wcStaffLeaveApplicationService *WcStaffLeaveApplicationService) CreateOAStaffLeaveApplication(wcStaffLeaveApplication *employee.WcStaffLeaveApplication) (err error) {
	oaStaffLeaveApplication := utils.OAStaffLeaveApplication{}
	var staffServer wcChatService.WcStaffService
	wcStaff, err := staffServer.GetWcStaff(strconv.Itoa(wcStaffLeaveApplication.StaffId))

	oaStaffLeaveApplication.DocSubject = "离职申请-" + wcStaff.Name
	oaStaffLeaveApplication.DocCreator = "liuyongbo"
	oaStaffLeaveApplication.DocStatus = "10"

	oaStaffLeaveApplication.StaffName = wcStaff.Name
	oaStaffLeaveApplication.LeaveDate = wcStaffLeaveApplication.LeaveDate
	oaStaffLeaveApplication.LeaveType = wcStaffLeaveApplication.LeaveType
	oaStaffLeaveApplication.LeaveResult = wcStaffLeaveApplication.LeaveResult

	landrayOa := utils.LandrayOa{}
	oaId, err := landrayOa.CreateOALeaveApplication(oaStaffLeaveApplication)
	if err != nil || strings.Index("code", oaId) > 0 {
		global.GVA_LOG.Error("创建员工OA离职申请失败", zap.Error(err))
		return err
	}
	return nil
}
