package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"time"
)

type WcSalaryTemplateSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

type WcSalaryTemplateRequest struct {
	weChat.WcSalaryTemplate
	Fields []WcSalaryTemplateItem `json:"fields" form:"fields"`
}

type WcSalaryTemplateItem struct {
	Field     string `json:"field" form:"field" `         //款项字段
	Name      string `json:"name" form:"name" `           //款项定义
	IsVisible *int   `json:"isVisible" form:"isVisible" ` //是否可见
}
