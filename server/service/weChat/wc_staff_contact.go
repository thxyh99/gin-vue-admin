package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
    weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcStaffContactService struct {
}

// CreateWcStaffContact 创建紧急联系人记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffContactService *WcStaffContactService) CreateWcStaffContact(wcStaffContact *weChat.WcStaffContact) (err error) {
	err = global.GVA_DB.Create(wcStaffContact).Error
	return err
}

// DeleteWcStaffContact 删除紧急联系人记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffContactService *WcStaffContactService)DeleteWcStaffContact(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffContact{},"id = ?",ID).Error
	return err
}

// DeleteWcStaffContactByIds 批量删除紧急联系人记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffContactService *WcStaffContactService)DeleteWcStaffContactByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffContact{},"id in ?",IDs).Error
	return err
}

// UpdateWcStaffContact 更新紧急联系人记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffContactService *WcStaffContactService)UpdateWcStaffContact(wcStaffContact weChat.WcStaffContact) (err error) {
	err = global.GVA_DB.Save(&wcStaffContact).Error
	return err
}

// GetWcStaffContact 根据ID获取紧急联系人记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffContactService *WcStaffContactService)GetWcStaffContact(ID string) (wcStaffContact weChat.WcStaffContact, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffContact).Error
	return
}

// GetWcStaffContactInfoList 分页获取紧急联系人记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffContactService *WcStaffContactService)GetWcStaffContactInfoList(info weChatReq.WcStaffContactSearch) (list []weChat.WcStaffContact, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffContact{})
    var wcStaffContacts []weChat.WcStaffContact
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
	
	err = db.Find(&wcStaffContacts).Error
	return  wcStaffContacts, total, err
}
