package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"strings"
)

// WcStaffMaterialsResponse 工资卡信息 结构体
type WcStaffMaterialsResponse struct {
	weChat.WcStaffMaterials
	StaffName            string   `json:"staffName"`            //员工名称
	JobNum               string   `json:"jobNum"`               //员工工号
	SkillCertificateList []string `json:"skillCertificateList"` //职称/技能证书
}

func (WcStaffMaterialsResponse) AssembleStaffMaterialsList(staffMaterials []weChat.WcStaffMaterials) (newStaffMaterials []WcStaffMaterialsResponse, err error) {
	var newStaffMater WcStaffMaterialsResponse

	for _, staffMaterial := range staffMaterials {
		newStaffMater.WcStaffMaterials = staffMaterial

		if newStaffMater.SkillCertificate != "" {
			newStaffMater.SkillCertificateList = strings.Split(staffMaterial.SkillCertificate, "###")
		}

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffMaterial.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffMaterialsList Err:", err)
			return
		}
		newStaffMater.StaffName = staff.Name
		newStaffMater.JobNum = staff.JobNum

		newStaffMaterials = append(newStaffMaterials, newStaffMater)
	}
	return
}

func (WcStaffMaterialsResponse) AssembleStaffMaterials(staffMaterials weChat.WcStaffMaterials) (newStaffMaterials WcStaffMaterialsResponse, err error) {
	newStaffMaterials.WcStaffMaterials = staffMaterials

	if staffMaterials.SkillCertificate != "" {
		newStaffMaterials.SkillCertificateList = strings.Split(staffMaterials.SkillCertificate, "###")
	}

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffMaterials.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffMaterials Err:", err)
		return
	}

	newStaffMaterials.StaffName = staff.Name
	newStaffMaterials.JobNum = staff.JobNum

	return
}
