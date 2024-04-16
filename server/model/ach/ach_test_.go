// 自动生成模板AchTest
package ach

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 绩效测试 结构体  AchTest
type AchTest struct {
 global.GVA_MODEL 
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;size:100;" binding:"required"`  //用户UUID 
      Username  string `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:100;" binding:"required"`  //用户登录名 
      Month  string `json:"month" form:"month" gorm:"column:month;comment:绩效月份;size:7;" binding:"required"`  //绩效月份 
      Amount  *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:绩效金额;size:10;" binding:"required"`  //绩效金额 
}


// TableName 绩效测试 AchTest自定义表名 ach_test
func (AchTest) TableName() string {
  return "ach_test"
}

