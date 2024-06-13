// 自动生成模板WcStaffAdjustlevelApplication
package employee

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 调级申请 结构体  WcStaffAdjustlevelApplication
type WcStaffAdjustlevelApplication struct {
	global.GVA_MODEL
	Title            string     `json:"title" form:"title" gorm:"column:title;comment:调动标题;size:100;"`                                                            //调动标题
	StaffId          uint       `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;"`                                                      //员工ID
	EffectiveDate    *time.Time `json:"effectiveDate" form:"effectiveDate" gorm:"column:effective_date;comment:生效日期;" binding:"required"`                         //生效日期
	SourceDepartment string     `json:"sourceDepartment" form:"sourceDepartment" gorm:"column:source_department;type:enum(10);comment:调动前部门;" binding:"required"` //调动前部门
	NewDepartment2   string     `json:"newDepartment2" form:"newDepartment2" gorm:"column:new_department2;type:enum(10);comment:调动后部门;" binding:"required"`       //调动后部门
	SourcePosition   string     `json:"sourcePosition" form:"sourcePosition" gorm:"column:source_position;type:enum(10);comment:调动前岗位;" binding:"required"`       //调动前岗位
	NewPosition      string     `json:"newPosition" form:"newPosition" gorm:"column:new_position;type:enum(10);comment:调动前岗位;" binding:"required"`                //调动前岗位
	SourceJoblevel   string     `json:"sourceJoblevel" form:"sourceJoblevel" gorm:"column:source_joblevel;type:enum(10);comment:调动前职级;" binding:"required"`       //调动前职级
	NewJoblevel      string     `json:"newJoblevel" form:"newJoblevel" gorm:"column:new_joblevel;type:enum(10);comment:调动后职级;" binding:"required"`                //调动后职级
	SourceManager    string     `json:"sourceManager" form:"sourceManager" gorm:"column:source_manager;type:enum(10);comment:调动前上级;" binding:"required"`          //调动前上级
	NewManager       string     `json:"newManager" form:"newManager" gorm:"column:new_manager;type:enum(10);comment:调动后上级;" binding:"required"`                   //调动后上级
	Memo             string     `json:"memo" form:"memo" gorm:"column:memo;comment:备注;size:255;" binding:"required"`                                              //备注
	Attachment       string     `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;" binding:"required"`                                     //附件
	OaId             string     `json:"oaId" form:"oaId" gorm:"column:oa_id;comment:OAID;size:50;"`                                                               //OAID
	OaStatus         uint       `json:"oaStatus" form:"oaStatus" gorm:"column:oa_status;comment:OA状态;"`                                                           //OA状态
	CreatedBy        uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 调级申请 WcStaffAdjustlevelApplication自定义表名 wc_staff_Adjustlevel_application
func (WcStaffAdjustlevelApplication) TableName() string {
	return "wc_staff_adjustlevel_application"
}
