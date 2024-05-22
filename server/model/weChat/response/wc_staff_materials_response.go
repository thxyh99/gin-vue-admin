package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffMaterialsResponseDiscard 工资卡信息 结构体
type WcStaffMaterialsResponseDiscard struct {
	weChat.WcStaffMaterials
	StaffName            string   `json:"staffName"`            //员工名称
	JobNum               string   `json:"jobNum"`               //员工工号
	SkillCertificateList []string `json:"skillCertificateList"` //职称/技能证书
}

type WcStaffMaterialsResponse struct {
	IdCardPortrait         []*WcFileResponse `json:"idCardPortrait"`         //身份证(人像)
	IdCardNational         []*WcFileResponse `json:"idCardNational"`         //身份证(国徽)
	EducationCertificate   []*WcFileResponse `json:"educationCertificate"`   //学历证书
	DegreeCertificate      []*WcFileResponse `json:"degreeCertificate"`      //学位证书
	ResignationCertificate []*WcFileResponse `json:"resignationCertificate"` //前公司离职证明
	OnboardingForm         []*WcFileResponse `json:"onboardingForm"`         //入职补充登记表
	TrialProvide           []*WcFileResponse `json:"trialProvide"`           //试用期管理规定
	PersonalResume         []*WcFileResponse `json:"personalResume"`         //个人简历
	SkillCertificate       []*WcFileResponse `json:"skillCertificate"`       //职称/技能证书
}
