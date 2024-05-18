// 自动生成模板WcStaffEmploymentApplication
package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 入职申请 结构体  WcStaffEmploymentApplication
type WcStaffEmploymentApplication struct {
 global.GVA_MODEL 
      Title  string `json:"title" form:"title" gorm:"column:title;comment:入职标题;size:100;"`  //入职标题 
      StaffId  *int `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;" binding:"required"`  //员工ID 
      EmploymentDate  *time.Time `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:入职日期;" binding:"required"`  //入职日期 
      EmploymentDepartment  string `json:"employmentDepartment" form:"employmentDepartment" gorm:"column:employment_department;type:enum(10);comment:入职部门;" binding:"required"`  //入职部门 
      JobPosition  string `json:"jobPosition" form:"jobPosition" gorm:"column:job_position;type:enum(10);comment:入职职位;" binding:"required"`  //入职职位 
      SocialNumber  string `json:"socialNumber" form:"socialNumber" gorm:"column:social_number;comment:社保电脑号;size:200;" binding:"required"`  //社保电脑号 
      AccountNumber  string `json:"accountNumber" form:"accountNumber" gorm:"column:account_number;comment:公积金账号;size:200;" binding:"required"`  //公积金账号 
      HouseholdType  string `json:"householdType" form:"householdType" gorm:"column:household_type;type:enum();comment:户籍类型;" binding:"required"`  //户籍类型 
      PaymentPlace  string `json:"paymentPlace" form:"paymentPlace" gorm:"column:payment_place;type:enum(255);comment:社保公积金缴纳地;" binding:"required"`  //社保公积金缴纳地 
      Gender  string `json:"gender" form:"gender" gorm:"column:gender;type:enum();comment:性别(0未知1男2女);" binding:"required"`  //性别(0未知1男2女) 
      Birthday  *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:出生年月;" binding:"required"`  //出生年月 
      NativePlace  string `json:"nativePlace" form:"nativePlace" gorm:"column:native_place;comment:籍贯;size:255;" binding:"required"`  //籍贯 
      Nationality  string `json:"nationality" form:"nationality" gorm:"column:nationality;type:enum(50);comment:民族;" binding:"required"`  //民族 
      Height  *float64 `json:"height" form:"height" gorm:"column:height;comment:身高;size:4;" binding:"required"`  //身高 
      Weight  *float64 `json:"weight" form:"weight" gorm:"column:weight;comment:体重;size:4;" binding:"required"`  //体重 
      Marriage  string `json:"marriage" form:"marriage" gorm:"column:marriage;type:enum();comment:婚否(1:已婚 2:未婚 3:其他);" binding:"required"`  //婚否(1:已婚 2:未婚 3:其他) 
      PoliticalOutlook  string `json:"politicalOutlook" form:"politicalOutlook" gorm:"column:political_outlook;type:enum();comment:政治面貌(1:团员 2:党员 3:群众 0:其他);" binding:"required"`  //政治面貌(1:团员 2:党员 3:群众 0:其他) 
      Education  string `json:"education" form:"education" gorm:"column:education;type:enum();comment:学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他);" binding:"required"`  //学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他) 
      Major  string `json:"major" form:"major" gorm:"column:major;comment:专业;size:200;" binding:"required"`  //专业 
      School  string `json:"school" form:"school" gorm:"column:school;comment:毕业院校;size:200;" binding:"required"`  //毕业院校 
      GraduationDate  *time.Time `json:"graduationDate" form:"graduationDate" gorm:"column:graduation_date;comment:毕业时间;" binding:"required"`  //毕业时间 
      Certificate  string `json:"certificate" form:"certificate" gorm:"column:certificate;comment:职称技能证书;" binding:"required"`  //职称技能证书 
      IdNumber  string `json:"idNumber" form:"idNumber" gorm:"column:id_number;comment:身份证号码;size:100;" binding:"required"`  //身份证号码 
      IdAddress  string `json:"idAddress" form:"idAddress" gorm:"column:id_address;comment:身份证地址;size:255;" binding:"required"`  //身份证地址 
      BankAccount  string `json:"bankAccount" form:"bankAccount" gorm:"column:bank_account;comment:银行账号;size:100;" binding:"required"`  //银行账号 
      BankName  string `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:开户支行信息;size:200;" binding:"required"`  //开户支行信息 
      ContactWechat  string `json:"contactWechat" form:"contactWechat" gorm:"column:contact_wechat;comment:本人联系微信;size:200;" binding:"required"`  //本人联系微信 
      Mobile  string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机;size:100;" binding:"required"`  //手机 
      HomeAddress  string `json:"homeAddress" form:"homeAddress" gorm:"column:home_address;comment:常住地址;size:255;" binding:"required"`  //常住地址 
      LicenseAttachment  string `json:"licenseAttachment" form:"licenseAttachment" gorm:"column:license_attachment;comment:证件资料附件上传;" binding:"required"`  //证件资料附件上传 
      RelationShip  string `json:"relationShip" form:"relationShip" gorm:"column:relation_ship;type:enum();comment:联系人关系(1:父母 2:配偶 3:子女 0:其他);" binding:"required"`  //联系人关系(1:父母 2:配偶 3:子女 0:其他) 
      RelationName  string `json:"relationName" form:"relationName" gorm:"column:relation_name;comment:紧急联系人姓名;size:200;" binding:"required"`  //紧急联系人姓名 
      RelationMobile  string `json:"relationMobile" form:"relationMobile" gorm:"column:relation_mobile;comment:紧急联系人电话;size:100;" binding:"required"`  //紧急联系人电话 
      RelationAddress  string `json:"relationAddress" form:"relationAddress" gorm:"column:relation_address;comment:联系人常住地址;size:255;" binding:"required"`  //联系人常住地址 
      IsCeopass  *bool `json:"isCeopass" form:"isCeopass" gorm:"column:is_ceopass;comment:是否经过首席执行官;" binding:"required"`  //是否经过首席执行官 
      IsBodychecknormal  *bool `json:"isBodychecknormal" form:"isBodychecknormal" gorm:"column:is_bodychecknormal;comment:入职体检是否正常;" binding:"required"`  //入职体检是否正常 
      JobLevel  string `json:"jobLevel" form:"jobLevel" gorm:"column:job_level;type:enum(10);comment:职级;" binding:"required"`  //职级 
      TryPeriod  string `json:"tryPeriod" form:"tryPeriod" gorm:"column:try_period;type:enum();comment:试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他);" binding:"required"`  //试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他) 
      EmployingOpinion  string `json:"employingOpinion" form:"employingOpinion" gorm:"column:employing_opinion;comment:用人部门意见;size:255;" binding:"required"`  //用人部门意见 
      HumanOpinion  string `json:"humanOpinion" form:"humanOpinion" gorm:"column:human_opinion;comment:人资部意见;size:255;" binding:"required"`  //人资部意见 
      UrgencyNotification  string `json:"urgencyNotification" form:"urgencyNotification" gorm:"column:urgency_notification;comment:通知紧急程度;size:255;" binding:"required"`  //通知紧急程度 
      OnboardingOpinion  string `json:"onboardingOpinion" form:"onboardingOpinion" gorm:"column:onboarding_opinion;type:enum(255);comment:入职意见;" binding:"required"`  //入职意见 
      OaId  string `json:"oaId" form:"oaId" gorm:"column:oa_id;comment:OAID;size:50;" binding:"required"`  //OAID 
      OaStatus  *bool `json:"oaStatus" form:"oaStatus" gorm:"column:oa_status;comment:OA状态;" binding:"required"`  //OA状态 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 入职申请 WcStaffEmploymentApplication自定义表名 wc_staff_employment_application
func (WcStaffEmploymentApplication) TableName() string {
  return "wc_staff_employment_application"
}

