package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"sync"
)

// WcStaffJobResponse 工作信息 结构体
type WcStaffJobResponse struct {
	weChat.WcStaffJob
	StaffName          string `json:"staffName"`          //员工名称
	JobNum             string `json:"jobNum"`             //员工工号
	TypeText           string `json:"typeText"`           //员工类型
	StatusText         string `json:"statusText"`         //员工状态
	TryPeriodText      string `json:"tryPeriodText"`      //试用期
	ExpenseAccountText string `json:"expenseAccountText"` //费用科目
	Position           string `json:"position"`           //职务信息
	PositionIds        []int  `json:"positionIds"`        //职务信息编号
	Department         string `json:"department"`         //部门信息
	DepartmentIds      []int  `json:"departmentIds"`      //部门信息编号
	RankTypeText       string `json:"rankTypeText"`       //职级类型
	RankText           string `json:"rankText"`           //职级
}

func (WcStaffJobResponse) AssembleStaffJobList(staffInfos []weChat.WcStaffJob) (newStaffInfos []WcStaffJobResponse, err error) {
	var newStaffJob WcStaffJobResponse
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

		var wg = sync.WaitGroup{}
		wg.Add(5)
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

		wg.Wait()

		newStaffInfos = append(newStaffInfos, newStaffJob)
	}
	return
}

func (WcStaffJobResponse) AssembleStaffJob(staffJob weChat.WcStaffJob) (newStaffJob WcStaffJobResponse, err error) {
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

	var wg = sync.WaitGroup{}
	wg.Add(5)
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

	wg.Wait()

	return
}
