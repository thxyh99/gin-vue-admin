package employee

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcThirdOaPushRouter struct {
}

// InitWcThirdOaPushRouter 初始化 OA回调日志 路由信息
func (s *WcThirdOaPushRouter) InitWcThirdOaPushRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wcThirdOaPushRouter := Router.Group("wcThirdOaPush").Use(middleware.OperationRecord())
	wcThirdOaPushRouterWithoutRecord := Router.Group("wcThirdOaPush")
	wcThirdOaPushRouterWithoutAuth := PublicRouter.Group("wcThirdOaPush")

	var wcThirdOaPushApi = v1.ApiGroupApp.EmployeeApiGroup.WcThirdOaPushApi
	{
		wcThirdOaPushRouter.POST("createWcThirdOaPush", wcThirdOaPushApi.CreateWcThirdOaPush)             // 更新OA回调日志
		wcThirdOaPushRouter.DELETE("deleteWcThirdOaPush", wcThirdOaPushApi.DeleteWcThirdOaPush)           // 删除OA回调日志
		wcThirdOaPushRouter.DELETE("deleteWcThirdOaPushByIds", wcThirdOaPushApi.DeleteWcThirdOaPushByIds) // 批量删除OA回调日志
		wcThirdOaPushRouter.PUT("updateWcThirdOaPush", wcThirdOaPushApi.UpdateWcThirdOaPush)              // 更新OA回调日志
	}
	{
		wcThirdOaPushRouterWithoutRecord.GET("findWcThirdOaPush", wcThirdOaPushApi.FindWcThirdOaPush)       // 根据ID获取OA回调日志
		wcThirdOaPushRouterWithoutRecord.GET("getWcThirdOaPushList", wcThirdOaPushApi.GetWcThirdOaPushList) // 获取OA回调日志列表
	}
	{
		wcThirdOaPushRouterWithoutAuth.GET("getWcThirdOaPushPublic", wcThirdOaPushApi.GetWcThirdOaPushPublic) // 获取OA回调日志列表
		wcThirdOaPushRouterWithoutAuth.POST("oaPush", wcThirdOaPushApi.CreateWcThirdOaPush)                   // 新建OA回调日志
	}
}
