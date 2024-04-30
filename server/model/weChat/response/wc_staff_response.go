// 自动生成模板WcStaff
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// 账号信息 结构体  WcStaffResponse
type WcStaffResponse struct {
	weChat.WcStaff
	Position     string `json:"position"`      //职务信息
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

		newStaffs = append(newStaffs, newStaff)
	}

	//for _, item := range newStaffs {
	//	fmt.Println(*item.Gender, item.GenderText, *item.IsLeader, item.IsLeaderText, *item.Status, item.StatusText)
	//}

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

	return
}
