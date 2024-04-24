package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
    weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcPositionService struct {
}

// CreateWcPosition 创建职位管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcPositionService *WcPositionService) CreateWcPosition(wcPosition *weChat.WcPosition) (err error) {
	err = global.GVA_DB.Create(wcPosition).Error
	return err
}

// DeleteWcPosition 删除职位管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcPositionService *WcPositionService)DeleteWcPosition(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcPosition{},"id = ?",ID).Error
	return err
}

// DeleteWcPositionByIds 批量删除职位管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcPositionService *WcPositionService)DeleteWcPositionByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcPosition{},"id in ?",IDs).Error
	return err
}

// UpdateWcPosition 更新职位管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcPositionService *WcPositionService)UpdateWcPosition(wcPosition weChat.WcPosition) (err error) {
	err = global.GVA_DB.Save(&wcPosition).Error
	return err
}

// GetWcPosition 根据ID获取职位管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcPositionService *WcPositionService)GetWcPosition(ID string) (wcPosition weChat.WcPosition, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcPosition).Error
	return
}

// GetWcPositionInfoList 分页获取职位管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (wcPositionService *WcPositionService)GetWcPositionInfoList(info weChatReq.WcPositionSearch) (list []weChat.WcPosition, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&weChat.WcPosition{})
    var wcPositions []weChat.WcPosition
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
	
	err = db.Find(&wcPositions).Error
	return  wcPositions, total, err
}
