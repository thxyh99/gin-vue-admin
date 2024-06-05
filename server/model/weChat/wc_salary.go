// 自动生成模板WcSalary
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// WcSalary 工资单发放结构体
type WcSalary struct {
	global.GVA_MODEL
	Month      string `json:"month" form:"month" gorm:"column:month;comment:所属月份(Ym);"`                     //所属月份(Ym)
	TemplateId *int   `json:"templateId" form:"templateId" gorm:"column:template_id;comment:工资模版;size:10;"` //工资模版
}

// TableName 工资单模版 WcSalary自定义表名 wc_salary
func (WcSalary) TableName() string {
	return "wc_salary"
}
