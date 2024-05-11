package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// WcStaffInfoResponse 账号信息 结构体
type WcStaffInfoResponse struct {
	weChat.WcStaffInfo
	StaffName            string `json:"staffName"`             //员工名称
	JobNum               string `json:"jobNum"`                //员工工号
	TypeText             string `json:"typeText"`              //员工考勤类型
	RankText             string `json:"rankText"`              //员工职级
	HouseholdTypeText    string `json:"householdTypeText"`     //户籍类型
	NationText           string `json:"nationText"`            //员工民族
	MarriageText         string `json:"marriageText"`          //员工婚否
	PoliticalOutlookText string `json:"politicalOutlookText" ` //员工政治面貌
}

func (WcStaffInfoResponse) AssembleStaffInfoList(staffInfos []weChat.WcStaffInfo) (newStaffInfos []WcStaffInfoResponse, err error) {
	var newStaffInfo WcStaffInfoResponse
	configInfo := config.GetConfigInfo()

	for _, staffInfo := range staffInfos {
		newStaffInfo.WcStaffInfo = staffInfo
		staffInfoTypeText, _ := utils.Find(configInfo.StaffInfoType, *staffInfo.Type)
		newStaffInfo.TypeText = staffInfoTypeText
		staffOfficeRankText, _ := utils.Find(configInfo.StaffOfficeRank, *staffInfo.Rank)
		newStaffInfo.RankText = staffOfficeRankText
		householdTypeText, _ := utils.Find(configInfo.HouseholdType, *staffInfo.HouseholdType)
		newStaffInfo.HouseholdTypeText = householdTypeText
		nationText, _ := utils.Find(configInfo.Nation, *staffInfo.Nation)
		newStaffInfo.NationText = nationText
		marriageText, _ := utils.Find(configInfo.Marriage, *staffInfo.Marriage)
		newStaffInfo.MarriageText = marriageText
		politicalOutlookText, _ := utils.Find(configInfo.PoliticalOutlook, *staffInfo.PoliticalOutlook)
		newStaffInfo.PoliticalOutlookText = politicalOutlookText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffInfo.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffInfoList Err:", err)
			return
		}
		newStaffInfo.StaffName = staff.Name
		newStaffInfo.JobNum = staff.JobNum

		newStaffInfos = append(newStaffInfos, newStaffInfo)
	}
	return
}

func (WcStaffInfoResponse) AssembleStaffInfo(staffInfo weChat.WcStaffInfo) (newStaffInfo WcStaffInfoResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffInfo.WcStaffInfo = staffInfo
	staffInfoTypeText, _ := utils.Find(configInfo.StaffInfoType, *staffInfo.Type)
	newStaffInfo.TypeText = staffInfoTypeText
	staffOfficeRankText, _ := utils.Find(configInfo.StaffOfficeRank, *staffInfo.Rank)
	newStaffInfo.RankText = staffOfficeRankText
	householdTypeText, _ := utils.Find(configInfo.HouseholdType, *staffInfo.HouseholdType)
	newStaffInfo.HouseholdTypeText = householdTypeText
	nationText, _ := utils.Find(configInfo.Nation, *staffInfo.Nation)
	newStaffInfo.NationText = nationText
	marriageText, _ := utils.Find(configInfo.Marriage, *staffInfo.Marriage)
	newStaffInfo.MarriageText = marriageText
	politicalOutlookText, _ := utils.Find(configInfo.PoliticalOutlook, *staffInfo.PoliticalOutlook)
	newStaffInfo.PoliticalOutlookText = politicalOutlookText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffInfo.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffInfo Err:", err)
		return
	}

	newStaffInfo.StaffName = staff.Name
	newStaffInfo.JobNum = staff.JobNum

	return
}
