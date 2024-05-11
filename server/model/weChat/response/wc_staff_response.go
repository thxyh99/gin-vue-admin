package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"sync"
)

// WcStaffResponse 账号信息 结构体
type WcStaffResponse struct {
	weChat.WcStaff
	Position      string `json:"position"`      //职务信息
	PositionIds   []int  `json:"positionIds"`   //职务信息编号
	Department    string `json:"department"`    //部门信息
	DepartmentIds []int  `json:"departmentIds"` //部门信息编号
	GenderText    string `json:"genderText"`    //性别信息
	IsLeaderText  string `json:"isLeaderText" ` //是否领导信息
	StatusText    string `json:"statusText"`    //状态信息
}

func (WcStaffResponse) AssembleStaffList(staffs []weChat.WcStaff) (newStaffs []WcStaffResponse, err error) {
	var newStaff WcStaffResponse
	configInfo := config.GetConfigInfo()

	for _, item := range staffs {
		newStaff.WcStaff = item
		gender, _ := utils.Find(configInfo.StaffGender, *item.Gender)
		newStaff.GenderText = gender
		isLeader, _ := utils.Find(configInfo.StaffIsLeader, *item.IsLeader)
		newStaff.IsLeaderText = isLeader
		status, _ := utils.Find(configInfo.StaffStatus, *item.Status)
		newStaff.StatusText = status

		var wg = sync.WaitGroup{}
		wg.Add(2)
		go func(item weChat.WcStaff) {
			defer wg.Done()

			//拼接员工职位信息
			var sPosition weChat.WcStaffPosition
			var positionIds []int
			var positionText string
			pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.id,p.name").
				Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", item.ID).Rows()
			if err != nil {
				fmt.Println("AssembleStaffList Position1 Err:", err)
				return
			} else {
				for pRows.Next() {
					var id int
					var name string
					err = pRows.Scan(&id, &name)
					if err != nil {
						fmt.Println("AssembleStaffList Position2 Err:", err)
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
			newStaff.PositionIds = positionIds
			newStaff.Position = positionText
		}(item)

		go func(item weChat.WcStaff) {
			defer wg.Done()

			//拼接员工部门信息
			var sDepartment weChat.WcStaffDepartment
			var departmentIds []int
			var departmentText string
			dRows, err := global.GVA_DB.Table(sDepartment.TableName()).Select("department_id").Where("staff_id=?", item.ID).Rows()
			if err != nil {
				fmt.Println("AssembleStaffList Department1 Err:", err)
				return
			} else {
				for dRows.Next() {
					var departmentId int
					err = dRows.Scan(&departmentId)
					if err != nil {
						fmt.Println("AssembleStaffList Department2 Err:", err)
						return
					} else {
						fullName := weChat.GetFullDepartmentById(departmentId)
						if err != nil {
							fmt.Println("AssembleStaffList Department3 Err:", err)
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
			newStaff.DepartmentIds = departmentIds
			newStaff.Department = departmentText
		}(item)

		wg.Wait()
		newStaffs = append(newStaffs, newStaff)
	}
	return
}

func (WcStaffResponse) AssembleStaff(staff weChat.WcStaff) (newStaff WcStaffResponse, err error) {
	configInfo := config.GetConfigInfo()
	gender, _ := utils.Find(configInfo.StaffGender, *staff.Gender)
	newStaff.GenderText = gender
	isLeader, _ := utils.Find(configInfo.StaffIsLeader, *staff.IsLeader)
	newStaff.IsLeaderText = isLeader
	status, _ := utils.Find(configInfo.StaffStatus, *staff.Status)
	newStaff.StatusText = status
	newStaff.WcStaff = staff

	//拼接员工职位信息
	var sPosition weChat.WcStaffPosition
	var positionIds []int
	var positionText string
	pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.id,p.name").
		Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", staff.ID).Rows()
	if err != nil {
		fmt.Println("AssembleStaff Position1 Err:", err)
		return
	} else {
		for pRows.Next() {
			var id int
			var name string
			err = pRows.Scan(&id, &name)
			if err != nil {
				fmt.Println("AssembleStaff Position2 Err:", err)
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
	newStaff.PositionIds = positionIds
	newStaff.Position = positionText

	//拼接员工部门信息
	var sDepartment weChat.WcStaffDepartment
	var departmentIds []int
	var departmentText string
	//dRows, err := global.GVA_DB.Table(sDepartment.TableName()+" as sd").Select("d.id").
	//	Joins("left join wc_department as d on d.id = sd.department_id").Where("sd.staff_id=?", staff.ID).Rows()
	dRows, err := global.GVA_DB.Table(sDepartment.TableName()).Select("department_id").Where("staff_id=?", staff.ID).Rows()
	if err != nil {
		fmt.Println("AssembleStaff Department1 Err:", err)
		return
	} else {
		for dRows.Next() {
			var departmentId int
			err = dRows.Scan(&departmentId)
			if err != nil {
				fmt.Println("AssembleStaff Department2 Err:", err)
				return
			} else {
				departmentIds = append(departmentIds, departmentId)
				fullName := weChat.GetFullDepartmentById(departmentId)
				if err != nil {
					fmt.Println("AssembleStaff Department3 Err:", err)
					return
				}
				if departmentText != "" {
					departmentText += ";" + fullName
				} else {
					departmentText = fullName
				}
			}
		}
	}
	newStaff.DepartmentIds = departmentIds
	newStaff.Department = departmentText

	return
}
