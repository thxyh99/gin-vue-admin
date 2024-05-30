// 自动生成模板WcSalaryTemplate
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 工资单模版 结构体  WcSalaryTemplate
type WcSalaryTemplate struct {
 global.GVA_MODEL 
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金);size:2;" binding:"required"`  //工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金) 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:200;" binding:"required"`  //名称 
      RankType  *int `json:"rankType" form:"rankType" gorm:"column:rank_type;comment:职级类型;size:6;" binding:"required"`  //职级类型 
}


// TableName 工资单模版 WcSalaryTemplate自定义表名 wc_salary_template
func (WcSalaryTemplate) TableName() string {
  return "wc_salary_template"
}

