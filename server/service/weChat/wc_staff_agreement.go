package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
)

type WcStaffAgreementService struct {
}

// CreateWcStaffAgreement 创建合同信息记录
func (wcStaffAgreementService *WcStaffAgreementService) CreateWcStaffAgreement(wcStaffAgreement *weChat.WcStaffAgreement) (err error) {
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
	err = global.GVA_DB.Save(&wcStaffAgreement).Error
	return err
}

// GetWcStaffAgreement 根据员工ID获取合同信息记录
func (wcStaffAgreementService *WcStaffAgreementService) GetWcStaffAgreement(ID string) (newStaffAgreement weChat2.WcStaffAgreementResponse, err error) {
	var staffAgreement weChat.WcStaffAgreement
	err = global.GVA_DB.Where("staff_id = ?", ID).First(&staffAgreement).Error
	if err != nil {
		return
	}

	newStaffAgreement, err = weChat2.WcStaffAgreementResponse{}.AssembleStaffAgreement(staffAgreement)

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

	list, err = weChat2.WcStaffAgreementResponse{}.AssembleStaffAgreementList(wcStaffAgreements)
	return
}
