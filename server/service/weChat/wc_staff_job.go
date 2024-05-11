package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
)

type WcStaffJobService struct {
}

// CreateWcStaffJob 创建工作信息记录
func (wcStaffJobService *WcStaffJobService) CreateWcStaffJob(wcStaffJob *weChat.WcStaffJob) (err error) {
	err = global.GVA_DB.Create(wcStaffJob).Error
	return err
}

// DeleteWcStaffJob 删除工作信息记录
func (wcStaffJobService *WcStaffJobService) DeleteWcStaffJob(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffJob{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffJobByIds 批量删除工作信息记录
func (wcStaffJobService *WcStaffJobService) DeleteWcStaffJobByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffJob{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffJob 更新工作信息记录
func (wcStaffJobService *WcStaffJobService) UpdateWcStaffJob(wcStaffJob weChat.WcStaffJob) (err error) {
	err = global.GVA_DB.Save(&wcStaffJob).Error
	return err
}

// GetWcStaffJob 根据ID获取工作信息记录
func (wcStaffJobService *WcStaffJobService) GetWcStaffJob(ID string) (newStaffJob weChat2.WcStaffJobResponse, err error) {
	var staffJob weChat.WcStaffJob
	err = global.GVA_DB.Where("id = ?", ID).First(&staffJob).Error
	if err != nil {
		return
	}
	newStaffJob, err = weChat2.WcStaffJobResponse{}.AssembleStaffJob(staffJob)
	return
}

// GetWcStaffJobInfoList 分页获取工作信息记录
func (wcStaffJobService *WcStaffJobService) GetWcStaffJobInfoList(info weChatReq.WcStaffJobSearch) (list []weChat2.WcStaffJobResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffJob{})
	var wcStaffJobs []weChat.WcStaffJob
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

	err = db.Find(&wcStaffJobs).Error
	if err != nil {
		return
	}

	list, err = weChat2.WcStaffJobResponse{}.AssembleStaffJobList(wcStaffJobs)
	return
}
