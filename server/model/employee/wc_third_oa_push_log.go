// 自动生成模板WcThirdOaPush
package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// OA回调日志 结构体  WcThirdOaPush
type WcThirdOaPush struct {
	global.GVA_MODEL
	AppKey   string `json:"appKey" form:"appKey" gorm:"column:app_key;comment:应用ID;size:50;" binding:"required"`        //应用ID
	Flowtype string `json:"flowtype" form:"flowtype" gorm:"column:flowtype;comment:流程类型;size:50;" binding:"required"`   //流程类型
	FdId     string `json:"fdId" form:"fdId" gorm:"column:fdId;comment:流程id;size:50;" binding:"required"`               //流程id
	Subject  string `json:"subject" form:"subject" gorm:"column:subject;comment:流程标题;size:255;" binding:"required"`     //流程标题
	FdTmplId string `json:"fdTmplId" form:"fdTmplId" gorm:"column:fdTmplId;comment:流程模板id;size:50;" binding:"required"` //流程模板id
	Ip       string `json:"ip" form:"ip" gorm:"column:ip;comment:推送ip;size:50;" binding:"required"`                     //推送ip
	Status   uint   `json:"status" form:"status" gorm:"column:status;comment:流程状态;size:50;" binding:"required"`         //流程状态
	Log      string `json:"log" form:"log" gorm:"column:log;comment:日志内容;"`                                             //日志内容
}

// TableName OA回调日志 WcThirdOaPush自定义表名 wc_thirdoa_push_log
func (WcThirdOaPush) TableName() string {
	return "wc_thirdoa_push_log"
}
