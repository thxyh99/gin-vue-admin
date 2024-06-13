package employee

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffAdjustlevelApplicationRouter struct {
}

// InitWcStaffAdjustlevelApplicationRouter 初始化 调级申请 路由信息
func (s *WcStaffAdjustlevelApplicationRouter) InitWcStaffAdjustlevelApplicationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wcStaffAdjustlevelApplicationRouter := Router.Group("wcStaffAdjustlevelApplication").Use(middleware.OperationRecord())
	wcStaffAdjustlevelApplicationRouterWithoutRecord := Router.Group("wcStaffAdjustlevelApplication")
	wcStaffAdjustlevelApplicationRouterWithoutAuth := PublicRouter.Group("wcStaffAdjustlevelApplication")

	var wcStaffAdjustlevelApplicationApi = v1.ApiGroupApp.EmployeeApiGroup.WcStaffAdjustlevelApplicationApi
	{
		wcStaffAdjustlevelApplicationRouter.POST("createWcStaffAdjustlevelApplication", wcStaffAdjustlevelApplicationApi.CreateWcStaffAdjustlevelApplication)             // 新建调级申请
		wcStaffAdjustlevelApplicationRouter.DELETE("deleteWcStaffAdjustlevelApplication", wcStaffAdjustlevelApplicationApi.DeleteWcStaffAdjustlevelApplication)           // 删除调级申请
		wcStaffAdjustlevelApplicationRouter.DELETE("deleteWcStaffAdjustlevelApplicationByIds", wcStaffAdjustlevelApplicationApi.DeleteWcStaffAdjustlevelApplicationByIds) // 批量删除调级申请
		wcStaffAdjustlevelApplicationRouter.PUT("updateWcStaffAdjustlevelApplication", wcStaffAdjustlevelApplicationApi.UpdateWcStaffAdjustlevelApplication)              // 更新调级申请
	}
	{
		wcStaffAdjustlevelApplicationRouterWithoutRecord.GET("findWcStaffAdjustlevelApplication", wcStaffAdjustlevelApplicationApi.FindWcStaffAdjustlevelApplication)       // 根据ID获取调级申请
		wcStaffAdjustlevelApplicationRouterWithoutRecord.GET("getWcStaffAdjustlevelApplicationList", wcStaffAdjustlevelApplicationApi.GetWcStaffAdjustlevelApplicationList) // 获取调级申请列表
	}
	{
		wcStaffAdjustlevelApplicationRouterWithoutAuth.GET("getWcStaffAdjustlevelApplicationPublic", wcStaffAdjustlevelApplicationApi.GetWcStaffAdjustlevelApplicationPublic) // 获取调级申请列表
		wcStaffAdjustlevelApplicationRouterWithoutAuth.GET("createOAStaffAdjustlevelApplication", wcStaffAdjustlevelApplicationApi.CreateOAStaffAdjustlevelApplication)       // 获取调级申请列表
	}
}
