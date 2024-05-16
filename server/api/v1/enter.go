package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/employee"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/weChat"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	WeChatApiGroup   weChat.ApiGroup
	OaApiGroup       oa.ApiGroup
	EmployeeApiGroup employee.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
