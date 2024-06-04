package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type WcSalaryField struct {
	global.GVA_MODEL
	TemplateId *int   `json:"type" form:"type" gorm:"column:type;comment:工资类型;size:10;" binding:"required"`                //工资模版ID
	Field      string `json:"field" form:"field" gorm:"column:field;comment:款项字段;size:50;" binding:"required"`             //款项字段
	Name       string `json:"name" form:"name" gorm:"column:name;comment:款项定义;size:100;" binding:"required"`               //款项定义
	IsVisible  *int   `json:"isVisible" form:"isVisible" gorm:"column:is_visible;comment:是否可见;size:1;" binding:"required"` //是否可见
}

// TableName 工资款项 WcSalaryField自定义表名 wc_salary_field
func (WcSalaryField) TableName() string {
	return "wc_salary_field"
}
