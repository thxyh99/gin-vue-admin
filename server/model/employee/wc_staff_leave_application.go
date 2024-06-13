// 自动生成模板WcStaffLeaveApplication
package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 离职申请 结构体  WcStaffLeaveApplication
type WcStaffLeaveApplication struct {
	global.GVA_MODEL
	Title         string         `json:"title" form:"title" gorm:"column:title;comment:离职标题;size:100;"`                                                  //离职标题
	StaffId       uint           `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:19;" binding:"required"`                         //员工ID
	LeaveDate     string         `json:"leaveDate" form:"leaveDate" gorm:"column:leave_date;comment:解除日期;" binding:"required"`                           //解除日期
	JobDepartment string         `json:"jobDepartment" form:"jobDepartment" gorm:"column:job_department;type:enum(10);comment:所属部门;" binding:"required"` //所属部门
	LeaveType     string         `json:"leaveType" form:"leaveType" gorm:"column:leave_type;type:enum();comment:原职位;" binding:"required"`                //原职位
	LeaveResult   string         `json:"leaveResult" form:"leaveResult" gorm:"column:leave_result;comment:事由;size:255;" binding:"required"`              //事由
	Attachment    datatypes.JSON `json:"attachment" form:"attachment" gorm:"column:attachment;comment:申请表;" binding:"required"`                          //申请表
	CheckList     datatypes.JSON `json:"checkList" form:"checkList" gorm:"column:check_list;comment:交接清单;" binding:"required"`                           //交接清单
	IsLeave       *bool          `json:"isLeave" form:"isLeave" gorm:"column:is_leave;comment:是否开具离职证明;" binding:"required"`                             //是否开具离职证明
	IsHome        *bool          `json:"isHome" form:"isHome" gorm:"column:is_home;comment:是否入住公司宿舍;" binding:"required"`                                //是否入住公司宿舍
	DormLocation  string         `json:"dormLocation" form:"dormLocation" gorm:"column:dorm_location;comment:宿舍所在地;size:40;" binding:"required"`         //宿舍所在地
	RoomNum       string         `json:"roomNum" form:"roomNum" gorm:"column:room_num;comment:房间门牌号;size:10;" binding:"required"`                        //房间门牌号
	SubmitOpinion string         `json:"submitOpinion" form:"submitOpinion" gorm:"column:submit_opinion;comment:提交意见;size:255;" binding:"required"`      //提交意见
	OaId          string         `json:"oaId" form:"oaId" gorm:"column:oa_id;comment:OAID;size:50;"`                                                     //OAID
	OaStatus      uint           `json:"oaStatus" form:"oaStatus" gorm:"column:oa_status;comment:OA状态;"`                                                 //OA状态
	CreatedBy     uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint           `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 离职申请 WcStaffLeaveApplication自定义表名 wc_staff_leave_application
func (WcStaffLeaveApplication) TableName() string {
	return "wc_staff_leave_application"
}
