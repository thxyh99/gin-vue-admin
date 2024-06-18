package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"strconv"
	"time"
)

type WcStaffAgreementService struct {
}

// CreateWcStaffAgreement 创建合同信息记录
func (wcStaffAgreementService *WcStaffAgreementService) CreateWcStaffAgreement(wcStaffAgreement *weChat.WcStaffAgreement) (err error) {
	if wcStaffAgreement.Times == nil {
		zero := 0
		wcStaffAgreement.Times = &zero
	}
	err = global.GVA_DB.Create(wcStaffAgreement).Error
	return err
}

// DeleteWcStaffAgreement 删除合同信息记录
func (wcStaffAgreementService *WcStaffAgreementService) DeleteWcStaffAgreement(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffAgreement{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffAgreementByIds 批量删除合同信息记录
func (wcStaffAgreementService *WcStaffAgreementService) DeleteWcStaffAgreementByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffAgreement{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffAgreement 更新合同信息记录
func (wcStaffAgreementService *WcStaffAgreementService) UpdateWcStaffAgreement(wcStaffAgreement weChat.WcStaffAgreement) (err error) {
	if wcStaffAgreement.Times == nil {
		zero := 0
		wcStaffAgreement.Times = &zero
	}

	fmt.Println(wcStaffAgreement.ID)
	staffAgreementId := wcStaffAgreement.ID

	err = global.GVA_DB.Save(&wcStaffAgreement).Error
	if err != nil {
		return
	}

	fmt.Println(wcStaffAgreement.ID)

	// 新增固定续签合同需要把之前的续签合同更新为已续签
	if staffAgreementId == 0 && *wcStaffAgreement.Type == 1 {
		var staffAgreement weChat.WcStaffAgreement
		cErr := global.GVA_DB.Debug().Table(staffAgreement.TableName()).Where("staff_id=? AND type=? AND start_day<?", wcStaffAgreement.StaffId, 1, wcStaffAgreement.StartDay).Updates(map[string]interface{}{
			"is_renew":   1,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		}).Error
		if cErr != nil {
			return cErr
		}
	}
	return err
}

// GetWcStaffAgreement 根据员工ID获取合同信息记录
func (wcStaffAgreementService *WcStaffAgreementService) GetWcStaffAgreement(ID string) (newStaffAgreement weChat2.WcStaffAgreementResponse, err error) {
	var staffAgreement weChat.WcStaffAgreement
	err = global.GVA_DB.Where("id = ?", ID).First(&staffAgreement).Error
	if err != nil {
		return
	}

	newStaffAgreement, err = wcStaffAgreementService.AssembleStaffAgreement(staffAgreement)

	return
}

// GetWcStaffAgreementInfoList 分页获取合同信息记录
func (wcStaffAgreementService *WcStaffAgreementService) GetWcStaffAgreementInfoList(info weChatReq.WcStaffAgreementSearch, staffId int) (list []weChat2.WcStaffAgreementResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffAgreement{})
	var wcStaffAgreements []weChat.WcStaffAgreement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if staffId != 0 {
		db = db.Where("staff_id=?", staffId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcStaffAgreements).Error
	if err != nil {
		return
	}

	list, err = wcStaffAgreementService.AssembleStaffAgreementList(wcStaffAgreements)
	return
}

func (wcStaffAgreementService *WcStaffAgreementService) AssembleStaffAgreementList(staffAgreements []weChat.WcStaffAgreement) (newStaffAgreements []weChat2.WcStaffAgreementResponse, err error) {
	var newStaffAgreement weChat2.WcStaffAgreementResponse
	configInfo := config.GetConfigInfo()

	for _, staffAgreement := range staffAgreements {
		newStaffAgreement.WcStaffAgreement = staffAgreement
		typeText, _ := utils.Find(configInfo.AgreementType, *staffAgreement.Type)
		newStaffAgreement.TypeText = typeText
		companyText, _ := utils.Find(configInfo.AgreementCompany, *staffAgreement.Company)
		newStaffAgreement.CompanyText = companyText
		timesText, _ := utils.Find(configInfo.RenewTimes, *staffAgreement.Times)
		newStaffAgreement.TimesText = timesText

		//获取合同期限（月）
		if staffAgreement.StartDay != nil && staffAgreement.EndDay != nil {
			period := utils.CalculateMonthDifference(*staffAgreement.StartDay, *staffAgreement.EndDay)
			newStaffAgreement.Period = strconv.Itoa(period)
		}

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffAgreement.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffAgreementList Err:", err)
			return
		}
		newStaffAgreement.StaffName = staff.Name
		newStaffAgreement.JobNum = staff.JobNum

		newStaffAgreements = append(newStaffAgreements, newStaffAgreement)
	}
	return
}

func (wcStaffAgreementService *WcStaffAgreementService) AssembleStaffAgreement(staffAgreement weChat.WcStaffAgreement) (newStaffAgreement weChat2.WcStaffAgreementResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffAgreement.WcStaffAgreement = staffAgreement
	typeText, _ := utils.Find(configInfo.AgreementType, *staffAgreement.Type)
	newStaffAgreement.TypeText = typeText
	companyText, _ := utils.Find(configInfo.AgreementCompany, *staffAgreement.Company)
	newStaffAgreement.CompanyText = companyText
	timesText, _ := utils.Find(configInfo.RenewTimes, *staffAgreement.Times)
	newStaffAgreement.TimesText = timesText

	//获取合同期限（月）
	if staffAgreement.StartDay != nil && staffAgreement.EndDay != nil {
		period := utils.CalculateMonthDifference(*staffAgreement.StartDay, *staffAgreement.EndDay)
		newStaffAgreement.Period = strconv.Itoa(period)
	}

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffAgreement.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffAgreement Err:", err)
		return
	}

	newStaffAgreement.StaffName = staff.Name
	newStaffAgreement.JobNum = staff.JobNum

	return
}
