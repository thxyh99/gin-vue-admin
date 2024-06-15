// 自动生成模板WcStaffAgreement
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 合同信息 结构体  WcStaffAgreement
type WcStaffAgreement struct {
	global.GVA_MODEL
	StaffId  *int       `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;" binding:"required"` //员工ID
	Company  string     `json:"company" form:"company" gorm:"column:company;comment:合同公司;size:200;" binding:"required"` //合同公司
	Type     *int       `json:"type" form:"type" gorm:"column:type;comment:合同类型;size:1;" binding:"required"`            //合同类型(1:固定期限劳动合同2:无固定期限劳动合同3:实习协议4:外包协议5:劳务派遣合同6:返聘协议7:培训协议)
	StartDay *time.Time `json:"startDay" form:"startDay" gorm:"column:start_day;comment:合同起始日;" binding:"required"`     //合同起始日
	EndDay   *time.Time `json:"endDay" form:"endDay" gorm:"column:end_day;comment:合同到期日;" binding:"required"`           //合同到期日
	Times    *int       `json:"times" form:"times" gorm:"column:times;comment:续签次数;size:10;"`                           //续签次数
	FileId   *int       `json:"fileId" form:"fileId" gorm:"column:file_id;comment:合同附件ID;size:10;"`                     //合同附件ID
	Period   *int       `json:"period" form:"period" gorm:"column:period;comment:合同期限（月）;size:10;"`                     //合同期限（月）
}

// TableName 合同信息 WcStaffAgreement自定义表名 wc_staff_agreement
func (WcStaffAgreement) TableName() string {
	return "wc_staff_agreement"
}
