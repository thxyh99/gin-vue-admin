package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffAgreementResponse 工作信息 结构体
type WcStaffAgreementResponse struct {
	weChat.WcStaffAgreement
	StaffName   string          `json:"staffName"`   //员工名称
	JobNum      string          `json:"jobNum"`      //员工工号
	TypeText    string          `json:"typeText"`    //合同类型
	CompanyText string          `json:"companyText"` //合同公司
	TimesText   string          `json:"timesText"`   //续签次数
	Period      string          `json:"period"`      //合同期限（月）
	Attachment  *WcFileResponse `json:"attachment"`  //合同附件下载链接
}
