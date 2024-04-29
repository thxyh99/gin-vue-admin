// 自动生成模板WcStaffContact
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 紧急联系人 结构体  WcStaffContact
type WcStaffContact struct {
	global.GVA_MODEL
	StaffId      *int   `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:11;" binding:"required"`                                    //员工ID
	Name         string `json:"name" form:"name" gorm:"column:name;comment:紧急联系人姓名;size:200;" binding:"required"`                                          //紧急联系人姓名
	Relationship *int   `json:"relationship" form:"relationship" gorm:"column:relationship;comment:联系人关系(1:父母 2:配偶 3:子女 0:其他);size:1;" binding:"required"` //联系人关系(1:父母 2:配偶 3:子女 0:其他)
	Mobile       string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:联系人电话;size:100;" binding:"required"`                                      //联系人电话
	Address      string `json:"address" form:"address" gorm:"column:address;comment:联系人常住地址;size:255;"`                                                    //联系人常住地址
}

// TableName 紧急联系人 WcStaffContact自定义表名 wc_staff_contact
func (WcStaffContact) TableName() string {
	return "wc_staff_contact"
}
