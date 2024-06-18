// 自动生成模板WcStaffLeaveApplication
package employee

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 员工离职申请 结构体  WcStaffLeaveApplication
type WcStaffLeaveApplication struct {
	global.GVA_MODEL
	Title         string     `json:"title" form:"title" gorm:"column:title;comment:离职标题;size:100;" binding:"required"`                         //离职标题
	StaffId       int        `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;" binding:"required"`                   //员工ID
	LeaveDate     *time.Time `json:"leaveDate" form:"leaveDate" gorm:"column:leave_date;comment:解除日期;" binding:"required"`                     //解除日期
	JobDepartment int        `json:"jobDepartment" form:"jobDepartment" gorm:"column:job_department;comment:所属部门;size:10;" binding:"required"` //所属部门
	JobPosition   int        `json:"jobPosition" form:"jobPosition" gorm:"column:job_position;comment:职务;size:10;" binding:"required"`         //职务
	LeaveType     string     `json:"leaveType" form:"leaveType" gorm:"column:leave_type;comment:离职类型;" binding:"required"`                     //离职类型
	LeaveResult   string     `json:"leaveResult" form:"leaveResult" gorm:"column:leave_result;comment:事由;size:255;" binding:"required"`        //事由
	Attachment    string     `json:"attachment" form:"attachment" gorm:"column:attachment;comment:申请表;"`                                       //申请表
	CheckList     string     `json:"checkList" form:"checkList" gorm:"column:check_list;comment:交接清单;"`                                        //交接清单
	IsLeave       *bool      `json:"isLeave" form:"isLeave" gorm:"column:is_leave;comment:是否开具离职证明;"`                                          //是否开具离职证明
	IsHavedorm    *bool      `json:"isHavedorm" form:"isHavedorm" gorm:"column:is_havedorm;comment:是否入住公司宿舍;"`                                 //是否入住公司宿舍
	DormLocation  string     `json:"dormLocation" form:"dormLocation" gorm:"column:dorm_location;comment:宿舍所在地;size:40;"`                      //宿舍所在地
	RoomNum       string     `json:"roomNum" form:"roomNum" gorm:"column:room_num;comment:房间门牌号;size:10;"`                                     //房间门牌号
	SubmitOpinion string     `json:"submitOpinion" form:"submitOpinion" gorm:"column:submit_opinion;comment:提交意见;size:255;"`                   //提交意见
	OaId          string     `json:"oaId" form:"oaId" gorm:"column:oa_id;comment:OAID;size:50;"`                                               //OAID
	OaStatus      int        `json:"oaStatus" form:"oaStatus" gorm:"column:oa_status;comment:OA状态;"`                                           //OA状态
	CreatedBy     int        `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:19;"`                                 //创建者
	UpdatedBy     int        `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:19;"`                                 //更新者
	DeletedBy     int        `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:19;"`                                 //删除者
}

// TableName 员工离职申请 WcStaffLeaveApplication自定义表名 wc_staff_leave_application
func (WcStaffLeaveApplication) TableName() string {
	return "wc_staff_leave_application"
}
