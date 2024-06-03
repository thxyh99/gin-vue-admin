// 自动生成模板WcSalaryTypeField
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 工资类型款项 结构体  WcSalaryTypeField
type WcSalaryTypeField struct {
 global.GVA_MODEL 
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:工资类型;size:2;" binding:"required"`  //工资类型 
      Field  string `json:"field" form:"field" gorm:"column:field;comment:款项字段;size:50;" binding:"required"`  //款项字段 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:款项定义;size:100;" binding:"required"`  //款项定义 
      IsRequired  *int `json:"isRequired" form:"isRequired" gorm:"column:is_required;comment:是否必选;size:1;" binding:"required"`  //是否必选 
}


// TableName 工资类型款项 WcSalaryTypeField自定义表名 wc_salary_type_field
func (WcSalaryTypeField) TableName() string {
  return "wc_salary_type_field"
}

