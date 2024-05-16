// 自动生成模板WcStaffPassApplication
package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// 转正申请 结构体  WcStaffPassApplication
type WcStaffPassApplication struct {
 global.GVA_MODEL 
      Title  string `json:"title" form:"title" gorm:"column:title;comment:转正标题;size:100;"`  //转正标题 
      StaffName  string `json:"staffName" form:"staffName" gorm:"column:staff_name;comment:姓名;size:200;" binding:"required"`  //姓名 
      EmploymentDate  *time.Time `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:入职日期;" binding:"required"`  //入职日期 
      JobDepartment  string `json:"jobDepartment" form:"jobDepartment" gorm:"column:job_department;type:enum(10);comment:所属部门;" binding:"required"`  //所属部门 
      JobPosition  string `json:"jobPosition" form:"jobPosition" gorm:"column:job_position;type:enum();comment:原职位;" binding:"required"`  //原职位 
      JobLevel  string `json:"jobLevel" form:"jobLevel" gorm:"column:job_level;type:enum();comment:原职级;" binding:"required"`  //原职级 
      TryPeriod  string `json:"tryPeriod" form:"tryPeriod" gorm:"column:try_period;type:enum();comment:试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他);" binding:"required"`  //试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他) 
      IsMananger  *bool `json:"isMananger" form:"isMananger" gorm:"column:is_mananger;comment:是否是部门负责人;" binding:"required"`  //是否是部门负责人 
      EmployementAttment  datatypes.JSON `json:"employementAttment" form:"employementAttment" gorm:"column:employement_attment;comment:附件;size:255;" binding:"required"`  //附件 
      SelfOpinion  string `json:"selfOpinion" form:"selfOpinion" gorm:"column:self_opinion;comment:个人自我鉴定;size:255;" binding:"required"`  //个人自我鉴定 
      TryOpinion  string `json:"tryOpinion" form:"tryOpinion" gorm:"column:try_opinion;comment:用人部门意见;size:255;" binding:"required"`  //用人部门意见 
      PassResponsibilities  string `json:"passResponsibilities" form:"passResponsibilities" gorm:"column:pass_responsibilities;comment:转正后工作职责;size:255;" binding:"required"`  //转正后工作职责 
      PassDate  *time.Time `json:"passDate" form:"passDate" gorm:"column:pass_date;comment:转正时间;" binding:"required"`  //转正时间 
      PassPosition  string `json:"passPosition" form:"passPosition" gorm:"column:pass_position;type:enum();comment:转正后职位;" binding:"required"`  //转正后职位 
      PassJoblevel  string `json:"passJoblevel" form:"passJoblevel" gorm:"column:pass_joblevel;type:enum();comment:转正后职级;" binding:"required"`  //转正后职级 
      SubmitOpinion  string `json:"submitOpinion" form:"submitOpinion" gorm:"column:submit_opinion;comment:提交意见;size:255;"`  //提交意见 
      Status  string `json:"status" form:"status" gorm:"column:status;type:enum();comment:状态;"`  //状态 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 转正申请 WcStaffPassApplication自定义表名 wc_staff_pass_application
func (WcStaffPassApplication) TableName() string {
  return "wc_staff_pass_application"
}

