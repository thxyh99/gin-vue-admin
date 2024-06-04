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

type WcSalaryTemplateCreateRequest struct {
	weChat.WcSalaryTemplate
	Fields []WcSalaryTemplateCreateItem `json:"fields" form:"fields"`
}

type WcSalaryTemplateCreateItem struct {
	Field     string `json:"field" form:"field" `         //款项字段
	Name      string `json:"name" form:"name" `           //款项定义
	IsVisible *int   `json:"isVisible" form:"isVisible" ` //是否可见
}

type WcSalaryTemplateUpdateRequest struct {
	weChat.WcSalaryTemplate
	Fields []WcSalaryTemplateUpdateItem `json:"fields" form:"fields"`
}

type WcSalaryTemplateUpdateItem struct {
	ID        *int `json:"id" form:"id" `               //款项ID
	IsVisible *int `json:"isVisible" form:"isVisible" ` //是否可见
}
