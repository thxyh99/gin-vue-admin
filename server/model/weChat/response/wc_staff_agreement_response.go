package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// WcStaffAgreementResponse 工作信息 结构体
type WcStaffAgreementResponse struct {
	weChat.WcStaffAgreement
	StaffName string `json:"staffName"` //员工名称
	JobNum    string `json:"jobNum"`    //员工工号
	TypeText  string `json:"typeText"`  //合同类型
}

func (WcStaffAgreementResponse) AssembleStaffAgreementList(staffAgreements []weChat.WcStaffAgreement) (newStaffAgreements []WcStaffAgreementResponse, err error) {
	var newStaffAgreement WcStaffAgreementResponse
	configInfo := config.GetConfigInfo()

	for _, staffAgreement := range staffAgreements {
		newStaffAgreement.WcStaffAgreement = staffAgreement
		typeText, _ := utils.Find(configInfo.AgreementType, *staffAgreement.Type)
		newStaffAgreement.TypeText = typeText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffAgreement.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffAgreementList Err:", err)
			return
		}
		newStaffAgreement.StaffName = staff.Name
		newStaffAgreement.JobNum = staff.JobNum

		newStaffAgreements = append(newStaffAgreements, newStaffAgreement)
	}
	return
}

func (WcStaffAgreementResponse) AssembleStaffAgreement(staffAgreement weChat.WcStaffAgreement) (newStaffAgreement WcStaffAgreementResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffAgreement.WcStaffAgreement = staffAgreement
	typeText, _ := utils.Find(configInfo.AgreementType, *staffAgreement.Type)
	newStaffAgreement.TypeText = typeText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffAgreement.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffAgreement Err:", err)
		return
	}

	newStaffAgreement.StaffName = staff.Name
	newStaffAgreement.JobNum = staff.JobNum

	return
}
