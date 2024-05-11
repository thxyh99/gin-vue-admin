package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// WcStaffContactResponse 工作信息 结构体
type WcStaffContactResponse struct {
	weChat.WcStaffContact
	StaffName        string `json:"staffName"`        //员工名称
	JobNum           string `json:"jobNum"`           //员工工号
	RelationshipText string `json:"relationshipText"` //联系人关系
}

func (WcStaffContactResponse) AssembleStaffContactList(staffContacts []weChat.WcStaffContact) (newStaffContact []WcStaffContactResponse, err error) {
	var newStaffEducation WcStaffContactResponse
	configInfo := config.GetConfigInfo()

	for _, staffContact := range staffContacts {
		newStaffEducation.WcStaffContact = staffContact
		relationshipText, _ := utils.Find(configInfo.Relationship, *staffContact.Relationship)
		newStaffEducation.RelationshipText = relationshipText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffContact.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffInfoList Err:", err)
			return
		}
		newStaffEducation.StaffName = staff.Name
		newStaffEducation.JobNum = staff.JobNum

		newStaffContact = append(newStaffContact, newStaffEducation)
	}
	return
}

func (WcStaffContactResponse) AssembleStaffContact(staffContact weChat.WcStaffContact) (newStaffContact WcStaffContactResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffContact.WcStaffContact = staffContact
	relationshipText, _ := utils.Find(configInfo.Relationship, *staffContact.Relationship)
	newStaffContact.RelationshipText = relationshipText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffContact.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffInfo Err:", err)
		return
	}

	newStaffContact.StaffName = staff.Name
	newStaffContact.JobNum = staff.JobNum

	return
}
