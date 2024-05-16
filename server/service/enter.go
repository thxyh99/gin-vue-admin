package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/employee"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/weChat"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	WeChatServiceGroup   weChat.ServiceGroup
	OaServiceGroup       oa.ServiceGroup
	EmployeeServiceGroup employee.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
