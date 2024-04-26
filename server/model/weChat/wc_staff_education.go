// 自动生成模板WcStaffEducation
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 学历信息 结构体  WcStaffEducation
type WcStaffEducation struct {
 global.GVA_MODEL 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID(SSO);size:11;" binding:"required"`  //用户ID(SSO) 
      Userid  string `json:"userid" form:"userid" gorm:"column:userid;comment:企微成员UserID;size:200;" binding:"required"`  //企微成员UserID 
      Education  *int `json:"education" form:"education" gorm:"column:education;comment:学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他);size:1;" binding:"required"`  //学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他) 
      School  string `json:"school" form:"school" gorm:"column:school;comment:毕业院校;size:200;" binding:"required"`  //毕业院校 
      Date  *time.Time `json:"date" form:"date" gorm:"column:date;comment:毕业日期;" binding:"required"`  //毕业日期 
      Major  string `json:"major" form:"major" gorm:"column:major;comment:专业;size:200;" binding:"required"`  //专业 
      Certificate  string `json:"certificate" form:"certificate" gorm:"column:certificate;comment:职称/技能证书;size:200;"`  //职称/技能证书 
}


// TableName 学历信息 WcStaffEducation自定义表名 wc_staff_education
func (WcStaffEducation) TableName() string {
  return "wc_staff_education"
}

