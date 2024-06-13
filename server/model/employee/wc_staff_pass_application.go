// 自动生成模板WcStaffPassApplication
package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 转正申请 结构体  WcStaffPassApplication
type WcStaffPassApplication struct {
	global.GVA_MODEL
	Title               string `json:"title" form:"title" gorm:"column:title;comment:转正标题;size:100;"`                                                                                    //转正标题
	StaffId             uint   `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;" binding:"required"`                                                           //员工ID
	Education           string `json:"education" form:"education" gorm:"column:education;type:enum();comment:学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他);" binding:"required"`      //学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他)
	EmploymentDate      string `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:入职日期;" binding:"required"`                                              //入职日期
	JobDepartment       string `json:"jobDepartment" form:"jobDepartment" gorm:"column:job_department;type:enum(10);comment:所属部门;" binding:"required"`                                   //所属部门
	SourcePosition      string `json:"sourcePosition" form:"sourcePosition" gorm:"column:source_position;type:enum(10);comment:原职位;" binding:"required"`                                 //原职位
	SourceLevel         string `json:"sourceLevel" form:"sourceLevel" gorm:"column:source_level;type:enum(10);comment:原职级;" binding:"required"`                                          //原职级
	TryPeriod           string `json:"tryPeriod" form:"tryPeriod" gorm:"column:try_period;type:enum();comment:试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他);" binding:"required"` //试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他)
	IsMananger          *bool  `json:"isMananger" form:"isMananger" gorm:"column:is_mananger;comment:是否是部门负责人;" binding:"required"`                                                      //是否是部门负责人
	Attachment          string `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;" binding:"required"`                                                             //附件
	SelfOpinion         string `json:"selfOpinion" form:"selfOpinion" gorm:"column:self_opinion;comment:个人自我鉴定;size:255;" binding:"required"`                                            //个人自我鉴定
	TryOpinion          string `json:"tryOpinion" form:"tryOpinion" gorm:"column:try_opinion;comment:用人部门意见;size:255;" binding:"required"`                                               //用人部门意见
	NewResponsibilities string `json:"newResponsibilities" form:"newResponsibilities" gorm:"column:new_responsibilities;type:enum(255);comment:转正后工作职责;" binding:"required"`             //转正后工作职责
	PassDate            string `json:"passDate" form:"passDate" gorm:"column:pass_date;type:enum();comment:转正时间;" binding:"required"`                                                    //转正时间
	NewPosition         string `json:"newPosition" form:"newPosition" gorm:"column:new_position;type:enum(10);comment:转正后职位;" binding:"required"`                                        //转正后职位
	NewJoblevel         string `json:"newJoblevel" form:"newJoblevel" gorm:"column:new_joblevel;type:enum(10);comment:转正后职级;" binding:"required"`                                        //转正后职级
	SubmitOpinion       string `json:"submitOpinion" form:"submitOpinion" gorm:"column:submit_opinion;comment:提交意见;size:255;" binding:"required"`                                        //提交意见
	OaId                string `json:"oaId" form:"oaId" gorm:"column:oa_id;comment:OAID;size:50;" binding:"required"`                                                                    //OAID
	OaStatus            uint   `json:"oaStatus" form:"oaStatus" gorm:"column:oa_status;comment:OA状态;" binding:"required"`                                                                //OA状态
	CreatedBy           uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy           uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy           uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 转正申请 WcStaffPassApplication自定义表名 wc_staff_pass_application
func (WcStaffPassApplication) TableName() string {
	return "wc_staff_pass_application"
}
