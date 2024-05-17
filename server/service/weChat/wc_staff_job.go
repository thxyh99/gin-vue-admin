package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"time"
)

type WcStaffJobService struct {
}

// CreateWcStaffJob 创建工作信息记录
func (wcStaffJobService *WcStaffJobService) CreateWcStaffJob(wcStaffJobRequest *weChatReq.WcStaffJobRequest) (err error) {
	var wcStaffJob = weChat.WcStaffJob{
		StaffId:        wcStaffJobRequest.StaffId,
		Type:           wcStaffJobRequest.Type,
		Status:         wcStaffJobRequest.Status,
		TryPeriod:      wcStaffJobRequest.TryPeriod,
		RankType:       wcStaffJobRequest.RankType,
		Rank:           wcStaffJobRequest.Rank,
		RankSalary:     wcStaffJobRequest.RankSalary,
		EmploymentDate: wcStaffJobRequest.EmploymentDate,
		FormalDate:     wcStaffJobRequest.FormalDate,
		ExpenseAccount: wcStaffJobRequest.ExpenseAccount,
	}
	err = global.GVA_DB.Create(wcStaffJob).Error

	// 更新员工职位信息
	pSize := len(wcStaffJobRequest.PositionIds)
	if pSize > 0 {
		items := make([]map[string]interface{}, 0, pSize)
		for _, pId := range wcStaffJobRequest.PositionIds {
			var item = make(map[string]interface{})
			item["staff_id"] = wcStaffJobRequest.StaffId
			item["position_id"] = pId
			item["created_at"] = time.Now()
			item["updated_at"] = time.Now()
			items = append(items, item)
			fmt.Println("position item", item)
		}
		cErr := global.GVA_DB.Table(weChat.WcStaffPosition{}.TableName()).CreateInBatches(&items, 1000).Error
		fmt.Println("err3:", cErr)
		if cErr != nil {
			return cErr
		}
	}

	// 更新员工部门信息
	dSize := len(wcStaffJobRequest.DepartmentIds)
	if dSize > 0 {
		items := make([]map[string]interface{}, 0, dSize)
		for _, dId := range wcStaffJobRequest.DepartmentIds {
			var item = make(map[string]interface{})
			item["staff_id"] = wcStaffJobRequest.StaffId
			item["department_id"] = dId
			item["created_at"] = time.Now()
			item["updated_at"] = time.Now()
			items = append(items, item)
			fmt.Println("department item", item)
		}
		cErr := global.GVA_DB.Table(weChat.WcStaffDepartment{}.TableName()).CreateInBatches(&items, 1000).Error
		fmt.Println("err4:", cErr)
		if cErr != nil {
			return cErr
		}
	}

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
func (wcStaffJobService *WcStaffJobService) UpdateWcStaffJob(wcStaffJobRequest *weChatReq.WcStaffJobRequest) (err error) {
	var wcStaffJob = weChat.WcStaffJob{
		StaffId:        wcStaffJobRequest.StaffId,
		Type:           wcStaffJobRequest.Type,
		Status:         wcStaffJobRequest.Status,
		TryPeriod:      wcStaffJobRequest.TryPeriod,
		RankType:       wcStaffJobRequest.RankType,
		Rank:           wcStaffJobRequest.Rank,
		RankSalary:     wcStaffJobRequest.RankSalary,
		EmploymentDate: wcStaffJobRequest.EmploymentDate,
		FormalDate:     wcStaffJobRequest.FormalDate,
		ExpenseAccount: wcStaffJobRequest.ExpenseAccount,
	}
	err = global.GVA_DB.Where("id=?", wcStaffJobRequest.ID).Updates(&wcStaffJob).Error
	if err != nil {
		return err
	}

	// 更新员工职位信息
	pSize := len(wcStaffJobRequest.PositionIds)
	if pSize > 0 {
		global.GVA_DB.Table(weChat.WcStaffPosition{}.TableName()).Where("staff_id=?", wcStaffJobRequest.StaffId).Unscoped().Delete(&weChat.WcStaffPosition{})
		items := make([]map[string]interface{}, 0, pSize)
		for _, pId := range wcStaffJobRequest.PositionIds {
			var item = make(map[string]interface{})
			item["staff_id"] = wcStaffJobRequest.StaffId
			item["position_id"] = pId
			item["created_at"] = time.Now()
			item["updated_at"] = time.Now()
			items = append(items, item)
			fmt.Println("item", item)
		}
		cErr := global.GVA_DB.Table(weChat.WcStaffPosition{}.TableName()).CreateInBatches(&items, 1000).Error
		fmt.Println("err3:", cErr)
		if cErr != nil {
			return cErr
		}
	}

	// 更新员工部门信息
	dSize := len(wcStaffJobRequest.DepartmentIds)
	if dSize > 0 {
		global.GVA_DB.Table(weChat.WcStaffDepartment{}.TableName()).Where("staff_id=?", wcStaffJobRequest.StaffId).Unscoped().Delete(&weChat.WcStaffDepartment{})
		items := make([]map[string]interface{}, 0, dSize)
		for _, dId := range wcStaffJobRequest.DepartmentIds {
			var item = make(map[string]interface{})
			item["staff_id"] = wcStaffJobRequest.StaffId
			item["department_id"] = dId
			item["created_at"] = time.Now()
			item["updated_at"] = time.Now()
			items = append(items, item)
			fmt.Println("item", item)
		}
		cErr := global.GVA_DB.Table(weChat.WcStaffDepartment{}.TableName()).CreateInBatches(&items, 1000).Error
		fmt.Println("err4:", cErr)
		if cErr != nil {
			return cErr
		}
	}

	return
	//err = global.GVA_DB.Save(&wcStaffJob).Error
	//return err
}

// GetWcStaffJob 根据员工ID获取工作信息记录
func (wcStaffJobService *WcStaffJobService) GetWcStaffJob(ID string) (newStaffJob weChat2.WcStaffJobResponse, err error) {
	var staffJob weChat.WcStaffJob
	err = global.GVA_DB.Where("staff_id = ?", ID).First(&staffJob).Error
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
