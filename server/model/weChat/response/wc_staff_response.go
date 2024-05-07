// 自动生成模板WcStaff
package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// 账号信息 结构体  WcStaffResponse
type WcStaffResponse struct {
	weChat.WcStaff
	Position     string `json:"position"`      //职务信息
	Department   string `json:"department"`    //部门信息
	GenderText   string `json:"genderText"`    //性别信息
	IsLeaderText string `json:"isLeaderText" ` //是否领导信息
	StatusText   string `json:"statusText"`    //状态信息
}

func (WcStaffResponse) Assemble(staffs []weChat.WcStaff) (newStaffs []WcStaffResponse) {
	var GenderMap = map[int]string{
		0: "未知",
		1: "男",
		2: "女",
	}
	var IsLeaderMaps = map[int]string{
		0: "否",
		1: "是",
	}
	var StatusMaps = map[int]string{
		1: "已激活",
		2: "已禁用",
		4: "未激活",
		5: "退出企业",
	}

	var newStaff WcStaffResponse
	for _, item := range staffs {
		newStaff.WcStaff = item
		newStaff.GenderText = GenderMap[*item.Gender]
		newStaff.IsLeaderText = IsLeaderMaps[*item.IsLeader]
		newStaff.StatusText = StatusMaps[*item.Status]

		//拼接员工职位信息
		var sPosition weChat.WcStaffPosition
		var positionText string
		pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.name").
			Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", item.ID).Rows()
		if err != nil {
			fmt.Println("position1 err:", err)
		} else {
			for pRows.Next() {
				var name string
				err = pRows.Scan(&name)
				if err != nil {
					fmt.Println("position2 err:", err)
				} else {
					if positionText != "" {
						positionText += ";" + name
					} else {
						positionText = name
					}
				}
			}
		}
		newStaff.Position = positionText

		//拼接员工部门信息
		var sDepartment weChat.WcStaffDepartment
		var departmentText string
		dRows, err := global.GVA_DB.Table(sDepartment.TableName()+" as sd").Select("d.name").
			Joins("left join wc_department as d on d.id = sd.department_id").Where("sd.staff_id=?", item.ID).Rows()
		if err != nil {
			fmt.Println("department1 err:", err)
		} else {
			for dRows.Next() {
				var name string
				err = dRows.Scan(&name)
				if err != nil {
					fmt.Println("department2 err:", err)
				} else {
					if departmentText != "" {
						departmentText += ";" + name
					} else {
						departmentText = name
					}
				}
			}
		}
		newStaff.Department = departmentText

		newStaffs = append(newStaffs, newStaff)
	}

	return
}

func (WcStaffResponse) AssembleItem(staff weChat.WcStaff) (wcStaffResponse WcStaffResponse) {
	var GenderMap = map[int]string{
		0: "未知",
		1: "男",
		2: "女",
	}
	var IsLeaderMaps = map[int]string{
		0: "否",
		1: "是",
	}
	var StatusMaps = map[int]string{
		1: "已激活",
		2: "已禁用",
		4: "未激活",
		5: "退出企业",
	}

	wcStaffResponse.GenderText = GenderMap[*staff.Gender]
	wcStaffResponse.IsLeaderText = IsLeaderMaps[*staff.IsLeader]
	wcStaffResponse.StatusText = StatusMaps[*staff.Status]
	wcStaffResponse.WcStaff = staff

	//拼接员工职位信息
	var sPosition weChat.WcStaffPosition
	var positionText string
	pRows, err := global.GVA_DB.Table(sPosition.TableName()+" as sp").Select("p.name").
		Joins("left join wc_position as p on p.id = sp.position_id").Where("sp.staff_id=?", staff.ID).Rows()
	if err != nil {
		fmt.Println("position1 err:", err)
	} else {
		for pRows.Next() {
			var name string
			err = pRows.Scan(&name)
			if err != nil {
				fmt.Println("position2 err:", err)
			} else {
				if positionText != "" {
					positionText += ";" + name
				} else {
					positionText = name
				}
			}
		}
	}
	wcStaffResponse.Position = positionText

	//拼接员工部门信息
	var sDepartment weChat.WcStaffDepartment
	var departmentText string
	dRows, err := global.GVA_DB.Table(sDepartment.TableName()+" as sd").Select("d.name").
		Joins("left join wc_department as d on d.id = sd.department_id").Where("sd.staff_id=?", staff.ID).Rows()
	if err != nil {
		fmt.Println("department1 err:", err)
	} else {
		for dRows.Next() {
			var name string
			err = dRows.Scan(&name)
			if err != nil {
				fmt.Println("department2 err:", err)
			} else {
				if departmentText != "" {
					departmentText += ";" + name
				} else {
					departmentText = name
				}
			}
		}
	}
	wcStaffResponse.Department = departmentText

	return
}
