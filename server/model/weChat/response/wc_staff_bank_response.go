package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffBankResponse 工资卡信息 结构体
type WcStaffBankResponse struct {
	weChat.WcStaffBank
	StaffName string `json:"staffName"` //员工名称
	JobNum    string `json:"jobNum"`    //员工工号
}

func (WcStaffBankResponse) AssembleStaffBankList(staffBanks []weChat.WcStaffBank) (newStaffBanks []WcStaffBankResponse, err error) {
	var newStaffBank WcStaffBankResponse

	for _, staffBank := range staffBanks {
		newStaffBank.WcStaffBank = staffBank

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffBank.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffBankList Err:", err)
			return
		}
		newStaffBank.StaffName = staff.Name
		newStaffBank.JobNum = staff.JobNum

		newStaffBanks = append(newStaffBanks, newStaffBank)
	}
	return
}

func (WcStaffBankResponse) AssembleStaffBank(staffBank weChat.WcStaffBank) (newStaffBank WcStaffBankResponse, err error) {
	newStaffBank.WcStaffBank = staffBank

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffBank.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffBank Err:", err)
		return
	}

	newStaffBank.StaffName = staff.Name
	newStaffBank.JobNum = staff.JobNum

	return
}
