package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffMaterialsResponse 工资卡信息 结构体
type WcStaffMaterialsResponse struct {
	weChat.WcStaffMaterials
	StaffName            string   `json:"staffName"`            //员工名称
	JobNum               string   `json:"jobNum"`               //员工工号
	SkillCertificateList []string `json:"skillCertificateList"` //职称/技能证书
}
