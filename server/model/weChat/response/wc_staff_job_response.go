package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcStaffJobResponse 工作信息 结构体
type WcStaffJobResponse struct {
	weChat.WcStaffJob
	StaffName          string `json:"staffName"`          //员工名称
	JobNum             string `json:"jobNum"`             //员工工号
	TypeText           string `json:"typeText"`           //员工类型
	StatusText         string `json:"statusText"`         //员工状态
	TryPeriodText      string `json:"tryPeriodText"`      //试用期
	ExpenseAccountText string `json:"expenseAccountText"` //费用科目
	Position           string `json:"position"`           //岗位信息
	PositionIds        []int  `json:"positionIds"`        //岗位信息编号
	Department         string `json:"department"`         //部门信息
	DepartmentIds      []int  `json:"departmentIds"`      //部门信息编号
	RankTypeText       string `json:"rankTypeText"`       //职级类型
	RankText           string `json:"rankText"`           //职级
	Leader             string `json:"leader"`             //直属领导
}
