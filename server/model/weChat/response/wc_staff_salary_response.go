package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffSalaryResponse 工资单信息 结构体
type WcStaffSalaryResponse struct {
	weChat.WcStaffSalary
	StaffName string `json:"staffName"` //员工名称
	JobNum    string `json:"jobNum"`    //员工工号
	TypeText  string `json:"typeText"`  //工资类型
}
