package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcStaffInfoService struct {
}

// CreateWcStaffInfo 创建个人信息记录

func (wcStaffInfoService *WcStaffInfoService) CreateWcStaffInfo(wcStaffInfo *weChat.WcStaffInfo) (err error) {
	err = global.GVA_DB.Create(wcStaffInfo).Error
	return err
}

// DeleteWcStaffInfo 删除个人信息记录

func (wcStaffInfoService *WcStaffInfoService) DeleteWcStaffInfo(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffInfo{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffInfoByIds 批量删除个人信息记录

func (wcStaffInfoService *WcStaffInfoService) DeleteWcStaffInfoByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffInfo{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffInfo 更新个人信息记录

func (wcStaffInfoService *WcStaffInfoService) UpdateWcStaffInfo(wcStaffInfo weChat.WcStaffInfo) (err error) {
	err = global.GVA_DB.Save(&wcStaffInfo).Error
	return err
}

// GetWcStaffInfo 根据ID获取个人信息记录

func (wcStaffInfoService *WcStaffInfoService) GetWcStaffInfo(ID string) (wcStaffInfo weChat.WcStaffInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffInfo).Error
	return
}

// GetWcStaffInfoInfoList 分页获取个人信息记录

func (wcStaffInfoService *WcStaffInfoService) GetWcStaffInfoInfoList(info weChatReq.WcStaffInfoSearch) (list []weChat.WcStaffInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffInfo{})
	var wcStaffInfos []weChat.WcStaffInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcStaffInfos).Error
	return wcStaffInfos, total, err
}