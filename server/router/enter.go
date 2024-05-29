package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/employee"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/weChat"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	WeChat   weChat.RouterGroup
	Employee employee.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
