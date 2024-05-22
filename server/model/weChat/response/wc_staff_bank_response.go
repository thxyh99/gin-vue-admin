package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffBankResponse 工资卡信息 结构体
type WcStaffBankResponse struct {
	weChat.WcStaffBank
	StaffName string `json:"staffName"` //员工名称
	JobNum    string `json:"jobNum"`    //员工工号
}
