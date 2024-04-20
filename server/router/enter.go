package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/weChat"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	WeChat  weChat.RouterGroup
	Oa      oa.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
