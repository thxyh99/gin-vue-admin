package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"time"
)

type WcStaffJobSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

// WcStaffJobRequest 工作信息 结构体
type WcStaffJobRequest struct {
	weChat.WcStaffJob
	PositionIds   []int `json:"positionIds"`   //职务信息
	DepartmentIds []int `json:"departmentIds"` //部门信息
}
