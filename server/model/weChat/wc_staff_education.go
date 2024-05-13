// 自动生成模板WcStaffEducation
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 学历信息 结构体  WcStaffEducation
type WcStaffEducation struct {
	global.GVA_MODEL
	StaffId      *int       `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:11;" binding:"required"`                                                        //员工ID
	Education    *int       `json:"education" form:"education" gorm:"column:education;comment:学历(1:大专以下 2:大专 3:专升本/自考本科 4:本科 5:重点本科 6:硕士 7:重点硕士 8:博士);size:1;" binding:"required"` //学历(1:大专以下 2:大专 3:专升本/自考本科 4:本科 5:重点本科 6:硕士 7:重点硕士 8:博士)
	School       string     `json:"school" form:"school" gorm:"column:school;comment:毕业院校;size:200;" binding:"required"`                                                           //毕业院校
	Date         *time.Time `json:"date" form:"date" gorm:"column:date;comment:毕业日期;" binding:"required"`                                                                          //毕业日期
	Major        string     `json:"major" form:"major" gorm:"column:major;comment:专业;size:200;" binding:"required"`                                                                //专业
	Certificate  string     `json:"certificate" form:"certificate" gorm:"column:certificate;comment:职称/技能证书;size:200;"`                                                            //职称/技能证书
	EducationPay *int       `json:"educationPay" form:"educationPay" gorm:"column:education_pay;comment:学历津贴;size:10;" binding:"required"`                                         //学历津贴
	SkillPay     *int       `json:"skillPay" form:"skillPay" gorm:"column:skill_pay;comment:职称技能津贴;size:10;" binding:"required"`                                                   //职称技能津贴
}

// TableName 学历信息 WcStaffEducation自定义表名 wc_staff_education
func (WcStaffEducation) TableName() string {
	return "wc_staff_education"
}
