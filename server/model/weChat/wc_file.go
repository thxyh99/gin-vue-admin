// 自动生成模板WcFile
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 文件管理 结构体  WcFile
type WcFile struct {
 global.GVA_MODEL 
      StaffId  *int `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:11;" binding:"required"`  //员工ID 
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:文件分类(0:其它 1:合同附件 2:身份证(人像) 3:身份证(国徽) 4:学历证书 5:学位证书 6:前公司离职证明 7:员工入职申请表 8:试用期管理规定 9:个人简历 10:职称/技能证书);size:2;" binding:"required"`  //文件分类(0:其它 1:合同附件 2:身份证(人像) 3:身份证(国徽) 4:学历证书 5:学位证书 6:前公司离职证明 7:员工入职申请表 8:试用期管理规定 9:个人简历 10:职称/技能证书) 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:255;" binding:"required"`  //文件名 
      Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址;size:255;" binding:"required"`  //文件地址 
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签;size:255;" binding:"required"`  //文件标签 
      Key  string `json:"key" form:"key" gorm:"column:key;comment:编号;size:255;" binding:"required"`  //编号 
}


// TableName 文件管理 WcFile自定义表名 wc_file
func (WcFile) TableName() string {
  return "wc_file"
}

