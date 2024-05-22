package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffResponse 账号信息 结构体
type WcStaffResponse struct {
	weChat.WcStaff
	GenderText           string `json:"genderText"`            //性别信息
	HouseholdTypeText    string `json:"householdTypeText"`     //户籍类型
	NationText           string `json:"nationText"`            //员工民族
	MarriageText         string `json:"marriageText"`          //员工婚否
	PoliticalOutlookText string `json:"politicalOutlookText" ` //员工政治面貌
}
