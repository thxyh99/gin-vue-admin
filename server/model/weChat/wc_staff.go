// 自动生成模板WcStaff
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 账号信息 结构体  WcStaff
type WcStaff struct {
	global.GVA_MODEL
	UserId    *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID(SSO);size:11;" binding:"required"`          //用户ID(SSO)
	Userid    string `json:"userid" form:"userid" gorm:"column:userid;comment:企微成员UserID;size:200;" binding:"required"`         //企微成员UserID
	JobNum    string `json:"jobNum" form:"jobNum" gorm:"column:job_num;comment:员工工号;size:100;" binding:"required"`              //员工工号
	Name      string `json:"name" form:"name" gorm:"column:name;comment:成员名称;size:200;" binding:"required"`                     //成员名称
	Alias     string `json:"alias" form:"alias" gorm:"column:alias;comment:别名;size:200;"`                                       //别名
	Gender    *int   `json:"gender" form:"gender" gorm:"column:gender;comment:性别(0未知1男2女);size:1;" binding:"required"`          //性别(0未知1男2女)
	IsLeader  *int   `json:"isLeader" form:"isLeader" gorm:"column:is_leader;comment:是否领导(1:是 0:否);size:1;" binding:"required"` //是否领导(1:是 0:否)
	Mobile    string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机;size:100;" binding:"required"`                 //手机
	Telephone string `json:"telephone" form:"telephone" gorm:"column:telephone;comment:座机;size:100;"`                           //座机
	Email     string `json:"email" form:"email" gorm:"column:email;comment:个人邮箱;size:200;"`                                     //个人邮箱
	Address   string `json:"address" form:"address" gorm:"column:address;comment:地址;size:255;"`                                 //地址
	BizMail   string `json:"bizMail" form:"bizMail" gorm:"column:biz_mail;comment:企业邮箱;size:200;"`                              //企业邮箱
	NameEn    string `json:"nameEn" form:"nameEn" gorm:"column:name_en;comment:英文名;size:200;"`                                  //英文名
	Status    *int   `json:"status" form:"status" gorm:"column:status;comment:激活状态(1已激活，2已禁用，4未激活，5退出企业);size:1;"`              //激活状态(1已激活，2已禁用，4未激活，5退出企业)
}

// TableName 账号信息 WcStaff自定义表名 wc_staff
func (WcStaff) TableName() string {
	return "wc_staff"
}
