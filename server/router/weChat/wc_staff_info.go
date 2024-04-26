package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffInfoRouter struct {
}

// InitWcStaffInfoRouter 初始化 个人信息 路由信息
func (s *WcStaffInfoRouter) InitWcStaffInfoRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffInfoRouter := Router.Group("wcStaffInfo").Use(middleware.OperationRecord())
	wcStaffInfoRouterWithoutRecord := Router.Group("wcStaffInfo")
	wcStaffInfoRouterWithoutAuth := PublicRouter.Group("wcStaffInfo")

	var wcStaffInfoApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffInfoApi
	{
		wcStaffInfoRouter.POST("createWcStaffInfo", wcStaffInfoApi.CreateWcStaffInfo)   // 新建个人信息
		wcStaffInfoRouter.DELETE("deleteWcStaffInfo", wcStaffInfoApi.DeleteWcStaffInfo) // 删除个人信息
		wcStaffInfoRouter.DELETE("deleteWcStaffInfoByIds", wcStaffInfoApi.DeleteWcStaffInfoByIds) // 批量删除个人信息
		wcStaffInfoRouter.PUT("updateWcStaffInfo", wcStaffInfoApi.UpdateWcStaffInfo)    // 更新个人信息
	}
	{
		wcStaffInfoRouterWithoutRecord.GET("findWcStaffInfo", wcStaffInfoApi.FindWcStaffInfo)        // 根据ID获取个人信息
		wcStaffInfoRouterWithoutRecord.GET("getWcStaffInfoList", wcStaffInfoApi.GetWcStaffInfoList)  // 获取个人信息列表
	}
	{
	    wcStaffInfoRouterWithoutAuth.GET("getWcStaffInfoPublic", wcStaffInfoApi.GetWcStaffInfoPublic)  // 获取个人信息列表
	}
}
