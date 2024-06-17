// 自动生成模板WcStaffJob
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 工作信息 结构体  WcStaffJob
type WcStaffJob struct {
	global.GVA_MODEL
	StaffId           *int       `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:11;" binding:"required"`                                                  //员工ID
	Type              *int       `json:"type" form:"type" gorm:"column:type;comment:员工类型(1:全职 2:兼职 3:跟岗实习 4:顶岗实习 5:退休返聘 6:劳务外包 0:其他);size:1;" binding:"required"`                 //员工类型(1:全职 2:兼职 3:跟岗实习 4:顶岗实习 5:退休返聘 6:劳务外包 0:其他)
	Status            *int       `json:"status" form:"status" gorm:"column:status;comment:员工状态(1:试用 2:正式 3:待离职 4:已离职);size:1;" binding:"required"`                                //员工状态(1:试用 2:正式 3:待离职 4:已离职)
	TryPeriod         *int       `json:"tryPeriod" form:"tryPeriod" gorm:"column:try_period;comment:试用期(1:无试用期 2:2个月 3:6个月 0:其他);size:1;" binding:"required"`                     //试用期(1:无试用期 2:2个月 3:6个月 0:其他)
	Level             *int       `json:"level" form:"level" gorm:"column:level;comment:层级;size:1;" binding:"required"`                                                            //层级
	IoType            *int       `json:"ioType" form:"ioType" gorm:"column:io_type;comment:内外勤;size:1;" binding:"required"`                                                       //内外勤
	RankType          *int       `json:"rankType" form:"rankType" gorm:"column:rank_type;comment:员工职级类型;size:2;" binding:"required"`                                              //员工职级类型
	Rank              *int       `json:"rank" form:"rank" gorm:"column:rank;comment:员工职级;size:2;" binding:"required"`                                                             //员工职级
	RankSalary        float64    `json:"rankSalary" form:"rankSalary" gorm:"column:rank_salary;comment:等级工资;size:10;" binding:"required"`                                         //等级工资
	EmploymentDate    *time.Time `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:入职容大日期;" binding:"required"`                                   //入职容大日期
	FormalDate        *time.Time `json:"formalDate" form:"formalDate" gorm:"column:formal_date;comment:转正日期;" binding:"required"`                                                 //转正日期
	PresumeFormalDate *time.Time `json:"presumeFormalDate" form:"presumeFormalDate" gorm:"column:presume_formal_date;comment:拟定转正日期;" binding:"required"`                         //拟定转正日期
	LeaveDate         *time.Time `json:"leaveDate" form:"leaveDate" gorm:"column:leave_date;comment:离职日期;"`                                                                       //离职日期
	ExpenseAccount    *int       `json:"expenseAccount" form:"expenseAccount" gorm:"column:expense_account;comment:费用科目(1:管理费用 2:研发费用 3:生产费用 4:销售费用);size:1;" binding:"required"` //费用科目(1:管理费用 2:研发费用 3:生产费用 4:销售费用)
	LeaderId          *int       `json:"leaderId" form:"leaderId" gorm:"column:leader_id;comment:直属领导ID;size:200;"`                                                               //直属领导ID
	PracticeStart     *time.Time `json:"practiceStart" form:"practiceStart" gorm:"column:practice_start;comment:实习协议起始日;"`                                                        //实习协议起始日
	PracticeEnd       *time.Time `json:"practiceEnd" form:"practiceEnd" gorm:"column:practice_end;comment:实习协议到期日;"`                                                              //实习协议到期日
	TrainStart        *time.Time `json:"trainStart" form:"trainStart" gorm:"column:train_start;comment:培训协议起始日;"`                                                                 //培训协议起始日
	TrainEnd          *time.Time `json:"trainEnd" form:"trainEnd" gorm:"column:train_end;comment:培训协议到期日;"`                                                                       //培训协议到期日
	TrainFee          float64    `json:"trainFee" form:"trainFee" gorm:"column:train_fee;comment:培训费用;size:10;" `                                                                 //培训费用
	HealthStart       *time.Time `json:"healthStart" form:"healthStart" gorm:"column:health_start;comment:健康证起始日;"`                                                               //健康证起始日
	HealthEnd         *time.Time `json:"healthEnd" form:"healthEnd" gorm:"column:health_end;comment:健康证到期日;"`                                                                     //健康证到期日
	Remark            string     `json:"remark" form:"remark" gorm:"column:remark;comment:特殊备注;size:200;" `                                                                       //特殊备注
}

// TableName 工作信息 WcStaffJob自定义表名 wc_staff_job
func (WcStaffJob) TableName() string {
	return "wc_staff_job"
}
