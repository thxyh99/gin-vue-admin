// 自动生成模板WcStaffLeaveApplication
package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 离职申请 结构体  WcStaffLeaveApplication
type WcStaffLeaveApplication struct {
 global.GVA_MODEL 
      Title  string `json:"title" form:"title" gorm:"column:title;comment:转正标题;size:100;"`  //转正标题 
      StaffName  string `json:"staffName" form:"staffName" gorm:"column:staff_name;comment:姓名;size:200;" binding:"required"`  //姓名 
      LeaveDate  *time.Time `json:"leaveDate" form:"leaveDate" gorm:"column:leave_date;comment:解除日期;" binding:"required"`  //解除日期 
      JobDepartment  string `json:"jobDepartment" form:"jobDepartment" gorm:"column:job_department;type:enum(10);comment:所属部门;" binding:"required"`  //所属部门 
      LeaveType  string `json:"leaveType" form:"leaveType" gorm:"column:leave_type;type:enum();comment:原职位;" binding:"required"`  //原职位 
      LeaveResult  string `json:"leaveResult" form:"leaveResult" gorm:"column:leave_result;comment:事由;size:255;" binding:"required"`  //事由 
      EmployementAttment  string `json:"employementAttment" form:"employementAttment" gorm:"column:employement_attment;comment:附件;size:255;" binding:"required"`  //附件 
      IsLeave  *bool `json:"isLeave" form:"isLeave" gorm:"column:is_leave;comment:是否开具离职证明;" binding:"required"`  //是否开具离职证明 
      IsHome  *bool `json:"isHome" form:"isHome" gorm:"column:is_home;comment:是否入住公司宿舍;" binding:"required"`  //是否入住公司宿舍 
      DormLocation  string `json:"dormLocation" form:"dormLocation" gorm:"column:dorm_location;comment:宿舍所在地;size:40;" binding:"required"`  //宿舍所在地 
      RoomNum  string `json:"roomNum" form:"roomNum" gorm:"column:room_num;comment:房间门牌号;size:10;" binding:"required"`  //房间门牌号 
      SubmitOpinion  string `json:"submitOpinion" form:"submitOpinion" gorm:"column:submit_opinion;comment:提交意见;size:255;"`  //提交意见 
      Status  string `json:"status" form:"status" gorm:"column:status;type:enum();comment:状态;" binding:"required"`  //状态 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 离职申请 WcStaffLeaveApplication自定义表名 wc_staff_leave_application
func (WcStaffLeaveApplication) TableName() string {
  return "wc_staff_leave_application"
}

