// 自动生成模板WcStaffJob
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 工作信息 结构体  WcStaffJob
type WcStaffJob struct {
 global.GVA_MODEL 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID(SSO);size:11;" binding:"required"`  //用户ID(SSO) 
      Userid  string `json:"userid" form:"userid" gorm:"column:userid;comment:企微成员UserID;size:200;" binding:"required"`  //企微成员UserID 
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:员工类型(1:全职 2:兼职 3:实习 4:劳务外包 5:无类型);size:1;" binding:"required"`  //员工类型(1:全职 2:兼职 3:实习 4:劳务外包 5:无类型) 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:员工状态(1:试用 2:正式 3:待离职 0:无状态);size:1;" binding:"required"`  //员工状态(1:试用 2:正式 3:待离职 0:无状态) 
      TryPeriod  *int `json:"tryPeriod" form:"tryPeriod" gorm:"column:try_period;comment:试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他);size:1;" binding:"required"`  //试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他) 
      EmploymentDate  *time.Time `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:入职日期;" binding:"required"`  //入职日期 
      FormalDate  *time.Time `json:"formalDate" form:"formalDate" gorm:"column:formal_date;comment:转正日期;" binding:"required"`  //转正日期 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:职位名称;size:200;" binding:"required"`  //职位名称 
}


// TableName 工作信息 WcStaffJob自定义表名 wc_staff_job
func (WcStaffJob) TableName() string {
  return "wc_staff_job"
}

