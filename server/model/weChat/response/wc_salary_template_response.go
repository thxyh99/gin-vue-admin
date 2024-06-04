package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcSalaryTemplateResponse  工资模版结构体
type WcSalaryTemplateResponse struct {
	weChat.WcSalaryTemplate
	TypeText     string                 `json:"typeText"`     //工资类型
	RankTypeText string                 `json:"rankTypeText"` //职级类型
	Fields       []weChat.WcSalaryField `json:"fields"`       //模版绑定的工资字段信息
}
