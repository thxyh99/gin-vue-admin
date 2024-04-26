package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
    weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcStaffService struct {
}

// CreateWcStaff 创建账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService) CreateWcStaff(wcStaff *weChat.WcStaff) (err error) {
	err = global.GVA_DB.Create(wcStaff).Error
	return err
}

// DeleteWcStaff 删除账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService)DeleteWcStaff(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaff{},"id = ?",ID).Error
	return err
}

// DeleteWcStaffByIds 批量删除账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService)DeleteWcStaffByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaff{},"id in ?",IDs).Error
	return err
}

// UpdateWcStaff 更新账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService)UpdateWcStaff(wcStaff weChat.WcStaff) (err error) {
	err = global.GVA_DB.Save(&wcStaff).Error
	return err
}

// GetWcStaff 根据ID获取账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService)GetWcStaff(ID string) (wcStaff weChat.WcStaff, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaff).Error
	return
}

// GetWcStaffInfoList 分页获取账号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffService *WcStaffService)GetWcStaffInfoList(info weChatReq.WcStaffSearch) (list []weChat.WcStaff, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&weChat.WcStaff{})
    var wcStaffs []weChat.WcStaff
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
	
	err = db.Find(&wcStaffs).Error
	return  wcStaffs, total, err
}
