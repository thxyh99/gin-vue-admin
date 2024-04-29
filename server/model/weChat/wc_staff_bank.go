// 自动生成模板WcStaffBank
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 银行卡信息 结构体  WcStaffBank
type WcStaffBank struct {
	global.GVA_MODEL
	StaffId    *int   `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:11;" binding:"required"`           //员工ID
	CardNumber string `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:银行卡号;size:100;" binding:"required"` //银行卡号
	Bank       string `json:"bank" form:"bank" gorm:"column:bank;comment:开户行;size:200;" binding:"required"`                     //开户行
}

// TableName 银行卡信息 WcStaffBank自定义表名 wc_staff_bank
func (WcStaffBank) TableName() string {
	return "wc_staff_bank"
}
