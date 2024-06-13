package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"time"
)

type WcStaffSearch struct {
	StartCreatedAt      *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt        *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StaffId             *int       `json:"staffId" form:"staffId"`
	EmploymentDateStart string     `json:"employmentDateStart" form:"employmentDateStart"`
	EmploymentDateEnd   string     `json:"employmentDateEnd" form:"employmentDateEnd"`
	HistoryDate         string     `json:"historyDate" form:"historyDate"`
	DepartmentIds       []int      `json:"departmentIds[]" form:"departmentIds[]"`

	request.PageInfo
}

// WcStaffRequest 账号信息 结构体
type WcStaffRequest struct {
	weChat.WcStaff
	PositionIds   []int `json:"positionIds"`   //职务信息
	DepartmentIds []int `json:"departmentIds"` //部门信息
}
