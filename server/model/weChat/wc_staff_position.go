// 自动生成模板WcStaffPosition
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 员工职位管理 结构体  WcStaffPosition
type WcStaffPosition struct {
	global.GVA_MODEL
	StaffId    *int `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;" binding:"required"`           //员工ID
	PositionId *int `json:"positionId" form:"positionId" gorm:"column:position_id;comment:岗位信息ID;size:6;" binding:"required"` //岗位信息ID
}

// TableName 员工职位管理 WcStaffPosition自定义表名 wc_staff_position
func (WcStaffPosition) TableName() string {
	return "wc_staff_position"
}
