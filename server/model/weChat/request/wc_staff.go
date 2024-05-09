package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"time"
)

type WcStaffSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

// WcStaffRequest 账号信息 结构体
type WcStaffRequest struct {
	weChat.WcStaff
	PositionIds   []int `json:"positionIds"`   //职务信息
	DepartmentIds []int `json:"departmentIds"` //部门信息
}
