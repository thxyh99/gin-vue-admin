package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
    weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcStaffSocialService struct {
}

// CreateWcStaffSocial 创建社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) CreateWcStaffSocial(wcStaffSocial *weChat.WcStaffSocial) (err error) {
	err = global.GVA_DB.Create(wcStaffSocial).Error
	return err
}

// DeleteWcStaffSocial 删除社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService)DeleteWcStaffSocial(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffSocial{},"id = ?",ID).Error
	return err
}

// DeleteWcStaffSocialByIds 批量删除社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService)DeleteWcStaffSocialByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffSocial{},"id in ?",IDs).Error
	return err
}

// UpdateWcStaffSocial 更新社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService)UpdateWcStaffSocial(wcStaffSocial weChat.WcStaffSocial) (err error) {
	err = global.GVA_DB.Save(&wcStaffSocial).Error
	return err
}

// GetWcStaffSocial 根据ID获取社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService)GetWcStaffSocial(ID string) (wcStaffSocial weChat.WcStaffSocial, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffSocial).Error
	return
}

// GetWcStaffSocialInfoList 分页获取社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService)GetWcStaffSocialInfoList(info weChatReq.WcStaffSocialSearch) (list []weChat.WcStaffSocial, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffSocial{})
    var wcStaffSocials []weChat.WcStaffSocial
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&wcStaffSocials).Error
	return  wcStaffSocials, total, err
}
