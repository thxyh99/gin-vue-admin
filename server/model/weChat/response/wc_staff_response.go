package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// WcStaffResponse 账号信息 结构体
type WcStaffResponse struct {
	weChat.WcStaff
	GenderText           string `json:"genderText"`            //性别信息
	HouseholdTypeText    string `json:"householdTypeText"`     //户籍类型
	NationText           string `json:"nationText"`            //员工民族
	MarriageText         string `json:"marriageText"`          //员工婚否
	PoliticalOutlookText string `json:"politicalOutlookText" ` //员工政治面貌
}

func (WcStaffResponse) AssembleStaffList(staffs []weChat.WcStaff) (newStaffs []WcStaffResponse, err error) {
	var newStaff WcStaffResponse
	configInfo := config.GetConfigInfo()

	for _, item := range staffs {
		newStaff.WcStaff = item
		gender, _ := utils.Find(configInfo.StaffGender, *item.Gender)
		newStaff.GenderText = gender
		householdTypeText, _ := utils.Find(configInfo.HouseholdType, *item.HouseholdType)
		newStaff.HouseholdTypeText = householdTypeText
		nationText, _ := utils.Find(configInfo.Nation, *item.Nation)
		newStaff.NationText = nationText
		marriageText, _ := utils.Find(configInfo.Marriage, *item.Marriage)
		newStaff.MarriageText = marriageText
		politicalOutlookText, _ := utils.Find(configInfo.PoliticalOutlook, *item.PoliticalOutlook)
		newStaff.PoliticalOutlookText = politicalOutlookText
		//
		//	var wg = sync.WaitGroup{}
		//	wg.Add(2)
		//	go func(item weChat.WcStaff) {
		//		defer wg.Done()
		//
		//		//拼接员工职位信息
		//		var sPosition weChat.WcStaffPosition
		//		var positionIds []int
		//		var positionText string
		//		pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.id,p.name").
		//			Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", item.ID).Rows()
		//		if err != nil {
		//			fmt.Println("AssembleStaffList Position1 Err:", err)
		//			return
		//		} else {
		//			for pRows.Next() {
		//				var id int
		//				var name string
		//				err = pRows.Scan(&id, &name)
		//				if err != nil {
		//					fmt.Println("AssembleStaffList Position2 Err:", err)
		//					return
		//				} else {
		//					positionIds = append(positionIds, id)
		//					if positionText != "" {
		//						positionText += ";" + name
		//					} else {
		//						positionText = name
		//					}
		//				}
		//			}
		//		}
		//		newStaff.PositionIds = positionIds
		//		newStaff.Position = positionText
		//	}(item)
		//
		//	go func(item weChat.WcStaff) {
		//		defer wg.Done()
		//
		//		//拼接员工部门信息
		//		var sDepartment weChat.WcStaffDepartment
		//		var departmentIds []int
		//		var departmentText string
		//		dRows, err := global.GVA_DB.Table(sDepartment.TableName()).Select("department_id").Where("staff_id=?", item.ID).Rows()
		//		if err != nil {
		//			fmt.Println("AssembleStaffList Department1 Err:", err)
		//			return
		//		} else {
		//			for dRows.Next() {
		//				var departmentId int
		//				err = dRows.Scan(&departmentId)
		//				if err != nil {
		//					fmt.Println("AssembleStaffList Department2 Err:", err)
		//					return
		//				} else {
		//					fullName := weChat.GetFullDepartmentById(departmentId)
		//					if err != nil {
		//						fmt.Println("AssembleStaffList Department3 Err:", err)
		//						return
		//					} else {
		//						departmentIds = append(departmentIds, departmentId)
		//						if departmentText != "" {
		//							departmentText += ";" + fullName
		//						} else {
		//							departmentText = fullName
		//						}
		//					}
		//				}
		//			}
		//		}
		//		newStaff.DepartmentIds = departmentIds
		//		newStaff.Department = departmentText
		//	}(item)
		//
		//	wg.Wait()
		newStaffs = append(newStaffs, newStaff)
	}
	return
}

func (WcStaffResponse) AssembleStaff(staff weChat.WcStaff) (newStaff WcStaffResponse, err error) {
	configInfo := config.GetConfigInfo()
	gender, _ := utils.Find(configInfo.StaffGender, *staff.Gender)
	newStaff.GenderText = gender
	householdTypeText, _ := utils.Find(configInfo.HouseholdType, *staff.HouseholdType)
	newStaff.HouseholdTypeText = householdTypeText
	nationText, _ := utils.Find(configInfo.Nation, *staff.Nation)
	newStaff.NationText = nationText
	marriageText, _ := utils.Find(configInfo.Marriage, *staff.Marriage)
	newStaff.MarriageText = marriageText
	politicalOutlookText, _ := utils.Find(configInfo.PoliticalOutlook, *staff.PoliticalOutlook)
	newStaff.PoliticalOutlookText = politicalOutlookText
	newStaff.WcStaff = staff

	////拼接员工职位信息
	//var sPosition weChat.WcStaffPosition
	//var positionIds []int
	//var positionText string
	//pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.id,p.name").
	//	Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", staff.ID).Rows()
	//if err != nil {
	//	fmt.Println("AssembleStaff Position1 Err:", err)
	//	return
	//} else {
	//	for pRows.Next() {
	//		var id int
	//		var name string
	//		err = pRows.Scan(&id, &name)
	//		if err != nil {
	//			fmt.Println("AssembleStaff Position2 Err:", err)
	//			return
	//		} else {
	//			positionIds = append(positionIds, id)
	//			if positionText != "" {
	//				positionText += ";" + name
	//			} else {
	//				positionText = name
	//			}
	//		}
	//	}
	//}
	//newStaff.PositionIds = positionIds
	//newStaff.Position = positionText
	//
	////拼接员工部门信息
	//var sDepartment weChat.WcStaffDepartment
	//var departmentIds []int
	//var departmentText string
	////dRows, err := global.GVA_DB.Table(sDepartment.TableName()+" as sd").Select("d.id").
	////	Joins("left join wc_department as d on d.id = sd.department_id").Where("sd.staff_id=?", staff.ID).Rows()
	//dRows, err := global.GVA_DB.Table(sDepartment.TableName()).Select("department_id").Where("staff_id=?", staff.ID).Rows()
	//if err != nil {
	//	fmt.Println("AssembleStaff Department1 Err:", err)
	//	return
	//} else {
	//	for dRows.Next() {
	//		var departmentId int
	//		err = dRows.Scan(&departmentId)
	//		if err != nil {
	//			fmt.Println("AssembleStaff Department2 Err:", err)
	//			return
	//		} else {
	//			departmentIds = append(departmentIds, departmentId)
	//			fullName := weChat.GetFullDepartmentById(departmentId)
	//			if err != nil {
	//				fmt.Println("AssembleStaff Department3 Err:", err)
	//				return
	//			}
	//			if departmentText != "" {
	//				departmentText += ";" + fullName
	//			} else {
	//				departmentText = fullName
	//			}
	//		}
	//	}
	//}
	//newStaff.DepartmentIds = departmentIds
	//newStaff.Department = departmentText

	return
}
