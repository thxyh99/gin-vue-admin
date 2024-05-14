// 自动生成模板WcRankType
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 职级管理 结构体  WcRankType
type WcRankType struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:职工等级名称;size:200;" binding:"required"` //类型名称
}

// TableName 职级管理 WcRankType自定义表名 wc_rank
func (WcRankType) TableName() string {
	return "wc_rank_type"
}
