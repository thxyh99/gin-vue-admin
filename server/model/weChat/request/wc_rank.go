package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WcRankSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

// CreateWcRankRequest 创建参数
type CreateWcRankRequest struct {
	Type *int   `json:"type" form:"type"  binding:"required"` //职工等级类型
	Name string `json:"name" form:"name"  binding:"required"` //职工等级名称
}

type SearchRankParams struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Type           *int       `json:"type" form:"type" ` //职工等级类型
	Name           string     `json:"name" form:"name" ` //职工等级名称
	request.PageInfo
}
