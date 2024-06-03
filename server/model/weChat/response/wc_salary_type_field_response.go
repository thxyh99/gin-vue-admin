package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcSalaryTypeFieldResponse  工资类型款项结构体
type WcSalaryTypeFieldResponse struct {
	weChat.WcSalaryTypeField
	TypeText       string `json:"typeText"`       //工资类型
	IsRequiredText string `json:"isRequiredText"` //是否必选款项类型
}
