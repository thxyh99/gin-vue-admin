package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WcStaffSalarySearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Month          string     `json:"month" form:"month"`
	StaffId        *int       `json:"staffId" form:"staffId"`

	request.PageInfo
}
