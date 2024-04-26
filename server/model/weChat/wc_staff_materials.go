// 自动生成模板WcStaffMaterials
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 证件材料 结构体  WcStaffMaterials
type WcStaffMaterials struct {
 global.GVA_MODEL 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID(SSO);size:10;" binding:"required"`  //用户ID(SSO) 
      Userid  string `json:"userid" form:"userid" gorm:"column:userid;comment:企微成员UserID;size:200;" binding:"required"`  //企微成员UserID 
      IdCardPortrait  string `json:"idCardPortrait" form:"idCardPortrait" gorm:"column:id_card_portrait;comment:身份证(人像);size:255;"`  //身份证(人像) 
      IdCardNational  string `json:"idCardNational" form:"idCardNational" gorm:"column:id_card_national;comment:身份证(国徽);size:255;"`  //身份证(国徽) 
      EducationCertificate  string `json:"educationCertificate" form:"educationCertificate" gorm:"column:education_certificate;comment:学历证书;size:255;"`  //学历证书 
      DegreeCertificate  string `json:"degreeCertificate" form:"degreeCertificate" gorm:"column:degree_certificate;comment:学位证书;size:255;"`  //学位证书 
      ResignationCertificate  string `json:"resignationCertificate" form:"resignationCertificate" gorm:"column:resignation_certificate;comment:前公司离职证明;size:255;"`  //前公司离职证明 
      OnboardingForm  string `json:"onboardingForm" form:"onboardingForm" gorm:"column:onboarding_form;comment:入职补充登记表;size:255;"`  //入职补充登记表 
}


// TableName 证件材料 WcStaffMaterials自定义表名 wc_staff_materials
func (WcStaffMaterials) TableName() string {
  return "wc_staff_materials"
}

