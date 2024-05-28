package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WcStaffSocialSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StaffId        *int       `json:"staffId" form:"staffId"`
	Type           *int       `json:"type" form:"type"`

	request.PageInfo
}
