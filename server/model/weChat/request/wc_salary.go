package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WcSalarySearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	MonthStart     string     `json:"monthStart" form:"monthStart"`
	MonthEnd       string     `json:"monthEnd" form:"monthEnd"`
	TemplateId     *int       `json:"templateId" form:"templateId"`

	request.PageInfo
}
