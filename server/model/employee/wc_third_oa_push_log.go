// 自动生成模板WcThirdOaPush
package employee

import (
	"time"
)

// OA回调日志 结构体  WcThirdOaPush
type WcThirdOaPush struct {
	Appkey     string    `json:"appkey" form:"appkey" gorm:"column:appkey;comment:应用ID;size:50;" binding:"required"`         //应用ID
	FdTmplId   string    `json:"fdTmplId" form:"fdTmplId" gorm:"column:fdTmplId;comment:流程模板id;size:50;" binding:"required"` //流程模板id
	FdId       string    `json:"fdId" form:"fdId" gorm:"column:fdId;comment:流程id;size:50;" binding:"required"`               //流程id
	Subject    string    `json:"subject" form:"subject" gorm:"column:subject;comment:流程标题;size:255;" binding:"required"`     //流程标题
	Status     uint      `json:"status" form:"status" gorm:"column:status;comment:流程状态;" binding:"required"`                 //流程状态
	FormValues string    `json:"formValues" form:"formValues" gorm:"column:formValues;comment:日志内容;" binding:"required"`     //日志内容
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:createTime;comment:创建时间;size:19;"`                //创建时间
	Attachment string    `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;"`                          //附件
	Sign       string    `json:"sign" form:"sign" gorm:"column:sign;comment:签名;size:255;"`                                   //签名
	Ip         string    `json:"ip" form:"ip" gorm:"column:ip;comment:推送ip;size:50;"`                                        //推送ip
}

// TableName OA回调日志 WcThirdOaPush自定义表名 wc_thirdoa_push_log
func (WcThirdOaPush) TableName() string {
	return "wc_thirdoa_push_log"
}
