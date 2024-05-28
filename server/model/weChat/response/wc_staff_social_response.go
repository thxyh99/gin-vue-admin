package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffSocialResponse 社保公积金信息 结构体
type WcStaffSocialResponse struct {
	weChat.WcStaffSocial
	StaffName          string `json:"staffName"`          //员工名称
	JobNum             string `json:"jobNum"`             //员工工号
	TypeText           string `json:"typeText"`           //类型
	CredentialTypeText string `json:"credentialTypeText"` //证件类型
}
