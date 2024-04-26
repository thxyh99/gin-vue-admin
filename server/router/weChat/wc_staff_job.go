package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffJobRouter struct {
}

// InitWcStaffJobRouter 初始化 工作信息 路由信息
func (s *WcStaffJobRouter) InitWcStaffJobRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffJobRouter := Router.Group("wcStaffJob").Use(middleware.OperationRecord())
	wcStaffJobRouterWithoutRecord := Router.Group("wcStaffJob")
	wcStaffJobRouterWithoutAuth := PublicRouter.Group("wcStaffJob")

	var wcStaffJobApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffJobApi
	{
		wcStaffJobRouter.POST("createWcStaffJob", wcStaffJobApi.CreateWcStaffJob)   // 新建工作信息
		wcStaffJobRouter.DELETE("deleteWcStaffJob", wcStaffJobApi.DeleteWcStaffJob) // 删除工作信息
		wcStaffJobRouter.DELETE("deleteWcStaffJobByIds", wcStaffJobApi.DeleteWcStaffJobByIds) // 批量删除工作信息
		wcStaffJobRouter.PUT("updateWcStaffJob", wcStaffJobApi.UpdateWcStaffJob)    // 更新工作信息
	}
	{
		wcStaffJobRouterWithoutRecord.GET("findWcStaffJob", wcStaffJobApi.FindWcStaffJob)        // 根据ID获取工作信息
		wcStaffJobRouterWithoutRecord.GET("getWcStaffJobList", wcStaffJobApi.GetWcStaffJobList)  // 获取工作信息列表
	}
	{
	    wcStaffJobRouterWithoutAuth.GET("getWcStaffJobPublic", wcStaffJobApi.GetWcStaffJobPublic)  // 获取工作信息列表
	}
}
