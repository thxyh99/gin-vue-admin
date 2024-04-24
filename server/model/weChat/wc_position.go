// 自动生成模板WcPosition
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 职位管理 结构体  WcPosition
type WcPosition struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:职位名称;size:200;" binding:"required"`  //职位名称 
}


// TableName 职位管理 WcPosition自定义表名 wc_position
func (WcPosition) TableName() string {
  return "wc_position"
}

