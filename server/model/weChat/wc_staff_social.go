// 自动生成模板WcStaffSocial
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 社保公积金管理 结构体  WcStaffSocial
type WcStaffSocial struct {
 global.GVA_MODEL 
      Account  string `json:"account" form:"account" gorm:"column:account;comment:账号(社保:电脑号|个人社保号 公积金:个人账号);size:100;" binding:"required"`  //账号(社保:电脑号|个人社保号 公积金:个人账号) 
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:类型(1:深圳社保 2:深圳公积金 3:东莞社保 4:东莞公积金);size:2;" binding:"required"`  //类型(1:深圳社保 2:深圳公积金 3:东莞社保 4:东莞公积金) 
      StaffId  *int `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:11;" binding:"required"`  //员工ID 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:200;" binding:"required"`  //姓名 
      CredentialType  *int `json:"credentialType" form:"credentialType" gorm:"column:credential_type;comment:证件类型(1:身份证 2:港澳通行证);size:2;" binding:"required"`  //证件类型(1:身份证 2:港澳通行证) 
      CredentialNumber  string `json:"credentialNumber" form:"credentialNumber" gorm:"column:credential_number;comment:证件号码;size:100;" binding:"required"`  //证件号码 
      TotalSocial  *float64 `json:"totalSocial" form:"totalSocial" gorm:"column:total_social;comment:社保缴费合计;size:10;" binding:"required"`  //社保缴费合计 
      TotalSocialSelf  *float64 `json:"totalSocialSelf" form:"totalSocialSelf" gorm:"column:total_social_self;comment:社保缴费个人合计;size:10;" binding:"required"`  //社保缴费个人合计 
      TotalSocialUnit  *float64 `json:"totalSocialUnit" form:"totalSocialUnit" gorm:"column:total_social_unit;comment:社保缴费单位合计;size:10;" binding:"required"`  //社保缴费单位合计 
      PensionBase  *float64 `json:"pensionBase" form:"pensionBase" gorm:"column:pension_base;comment:养老缴纳基数;size:10;" binding:"required"`  //养老缴纳基数 
      PensionSelf  *float64 `json:"pensionSelf" form:"pensionSelf" gorm:"column:pension_self;comment:养老个人缴费;size:10;" binding:"required"`  //养老个人缴费 
      PensionUnit  *float64 `json:"pensionUnit" form:"pensionUnit" gorm:"column:pension_unit;comment:养老单位缴费;size:10;" binding:"required"`  //养老单位缴费 
      MedicalBase  *float64 `json:"medicalBase" form:"medicalBase" gorm:"column:medical_base;comment:医疗缴纳基数;size:10;" binding:"required"`  //医疗缴纳基数 
      MedicalSelf  *float64 `json:"medicalSelf" form:"medicalSelf" gorm:"column:medical_self;comment:医疗个人缴费;size:10;" binding:"required"`  //医疗个人缴费 
      MedicalUnit  *float64 `json:"medicalUnit" form:"medicalUnit" gorm:"column:medical_unit;comment:医疗单位缴费;size:10;" binding:"required"`  //医疗单位缴费 
      UnemployedBase  *float64 `json:"unemployedBase" form:"unemployedBase" gorm:"column:unemployed_base;comment:失业缴纳基数;size:10;" binding:"required"`  //失业缴纳基数 
      UnemployedSelf  *float64 `json:"unemployedSelf" form:"unemployedSelf" gorm:"column:unemployed_self;comment:失业个人缴费;size:10;" binding:"required"`  //失业个人缴费 
      UnemployedUnit  *float64 `json:"unemployedUnit" form:"unemployedUnit" gorm:"column:unemployed_unit;comment:失业单位缴费;size:10;" binding:"required"`  //失业单位缴费 
      IcInsuranceBase  *float64 `json:"icInsuranceBase" form:"icInsuranceBase" gorm:"column:ic_insurance_base;comment:工商保险缴费基数;size:10;" binding:"required"`  //工商保险缴费基数 
      IcInsuranceUnit  *float64 `json:"icInsuranceUnit" form:"icInsuranceUnit" gorm:"column:ic_insurance_unit;comment:工商保险单位缴费;size:10;" binding:"required"`  //工商保险单位缴费 
      BirthBase  *float64 `json:"birthBase" form:"birthBase" gorm:"column:birth_base;comment:生育医疗缴费基数;size:10;" binding:"required"`  //生育医疗缴费基数 
      BirthUnit  *float64 `json:"birthUnit" form:"birthUnit" gorm:"column:birth_unit;comment:生育医疗单位缴费;size:10;" binding:"required"`  //生育医疗单位缴费 
      TotalHousing  *float64 `json:"totalHousing" form:"totalHousing" gorm:"column:total_housing;comment:公积金缴费合计;size:10;" binding:"required"`  //公积金缴费合计 
      TotalHousingSelf  *float64 `json:"totalHousingSelf" form:"totalHousingSelf" gorm:"column:total_housing_self;comment:公积金个人缴费合计;size:10;" binding:"required"`  //公积金个人缴费合计 
      TotalHousingUnit  *float64 `json:"totalHousingUnit" form:"totalHousingUnit" gorm:"column:total_housing_unit;comment:公积金公司缴费合计;size:10;" binding:"required"`  //公积金公司缴费合计 
      HousingBase  *float64 `json:"housingBase" form:"housingBase" gorm:"column:housing_base;comment:缴存基数;size:10;" binding:"required"`  //缴存基数 
      HousingRatioSelf  *float64 `json:"housingRatioSelf" form:"housingRatioSelf" gorm:"column:housing_ratio_self;comment:个人缴存比例;size:10;" binding:"required"`  //个人缴存比例 
      HousingRatioUnit  *float64 `json:"housingRatioUnit" form:"housingRatioUnit" gorm:"column:housing_ratio_unit;comment:单位缴存比例;size:10;" binding:"required"`  //单位缴存比例 
      PeriodStart  string `json:"periodStart" form:"periodStart" gorm:"column:period_start;comment:费款所属期起;size:10;" binding:"required"`  //费款所属期起 
      PeriodEnd  string `json:"periodEnd" form:"periodEnd" gorm:"column:period_end;comment:费款所属期止;size:10;" binding:"required"`  //费款所属期止 
}


// TableName 社保公积金管理 WcStaffSocial自定义表名 wc_staff_social
func (WcStaffSocial) TableName() string {
  return "wc_staff_social"
}

