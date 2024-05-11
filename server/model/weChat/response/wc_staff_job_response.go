package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// WcStaffJobResponse 工作信息 结构体
type WcStaffJobResponse struct {
	weChat.WcStaffJob
	StaffName     string `json:"staffName"`     //员工名称
	JobNum        string `json:"jobNum"`        //员工工号
	TypeText      string `json:"typeText"`      //员工类型
	StatusText    string `json:"statusText"`    //员工状态
	TryPeriodText string `json:"tryPeriodText"` //试用期
}

func (WcStaffJobResponse) AssembleStaffJobList(staffInfos []weChat.WcStaffJob) (newStaffInfos []WcStaffJobResponse, err error) {
	var newStaffJob WcStaffJobResponse
	configInfo := config.GetConfigInfo()

	for _, staffJob := range staffInfos {
		newStaffJob.WcStaffJob = staffJob
		staffJobTypeText, _ := utils.Find(configInfo.StaffJobType, *staffJob.Type)
		newStaffJob.TypeText = staffJobTypeText
		staffJobStatusText, _ := utils.Find(configInfo.StaffJobStatus, *staffJob.Status)
		newStaffJob.StatusText = staffJobStatusText
		staffJobTryPeriodText, _ := utils.Find(configInfo.StaffJobTryPeriod, *staffJob.TryPeriod)
		newStaffJob.TryPeriodText = staffJobTryPeriodText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffJob.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffInfoList Err:", err)
			return
		}
		newStaffJob.StaffName = staff.Name
		newStaffJob.JobNum = staff.JobNum

		newStaffInfos = append(newStaffInfos, newStaffJob)
	}
	return
}

func (WcStaffJobResponse) AssembleStaffJob(staffJob weChat.WcStaffJob) (newStaffJob WcStaffJobResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffJob.WcStaffJob = staffJob
	staffJobTypeText, _ := utils.Find(configInfo.StaffJobType, *staffJob.Type)
	newStaffJob.TypeText = staffJobTypeText
	staffJobStatusText, _ := utils.Find(configInfo.StaffJobStatus, *staffJob.Status)
	newStaffJob.StatusText = staffJobStatusText
	staffJobTryPeriodText, _ := utils.Find(configInfo.StaffJobTryPeriod, *staffJob.TryPeriod)
	newStaffJob.TryPeriodText = staffJobTryPeriodText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffJob.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffInfo Err:", err)
		return
	}

	newStaffJob.StaffName = staff.Name
	newStaffJob.JobNum = staff.JobNum

	return
}
