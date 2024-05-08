package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
)

type WcStaffPositionService struct {
}

// CreateWcStaffPosition 创建员工职位管理记录

func (wcStaffPositionService *WcStaffPositionService) CreateWcStaffPosition(wcStaffPosition *weChat.WcStaffPosition) (err error) {
	err = global.GVA_DB.Create(wcStaffPosition).Error
	return err
}

// DeleteWcStaffPosition 删除员工职位管理记录

func (wcStaffPositionService *WcStaffPositionService) DeleteWcStaffPosition(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffPosition{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffPositionByIds 批量删除员工职位管理记录

func (wcStaffPositionService *WcStaffPositionService) DeleteWcStaffPositionByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffPosition{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffPosition 更新员工职位管理记录

func (wcStaffPositionService *WcStaffPositionService) UpdateWcStaffPosition(wcStaffPosition weChat.WcStaffPosition) (err error) {
	err = global.GVA_DB.Save(&wcStaffPosition).Error
	return err
}

// GetWcStaffPosition 根据ID获取员工职位管理记录

func (wcStaffPositionService *WcStaffPositionService) GetWcStaffPosition(ID string) (wcStaffPosition weChat.WcStaffPosition, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcStaffPosition).Error
	return
}

// GetWcStaffPositionInfoList 分页获取员工职位管理记录

func (wcStaffPositionService *WcStaffPositionService) GetWcStaffPositionInfoList(info weChatReq.WcStaffPositionSearch) (list []weChat.WcStaffPosition, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffPosition{})
	var wcStaffPositions []weChat.WcStaffPosition
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

	err = db.Find(&wcStaffPositions).Error
	return wcStaffPositions, total, err
}
