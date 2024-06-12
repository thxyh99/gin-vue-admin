// 自动生成模板WcStaff
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 账号信息 结构体  WcStaff
type WcStaff struct {
	global.GVA_MODEL
	UserId           *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID(SSO);size:11;" binding:"required"`                                                                           //用户ID(SSO)
	Userid           string     `json:"userid" form:"userid" gorm:"column:userid;comment:企微成员UserID;size:200;" binding:"required"`                                                                          //企微成员UserID
	Name             string     `json:"name" form:"name" gorm:"column:name;comment:姓名;size:200;" binding:"required"`                                                                                        //姓名
	JobNum           string     `json:"jobNum" form:"jobNum" gorm:"column:job_num;comment:员工编码;size:100;" binding:"required"`                                                                               //员工编码
	Mobile           string     `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机;size:100;" binding:"required"`                                                                                  //手机
	IdNumber         string     `json:"idNumber" form:"idNumber" gorm:"column:id_number;comment:身份证号;size:100;" binding:"required"`                                                                         //身份证号
	IdAddress        string     `json:"idAddress" form:"idAddress" gorm:"column:id_address;comment:身份证地址;size:255;" binding:"required"`                                                                     //身份证地址
	HouseholdType    *int       `json:"householdType" form:"householdType" gorm:"column:household_type;comment:户籍类型(1:本地城镇 2:本地农村 3:外地城镇[省内] 4:外地农村[省内] 5:外地城镇[省外] 6:外地农村[省外]);size:1;" binding:"required"` //户籍类型(1:本地城镇 2:本地农村 3:外地城镇[省内] 4:外地农村[省内] 5:外地城镇[省外] 6:外地农村[省外])
	Gender           *int       `json:"gender" form:"gender" gorm:"column:gender;comment:性别(0未知1男2女);size:1;" binding:"required"`                                                                           //性别(0未知1男2女)
	Birthday         *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:出生日期;" binding:"required"`                                                                                   //出生日期
	NativePlace      string     `json:"nativePlace" form:"nativePlace" gorm:"column:native_place;comment:籍贯;size:255;" binding:"required"`                                                                  //籍贯
	Nation           *int       `json:"nation" form:"nation" gorm:"column:nation;comment:民族;size:2;" binding:"required"`                                                                                    //民族
	Height           *float64   `json:"height" form:"height" gorm:"column:height;comment:身高;size:4;" binding:"required"`                                                                                    //身高
	Weight           *float64   `json:"weight" form:"weight" gorm:"column:weight;comment:体重;size:4;" binding:"required"`                                                                                    //体重
	Marriage         *int       `json:"marriage" form:"marriage" gorm:"column:marriage;comment:婚姻状况(1:已婚 2:未婚 3:其他);size:1;" binding:"required"`                                                            //婚姻状况(1:已婚 2:未婚 3:其他)
	PoliticalOutlook *int       `json:"politicalOutlook" form:"politicalOutlook" gorm:"column:political_outlook;comment:政治面貌(1:团员 2:党员 3:群众 0:其他);size:1;" binding:"required"`                              //政治面貌(1:团员 2:党员 3:群众 0:其他)
	Address          string     `json:"address" form:"address" gorm:"column:address;comment:现居住地址;size:255;" binding:"required"`                                                                            //现居住地址
	SocialNumber     string     `json:"socialNumber" form:"socialNumber" gorm:"column:social_number;comment:社保电脑号;size:200;"`                                                                               //社保电脑号
	AccountNumber    string     `json:"accountNumber" form:"accountNumber" gorm:"column:account_number;comment:公积金账号;size:200;"`                                                                            //公积金账号
	PaymentPlace     string     `json:"paymentPlace" form:"paymentPlace" gorm:"column:payment_place;comment:社保公积金缴纳地;size:255;"`
}

// TableName 账号信息 WcStaff自定义表名 wc_staff
func (WcStaff) TableName() string {
	return "wc_staff"
}
