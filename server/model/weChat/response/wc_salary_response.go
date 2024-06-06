package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcSalaryResponse  工资发放结构体
type WcSalaryResponse struct {
	weChat.WcSalary
	Name         string `json:"name"`         //工资模版名称
	Type         *int   `json:"type"`         //工资类型
	TypeText     string `json:"typeText"`     //工资类型
	RankType     *int   `json:"rankType"`     //职级类型
	RankTypeText string `json:"rankTypeText"` //职级类型
}
