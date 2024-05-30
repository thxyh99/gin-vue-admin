// 自动生成模板WcSalary
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 工资单模版 结构体  WcSalary
type WcSalary struct {
 global.GVA_MODEL 
      Date  string `json:"date" form:"date" gorm:"column:date;comment:所属月份(Ym);"`  //所属月份(Ym) 
      TemplateId  *int `json:"templateId" form:"templateId" gorm:"column:template_id;comment:工资模版;size:10;"`  //工资模版 
}


// TableName 工资单模版 WcSalary自定义表名 wc_salary
func (WcSalary) TableName() string {
  return "wc_salary"
}

