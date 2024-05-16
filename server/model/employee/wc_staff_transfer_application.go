// 自动生成模板WcStaffTransferApplication
package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 调动申请 结构体  WcStaffTransferApplication
type WcStaffTransferApplication struct {
	global.GVA_MODEL
	Title              string     `json:"title" form:"title" gorm:"column:title;comment:调动标题;size:100;"`                                                //调动标题
	ApplyTitle         string     `json:"applyTitle" form:"applyTitle" gorm:"column:apply_title;comment:申请方;size:200;" binding:"required"`              //申请方
	EmploymentDate     *time.Time `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:申请日期;" binding:"required"`          //申请日期
	TransferType       string     `json:"transferType" form:"transferType" gorm:"column:transfer_type;type:enum();comment:调动类型;" binding:"required"`    //调动类型
	TransferStaff      *int       `json:"transferStaff" form:"transferStaff" gorm:"column:transfer_staff;comment:被调动人;size:10;" binding:"required"`     //被调动人
	JobDepartment      string     `json:"jobDepartment" form:"jobDepartment" gorm:"column:job_department;type:enum();comment:原部门;" binding:"required"`  //原部门
	JobPosition        string     `json:"jobPosition" form:"jobPosition" gorm:"column:job_position;type:enum();comment:原职位;" binding:"required"`        //原职位
	ToDepartment       string     `json:"toDepartment" form:"toDepartment" gorm:"column:to_department;type:enum();comment:调往部门;" binding:"required"`    //调往部门
	ToPosition         string     `json:"toPosition" form:"toPosition" gorm:"column:to_position;type:enum();comment:调往职位;" binding:"required"`          //调往职位
	TransferResult     string     `json:"transferResult" form:"transferResult" gorm:"column:transfer_result;comment:调动事由;size:255;" binding:"required"` //调动事由
	SourceResult       string     `json:"sourceResult" form:"sourceResult" gorm:"column:source_result;comment:原部门意见;size:255;" binding:"required"`      //原部门意见
	ToResult           string     `json:"toResult" form:"toResult" gorm:"column:to_result;comment:调往部门意见;size:255;" binding:"required"`                 //调往部门意见
	ToDate             *time.Time `json:"toDate" form:"toDate" gorm:"column:to_date;comment:调入时间;" binding:"required"`                                  //调入时间
	InspectionPerion   *int       `json:"inspectionPerion" form:"inspectionPerion" gorm:"column:inspection_perion;comment:考察期;" binding:"required"`     //考察期
	EmployementAttment string     `json:"employementAttment" form:"employementAttment" gorm:"column:employement_attment;comment:附件;size:255;"`          //附件
	SubmitOpinion      string     `json:"submitOpinion" form:"submitOpinion" gorm:"column:submit_opinion;comment:提交意见;size:255;"`                       //提交意见
	Status             string     `json:"status" form:"status" gorm:"column:status;type:enum();comment:状态;" binding:"required"`                         //状态
	CreatedBy          uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 调动申请 WcStaffTransferApplication自定义表名 wc_staff_transfer_application
func (WcStaffTransferApplication) TableName() string {
	return "wc_staff_transfer_application"
}
