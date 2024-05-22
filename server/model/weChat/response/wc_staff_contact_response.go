package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffContactResponse 工作信息 结构体
type WcStaffContactResponse struct {
	weChat.WcStaffContact
	StaffName        string `json:"staffName"`        //员工名称
	JobNum           string `json:"jobNum"`           //员工工号
	RelationshipText string `json:"relationshipText"` //联系人关系
}
