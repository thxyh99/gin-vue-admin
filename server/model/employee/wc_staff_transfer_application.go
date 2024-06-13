// 自动生成模板WcStaffTransferApplication
package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 调动申请 结构体  WcStaffTransferApplication
type WcStaffTransferApplication struct {
	global.GVA_MODEL
	Title            string `json:"title" form:"title" gorm:"column:title;comment:调动标题;size:100;"`                                                          //调动标题
	ApplyTitle       string `json:"applyTitle" form:"applyTitle" gorm:"column:apply_title;comment:申请方;size:200;" binding:"required"`                        //申请方
	EmploymentDate   string `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:申请日期;" binding:"required"`                    //申请日期
	TransferType     string `json:"transferType" form:"transferType" gorm:"column:transfer_type;type:enum();comment:调动类型;" binding:"required"`              //调动类型
	TransferStaff    string `json:"transferStaff" form:"transferStaff" gorm:"column:transfer_staff;type:enum(10);comment:被调动人;" binding:"required"`         //被调动人
	SourceDepartment string `json:"sourceDepartment" form:"sourceDepartment" gorm:"column:source_department;type:enum(10);comment:原部门;" binding:"required"` //原部门
	SourcePosition   string `json:"sourcePosition" form:"sourcePosition" gorm:"column:source_position;type:enum(10);comment:原职位;" binding:"required"`       //原职位
	NewDepartment    string `json:"newDepartment" form:"newDepartment" gorm:"column:new_department;type:enum(10);comment:调往部门;" binding:"required"`         //调往部门
	NewPosition      string `json:"newPosition" form:"newPosition" gorm:"column:new_position;type:enum(10);comment:调往职位;" binding:"required"`               //调往职位
	TransferResult   string `json:"transferResult" form:"transferResult" gorm:"column:transfer_result;comment:调动事由;size:255;" binding:"required"`           //调动事由
	SourceResult     string `json:"sourceResult" form:"sourceResult" gorm:"column:source_result;comment:原部门意见;size:255;" binding:"required"`                //原部门意见
	ToResult         string `json:"toResult" form:"toResult" gorm:"column:to_result;comment:调往部门意见;size:255;" binding:"required"`                           //调往部门意见
	ToDate           string `json:"toDate" form:"toDate" gorm:"column:to_date;comment:调入时间;" binding:"required"`                                            //调入时间
	InspectionPerion string `json:"inspectionPerion" form:"inspectionPerion" gorm:"column:inspection_perion;type:enum();comment:考察期;" binding:"required"`   //考察期
	Attachment       string `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;" binding:"required"`                                   //附件
	SubmitOpinion    string `json:"submitOpinion" form:"submitOpinion" gorm:"column:submit_opinion;comment:提交意见;size:255;" binding:"required"`              //提交意见
	OaId             string `json:"oaId" form:"oaId" gorm:"column:oa_id;comment:OAID;size:50;"`                                                             //OAID
	OaStatus         uint   `json:"oaStatus" form:"oaStatus" gorm:"column:oa_status;comment:OA状态;"`                                                         //OA状态
	CreatedBy        uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 调动申请 WcStaffTransferApplication自定义表名 wc_staff_transfer_application
func (WcStaffTransferApplication) TableName() string {
	return "wc_staff_transfer_application"
}
