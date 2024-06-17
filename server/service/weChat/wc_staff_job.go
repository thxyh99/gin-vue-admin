package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"sync"
	"time"
)

type WcStaffJobService struct {
}

// CreateWcStaffJob 创建工作信息记录
func (wcStaffJobService *WcStaffJobService) CreateWcStaffJob(wcStaffJobRequest *weChatReq.WcStaffJobRequest) (err error) {
	if wcStaffJobRequest.LeaderId == nil {
		zero := 0
		wcStaffJobRequest.LeaderId = &zero
	}
	var wcStaffJob = weChat.WcStaffJob{
		StaffId:                   wcStaffJobRequest.StaffId,
		Type:                      wcStaffJobRequest.Type,
		Status:                    wcStaffJobRequest.Status,
		TryPeriod:                 wcStaffJobRequest.TryPeriod,
		RankType:                  wcStaffJobRequest.RankType,
		Rank:                      wcStaffJobRequest.Rank,
		RankSalary:                wcStaffJobRequest.RankSalary,
		EmploymentDate:            wcStaffJobRequest.EmploymentDate,
		EmploymentHeadquarterDate: wcStaffJobRequest.EmploymentHeadquarterDate,
		FormalDate:                wcStaffJobRequest.FormalDate,
		ExpenseAccount:            wcStaffJobRequest.ExpenseAccount,
		LeaderId:                  wcStaffJobRequest.LeaderId,
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
	wcStaffJob := wcStaffJobRequest.WcStaffJob
	// 直属领导非必填可以为空
	if wcStaffJob.LeaderId == nil {
		zero := 0
		wcStaffJob.LeaderId = &zero
	}
	err = global.GVA_DB.Save(&wcStaffJob).Error
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
}

// GetWcStaffJob 根据员工ID获取工作信息记录
func (wcStaffJobService *WcStaffJobService) GetWcStaffJob(ID string) (newStaffJob weChat2.WcStaffJobResponse, err error) {
	var staffJob weChat.WcStaffJob
	err = global.GVA_DB.Where("staff_id = ?", ID).First(&staffJob).Error
	if err != nil {
		return
	}
	newStaffJob, err = wcStaffJobService.AssembleStaffJob(staffJob)
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

	list, err = wcStaffJobService.AssembleStaffJobList(wcStaffJobs)
	return
}

func (wcStaffJobService *WcStaffJobService) AssembleStaffJobList(staffInfos []weChat.WcStaffJob) (newStaffInfos []weChat2.WcStaffJobResponse, err error) {
	var newStaffJob weChat2.WcStaffJobResponse
	configInfo := config.GetConfigInfo()

	for _, staffJob := range staffInfos {
		newStaffJob.WcStaffJob = staffJob
		staffJobTypeText, _ := utils.Find(configInfo.StaffJobType, *staffJob.Type)
		newStaffJob.TypeText = staffJobTypeText
		staffJobStatusText, _ := utils.Find(configInfo.StaffJobStatus, *staffJob.Status)
		newStaffJob.StatusText = staffJobStatusText
		staffJobTryPeriodText, _ := utils.Find(configInfo.StaffJobTryPeriod, *staffJob.TryPeriod)
		newStaffJob.TryPeriodText = staffJobTryPeriodText
		expenseAccountText, _ := utils.Find(configInfo.ExpenseAccount, *staffJob.ExpenseAccount)
		newStaffJob.ExpenseAccountText = expenseAccountText
		levelText, _ := utils.Find(configInfo.Level, *staffJob.Level)
		newStaffJob.LevelText = levelText
		ioTypeText, _ := utils.Find(configInfo.IoType, *staffJob.IoType)
		newStaffJob.IoTypeText = ioTypeText

		var wg = sync.WaitGroup{}
		wg.Add(6)
		go func(item weChat.WcStaffJob) {
			defer wg.Done()

			//拼接员工职位信息
			var sPosition weChat.WcStaffPosition
			var positionIds []int
			var positionText string
			pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.id,p.name").
				Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", item.StaffId).Rows()
			if err != nil {
				fmt.Println("AssembleStaffJobList Position1 Err:", err)
				return
			} else {
				for pRows.Next() {
					var id int
					var name string
					err = pRows.Scan(&id, &name)
					if err != nil {
						fmt.Println("AssembleStaffJobList Position2 Err:", err)
						return
					} else {
						positionIds = append(positionIds, id)
						if positionText != "" {
							positionText += ";" + name
						} else {
							positionText = name
						}
					}
				}
			}
			newStaffJob.PositionIds = positionIds
			newStaffJob.Position = positionText
		}(staffJob)

		go func(item weChat.WcStaffJob) {
			defer wg.Done()

			//拼接员工部门信息
			var sDepartment weChat.WcStaffDepartment
			var departmentIds []int
			var departmentText string
			dRows, err := global.GVA_DB.Table(sDepartment.TableName()).Select("department_id").Where("staff_id=?", item.StaffId).Rows()
			if err != nil {
				fmt.Println("AssembleStaffJobList Department1 Err:", err)
				return
			} else {
				for dRows.Next() {
					var departmentId int
					err = dRows.Scan(&departmentId)
					if err != nil {
						fmt.Println("AssembleStaffJobList Department2 Err:", err)
						return
					} else {
						fullName := weChat.GetFullDepartmentById(departmentId)
						if err != nil {
							fmt.Println("AssembleStaffJobList Department3 Err:", err)
							return
						} else {
							departmentIds = append(departmentIds, departmentId)
							if departmentText != "" {
								departmentText += ";" + fullName
							} else {
								departmentText = fullName
							}
						}
					}
				}
			}
			newStaffJob.DepartmentIds = departmentIds
			newStaffJob.Department = departmentText
		}(staffJob)

		go func(item weChat.WcStaffJob) {
			defer wg.Done()

			var rt weChat.WcRankType
			err = global.GVA_DB.Table(rt.TableName()).Where("id=?", staffJob.RankType).First(&rt).Error
			if err != nil {
				fmt.Println("AssembleStaffJobList Err:", err)
				return
			}
			newStaffJob.RankTypeText = rt.Name
		}(staffJob)

		go func(item weChat.WcStaffJob) {
			defer wg.Done()

			var r weChat.WcRank
			err = global.GVA_DB.Table(r.TableName()).Where("id=?", staffJob.Rank).First(&r).Error
			if err != nil {
				fmt.Println("AssembleStaffJobList Err:", err)
				return
			}
			newStaffJob.RankText = r.Name
		}(staffJob)

		go func(item weChat.WcStaffJob) {
			defer wg.Done()

			//获取员工名称工号
			var staff weChat.WcStaff
			err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffJob.StaffId).First(&staff).Error
			if err != nil {
				fmt.Println("AssembleStaffJobList Err:", err)
				return
			}
			newStaffJob.StaffName = staff.Name
			newStaffJob.JobNum = staff.JobNum
		}(staffJob)

		go func(item weChat.WcStaffJob) {
			defer wg.Done()

			fmt.Println("staffJob.LeaderId", staffJob.LeaderId)

			//获取员工直属领导
			var staffLeader weChat.WcStaff
			err = global.GVA_DB.Table(staffLeader.TableName()).Where("id=?", staffJob.LeaderId).First(&staffLeader).Error
			if err != nil {
				fmt.Println("AssembleStaffJobList Err:", err)
				return
			}
			newStaffJob.Leader = staffLeader.Name
		}(staffJob)

		wg.Wait()

		newStaffInfos = append(newStaffInfos, newStaffJob)
	}
	return
}

func (wcStaffJobService *WcStaffJobService) AssembleStaffJob(staffJob weChat.WcStaffJob) (newStaffJob weChat2.WcStaffJobResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffJob.WcStaffJob = staffJob
	staffJobTypeText, _ := utils.Find(configInfo.StaffJobType, *staffJob.Type)
	newStaffJob.TypeText = staffJobTypeText
	staffJobStatusText, _ := utils.Find(configInfo.StaffJobStatus, *staffJob.Status)
	newStaffJob.StatusText = staffJobStatusText
	staffJobTryPeriodText, _ := utils.Find(configInfo.StaffJobTryPeriod, *staffJob.TryPeriod)
	newStaffJob.TryPeriodText = staffJobTryPeriodText
	expenseAccountText, _ := utils.Find(configInfo.ExpenseAccount, *staffJob.ExpenseAccount)
	newStaffJob.ExpenseAccountText = expenseAccountText
	levelText, _ := utils.Find(configInfo.Level, *staffJob.Level)
	newStaffJob.LevelText = levelText
	ioTypeText, _ := utils.Find(configInfo.IoType, *staffJob.IoType)
	newStaffJob.IoTypeText = ioTypeText

	var wg = sync.WaitGroup{}
	wg.Add(6)
	go func(item weChat.WcStaffJob) {
		defer wg.Done()

		//拼接员工职位信息
		var sPosition weChat.WcStaffPosition
		var positionIds []int
		var positionText string
		pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.id,p.name").
			Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", item.StaffId).Rows()
		if err != nil {
			fmt.Println("AssembleStaffJobList Position1 Err:", err)
			return
		} else {
			for pRows.Next() {
				var id int
				var name string
				err = pRows.Scan(&id, &name)
				if err != nil {
					fmt.Println("AssembleStaffJobList Position2 Err:", err)
					return
				} else {
					positionIds = append(positionIds, id)
					if positionText != "" {
						positionText += ";" + name
					} else {
						positionText = name
					}
				}
			}
		}
		newStaffJob.PositionIds = positionIds
		newStaffJob.Position = positionText
	}(staffJob)

	go func(item weChat.WcStaffJob) {
		defer wg.Done()

		//拼接员工部门信息
		var sDepartment weChat.WcStaffDepartment
		var departmentIds []int
		var departmentText string
		dRows, err := global.GVA_DB.Table(sDepartment.TableName()).Select("department_id").Where("staff_id=?", item.StaffId).Rows()
		if err != nil {
			fmt.Println("AssembleStaffJobList Department1 Err:", err)
			return
		} else {
			for dRows.Next() {
				var departmentId int
				err = dRows.Scan(&departmentId)
				if err != nil {
					fmt.Println("AssembleStaffJobList Department2 Err:", err)
					return
				} else {
					fullName := weChat.GetFullDepartmentById(departmentId)
					if err != nil {
						fmt.Println("AssembleStaffJobList Department3 Err:", err)
						return
					} else {
						departmentIds = append(departmentIds, departmentId)
						if departmentText != "" {
							departmentText += ";" + fullName
						} else {
							departmentText = fullName
						}
					}
				}
			}
		}
		newStaffJob.DepartmentIds = departmentIds
		newStaffJob.Department = departmentText
	}(staffJob)

	go func(item weChat.WcStaffJob) {
		defer wg.Done()

		var rt weChat.WcRankType
		err = global.GVA_DB.Table(rt.TableName()).Where("id=?", staffJob.RankType).First(&rt).Error
		if err != nil {
			fmt.Println("AssembleStaffJobList Err:", err)
			return
		}
		newStaffJob.RankTypeText = rt.Name
	}(staffJob)

	go func(item weChat.WcStaffJob) {
		defer wg.Done()

		var r weChat.WcRank
		err = global.GVA_DB.Table(r.TableName()).Where("id=?", staffJob.Rank).First(&r).Error
		if err != nil {
			fmt.Println("AssembleStaffJobList Err:", err)
			return
		}
		newStaffJob.RankText = r.Name
	}(staffJob)

	go func(item weChat.WcStaffJob) {
		defer wg.Done()

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffJob.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffJobList Err:", err)
			return
		}
		newStaffJob.StaffName = staff.Name
		newStaffJob.JobNum = staff.JobNum
	}(staffJob)

	go func(item weChat.WcStaffJob) {
		defer wg.Done()

		fmt.Println("staffJob.LeaderId", staffJob.LeaderId)

		//获取员工直属领导
		var staffLeader weChat.WcStaff
		err = global.GVA_DB.Table(staffLeader.TableName()).Where("id=?", staffJob.LeaderId).First(&staffLeader).Error
		if err != nil {
			fmt.Println("AssembleStaffJobList Err:", err)
			return
		}
		newStaffJob.Leader = staffLeader.Name
	}(staffJob)

	wg.Wait()

	return
}

// GetStaffPosition 获取员工职位
func GetStaffPosition(id int) (positionText string) {
	//拼接员工职位信息
	var sPosition weChat.WcStaffPosition
	var positionIds []int
	pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.id,p.name").
		Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", id).Rows()
	if err != nil {
		fmt.Println("AssembleStaffJobList Position1 Err:", err)
		return
	} else {
		for pRows.Next() {
			var id int
			var name string
			err = pRows.Scan(&id, &name)
			if err != nil {
				fmt.Println("AssembleStaffJobList Position2 Err:", err)
				return
			} else {
				positionIds = append(positionIds, id)
				if positionText != "" {
					positionText += ";" + name
				} else {
					positionText = name
				}
			}
		}
	}
	return
}

// GetStaffDepartment 获取员工部门信息
func GetStaffDepartment(id int) (departmentText string) {
	//拼接员工部门信息
	var sDepartment weChat.WcStaffDepartment
	var departmentIds []int
	dRows, err := global.GVA_DB.Table(sDepartment.TableName()).Select("department_id").Where("staff_id=?", id).Rows()
	if err != nil {
		fmt.Println("AssembleStaffJobList Department1 Err:", err)
		return
	} else {
		for dRows.Next() {
			var departmentId int
			err = dRows.Scan(&departmentId)
			if err != nil {
				fmt.Println("AssembleStaffJobList Department2 Err:", err)
				return
			} else {
				fullName := weChat.GetFullDepartmentById(departmentId)
				if err != nil {
					fmt.Println("AssembleStaffJobList Department3 Err:", err)
					return
				} else {
					departmentIds = append(departmentIds, departmentId)
					if departmentText != "" {
						departmentText += ";" + fullName
					} else {
						departmentText = fullName
					}
				}
			}
		}
	}
	return
}
