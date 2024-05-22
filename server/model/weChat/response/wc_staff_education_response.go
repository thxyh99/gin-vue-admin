package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffEducationResponse 工作信息 结构体
type WcStaffEducationResponse struct {
	weChat.WcStaffEducation
	StaffName     string `json:"staffName"`     //员工名称
	JobNum        string `json:"jobNum"`        //员工工号
	EducationText string `json:"educationText"` //学历
}
