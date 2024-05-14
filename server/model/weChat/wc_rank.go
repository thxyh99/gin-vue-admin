// 自动生成模板WcRank
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 职级管理 结构体  WcRank
type WcRank struct {
 global.GVA_MODEL 
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:职工等级类型;size:6;" binding:"required"`  //职工等级类型 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:职工等级名称;size:200;" binding:"required"`  //职工等级名称 
}


// TableName 职级管理 WcRank自定义表名 wc_rank
func (WcRank) TableName() string {
  return "wc_rank"
}

