package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
    weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcStaffMaterialsService struct {
}

// CreateWcStaffMaterials 创建证件材料记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffMaterialsService *WcStaffMaterialsService) CreateWcStaffMaterials(wcStaffMaterials *weChat.WcStaffMaterials) (err error) {
	err = global.GVA_DB.Create(wcStaffMaterials).Error
	return err
}

// DeleteWcStaffMaterials 删除证件材料记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffMaterialsService *WcStaffMaterialsService)DeleteWcStaffMaterials(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffMaterials{},"id = ?",ID).Error
	return err
}

// DeleteWcStaffMaterialsByIds 批量删除证件材料记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffMaterialsService *WcStaffMaterialsService)DeleteWcStaffMaterialsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffMaterials{},"id in ?",IDs).Error
	return err
}

// UpdateWcStaffMaterials 更新证件材料记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffMaterialsService *WcStaffMaterialsService)UpdateWcStaffMaterials(wcStaffMaterials weChat.WcStaffMaterials) (err error) {
	err = global.GVA_DB.Save(&wcStaffMaterials).Error
	return err
}

// GetWcStaffMaterials 根据ID获取证件材料记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffMaterialsService *WcStaffMaterialsService)GetWcStaffMaterials(ID string) (wcStaffMaterials weChat.WcStaffMaterials, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffMaterials).Error
	return
}

// GetWcStaffMaterialsInfoList 分页获取证件材料记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcStaffMaterialsService *WcStaffMaterialsService)GetWcStaffMaterialsInfoList(info weChatReq.WcStaffMaterialsSearch) (list []weChat.WcStaffMaterials, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffMaterials{})
    var wcStaffMaterialss []weChat.WcStaffMaterials
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
	
	err = db.Find(&wcStaffMaterialss).Error
	return  wcStaffMaterialss, total, err
}
