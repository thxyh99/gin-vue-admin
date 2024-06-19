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
	PaymentPlaceText     string `json:"paymentPlaceText" `     //社保公积金缴纳地
}

type WcStaffStatisticsResponse struct {
	OnJobCount        int        `json:"onJobCount"`        //在职人数
	FullTimeCount     int        `json:"fullTimeCount"`     //全职人数
	PartTimeCount     int        `json:"partTimeCount"`     //兼职人数
	FollowCount       int        `json:"followUpJobCount"`  //跟岗实习人数
	ReplaceCount      int        `json:"replaceCount"`      //顶岗实习人数
	RetireCount       int        `json:"retireCount"`       //退休返聘人数
	OutsourcingCount  int        `json:"outsourcingCount"`  //劳务外包人数
	ToBeEmployedCount int        `json:"toBeEmployedCount"` //待入职人数
	ProbationCount    int        `json:"probationCount"`    //试用人数
	FormalCount       int        `json:"formalCount"`       //正式人数
	ToBeDepartedCount int        `json:"toBeDepartedCount"` //待离职人数
	TodoCount         int        `json:"todoCount"`         //待办事项个数
	TodoList          []TodoList `json:"todoList"`          //待办事项列表
}

type TodoList struct {
	Type  int        `json:"type"` //0:全部 1:超期未处理 2:今日到期 3:提前提醒
	Total int        `json:"total"`
	Label string     `json:"label"`
	List  []TodoItem `json:"list"`
}

type TodoItem struct {
	Title   string `json:"title"`
	Label   string `json:"label"`
	Value   string `json:"value"`
	StaffId int    `json:"staffId"`
}
