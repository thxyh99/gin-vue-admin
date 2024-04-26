// 自动生成模板WcStaffBank
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 银行卡信息 结构体  WcStaffBank
type WcStaffBank struct {
 global.GVA_MODEL 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID(SSO);size:10;" binding:"required"`  //用户ID(SSO) 
      Userid  string `json:"userid" form:"userid" gorm:"column:userid;comment:企微成员UserID;size:200;" binding:"required"`  //企微成员UserID 
      CardNumber  string `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:银行卡号;size:100;" binding:"required"`  //银行卡号 
      Bank  string `json:"bank" form:"bank" gorm:"column:bank;comment:开户行;size:200;" binding:"required"`  //开户行 
}


// TableName 银行卡信息 WcStaffBank自定义表名 wc_staff_bank
func (WcStaffBank) TableName() string {
  return "wc_staff_bank"
}

