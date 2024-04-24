// 自动生成模板WcStaff
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 员工管理 结构体  WcStaff
type WcStaff struct {
 global.GVA_MODEL 
      Userid  string `json:"userid" form:"userid" gorm:"column:userid;comment:成员UserID;size:200;" binding:"required"`  //成员UserID 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:成员名称;size:200;" binding:"required"`  //成员名称 
      Department  string `json:"department" form:"department" gorm:"column:department;comment:成员所属部门ID;size:200;" binding:"required"`  //成员所属部门ID 
      Order  *int `json:"order" form:"order" gorm:"column:order;comment:排序;size:1;" binding:"required"`  //排序 
      PositionId  *int `json:"positionId" form:"positionId" gorm:"column:position_id;comment:职务信息ID;size:10;" binding:"required"`  //职务信息ID 
      Position  string `json:"position" form:"position" gorm:"column:position;comment:职务信息;size:200;" binding:"required"`  //职务信息 
      Gender  *int `json:"gender" form:"gender" gorm:"column:gender;comment:性别(0未知1男2女);size:1;" binding:"required"`  //性别(0未知1男2女) 
      Email  string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:200;"`  //邮箱 
      BizMail  string `json:"bizMail" form:"bizMail" gorm:"column:biz_mail;comment:企业邮箱;size:200;"`  //企业邮箱 
      IsLeaderInDept  string `json:"isLeaderInDept" form:"isLeaderInDept" gorm:"column:is_leader_in_dept;comment:表示在所在的部门内是否为部门负责人;size:200;"`  //表示在所在的部门内是否为部门负责人 
      DirectLeader  string `json:"directLeader" form:"directLeader" gorm:"column:direct_leader;comment:直属上级UserID;size:200;"`  //直属上级UserID 
      Avatar  string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;size:255;"`  //头像 
      ThumbAvatar  string `json:"thumbAvatar" form:"thumbAvatar" gorm:"column:thumb_avatar;comment:头像缩略图;size:255;"`  //头像缩略图 
      Telephone  string `json:"telephone" form:"telephone" gorm:"column:telephone;comment:座机;size:100;"`  //座机 
      Alias  string `json:"alias" form:"alias" gorm:"column:alias;comment:别名;size:200;"`  //别名 
      Address  string `json:"address" form:"address" gorm:"column:address;comment:地址;size:255;"`  //地址 
      OpenUserid  string `json:"openUserid" form:"openUserid" gorm:"column:open_userid;comment:全局唯一;size:100;"`  //全局唯一 
      MainDepartment  *int `json:"mainDepartment" form:"mainDepartment" gorm:"column:main_department;comment:主部门;size:10;"`  //主部门 
      Extattr  string `json:"extattr" form:"extattr" gorm:"column:extattr;comment:扩展属性;type:text;"`  //扩展属性 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:激活状态(1已激活，2已禁用，4未激活，5退出企业);size:1;" binding:"required"`  //激活状态(1已激活，2已禁用，4未激活，5退出企业) 
      QrCode  string `json:"qrCode" form:"qrCode" gorm:"column:qr_code;comment:员工个人二维码;size:255;"`  //员工个人二维码 
      ExternalPosition  string `json:"externalPosition" form:"externalPosition" gorm:"column:external_position;comment:对外职务;size:200;"`  //对外职务 
      ExternalProfile  string `json:"externalProfile" form:"externalProfile" gorm:"column:external_profile;comment:成员对外属性;type:text;"`  //成员对外属性 
}


// TableName 员工管理 WcStaff自定义表名 wc_staff
func (WcStaff) TableName() string {
  return "wc_staff"
}

