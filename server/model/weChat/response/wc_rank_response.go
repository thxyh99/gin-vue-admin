package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcRankResponse  员工职级结构体
type WcRankResponse struct {
	weChat.WcRank
	TypeText string `json:"typeText"` //职级类型
}

// WcRankItem 员工等级名称
type WcRankItem struct {
	ID   int    `json:"ID"`   // ID
	Name string `json:"name"` //等级名称
}
