package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// WcStaffEducationResponse 工作信息 结构体
type WcStaffEducationResponse struct {
	weChat.WcStaffEducation
	StaffName     string `json:"staffName"`     //员工名称
	JobNum        string `json:"jobNum"`        //员工工号
	EducationText string `json:"educationText"` //学历
}

func (WcStaffEducationResponse) AssembleStaffEducationList(staffEducations []weChat.WcStaffEducation) (newStaffEducations []WcStaffEducationResponse, err error) {
	var newStaffEducation WcStaffEducationResponse
	configInfo := config.GetConfigInfo()

	for _, staffEducation := range staffEducations {
		newStaffEducation.WcStaffEducation = staffEducation
		educationText, _ := utils.Find(configInfo.Education, *staffEducation.Education)
		newStaffEducation.EducationText = educationText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffEducation.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffInfoList Err:", err)
			return
		}
		newStaffEducation.StaffName = staff.Name
		newStaffEducation.JobNum = staff.JobNum

		newStaffEducations = append(newStaffEducations, newStaffEducation)
	}
	return
}

func (WcStaffEducationResponse) AssembleStaffEducation(staffEducation weChat.WcStaffEducation) (newStaffEducation WcStaffEducationResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffEducation.WcStaffEducation = staffEducation
	educationText, _ := utils.Find(configInfo.Education, *staffEducation.Education)
	newStaffEducation.EducationText = educationText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffEducation.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffInfo Err:", err)
		return
	}

	newStaffEducation.StaffName = staff.Name
	newStaffEducation.JobNum = staff.JobNum

	return
}
