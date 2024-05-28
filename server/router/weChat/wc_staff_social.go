package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffSocialRouter struct {
}

// InitWcStaffSocialRouter 初始化 社保公积金管理 路由信息
func (s *WcStaffSocialRouter) InitWcStaffSocialRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wcStaffSocialRouter := Router.Group("wcStaffSocial").Use(middleware.OperationRecord())
	wcStaffSocialRouterWithoutRecord := Router.Group("wcStaffSocial")
	wcStaffSocialRouterWithoutAuth := PublicRouter.Group("wcStaffSocial")

	var wcStaffSocialApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffSocialApi
	{
		wcStaffSocialRouter.POST("createWcStaffSocial", wcStaffSocialApi.CreateWcStaffSocial)             // 新建社保公积金管理
		wcStaffSocialRouter.DELETE("deleteWcStaffSocial", wcStaffSocialApi.DeleteWcStaffSocial)           // 删除社保公积金管理
		wcStaffSocialRouter.DELETE("deleteWcStaffSocialByIds", wcStaffSocialApi.DeleteWcStaffSocialByIds) // 批量删除社保公积金管理
		wcStaffSocialRouter.PUT("updateWcStaffSocial", wcStaffSocialApi.UpdateWcStaffSocial)              // 更新社保公积金管理
		wcStaffSocialRouter.POST("importExcel", wcStaffSocialApi.ImportExcel)                             // 导入账号信息
	}
	{
		wcStaffSocialRouterWithoutRecord.GET("findWcStaffSocial", wcStaffSocialApi.FindWcStaffSocial)       // 根据ID获取社保公积金管理
		wcStaffSocialRouterWithoutRecord.GET("getWcStaffSocialList", wcStaffSocialApi.GetWcStaffSocialList) // 获取社保公积金管理列表
	}
	{
		wcStaffSocialRouterWithoutAuth.GET("getWcStaffSocialPublic", wcStaffSocialApi.GetWcStaffSocialPublic) // 获取社保公积金管理列表
	}
}
