// 自动生成模板WcStaffJob
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 工作信息 结构体  WcStaffJob
type WcStaffJob struct {
	global.GVA_MODEL
	StaffId        *int       `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:11;" binding:"required"`                                                  //员工ID
	Type           *int       `json:"type" form:"type" gorm:"column:type;comment:员工工作类型(1:全职 2:兼职 3:跟岗实习 4:顶岗实习 5:退休返聘 6:劳务外包 0:其他);size:1;" binding:"required"`               //员工工作类型(1:全职 2:兼职 3:跟岗实习 4:顶岗实习 5:退休返聘 6:劳务外包 0:其他)
	Status         *int       `json:"status" form:"status" gorm:"column:status;comment:员工状态(1:试用 2:正式 3:待离职 4:已离职);size:1;" binding:"required"`                                //员工状态(1:试用 2:正式 3:待离职 4:已离职)
	TryPeriod      *int       `json:"tryPeriod" form:"tryPeriod" gorm:"column:try_period;comment:试用期(1:无试用期 2:2个月 3:6个月 0:其他);size:1;" binding:"required"`                     //试用期(1:无试用期 2:2个月 3:6个月 0:其他)
	EmploymentDate *time.Time `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:入职日期;" binding:"required"`                                     //入职日期
	FormalDate     *time.Time `json:"formalDate" form:"formalDate" gorm:"column:formal_date;comment:转正日期;" binding:"required"`                                                 //转正日期
	Name           string     `json:"name" form:"name" gorm:"column:name;comment:职位名称;size:200;" binding:"required"`                                                           //职位名称
	ExpenseAccount *int       `json:"expenseAccount" form:"expenseAccount" gorm:"column:expense_account;comment:费用科目(1:管理费用 2:研发费用 3:生产费用 4:销售费用);size:1;" binding:"required"` //费用科目(1:管理费用 2:研发费用 3:生产费用 4:销售费用)
}

// TableName 工作信息 WcStaffJob自定义表名 wc_staff_job
func (WcStaffJob) TableName() string {
	return "wc_staff_job"
}
