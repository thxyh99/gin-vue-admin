package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffAjustlevelApplicationRouter struct {
}

// InitWcStaffAjustlevelApplicationRouter 初始化 调级申请 路由信息
func (s *WcStaffAjustlevelApplicationRouter) InitWcStaffAjustlevelApplicationRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffAjustlevelApplicationRouter := Router.Group("wcStaffAjustlevelApplication").Use(middleware.OperationRecord())
	wcStaffAjustlevelApplicationRouterWithoutRecord := Router.Group("wcStaffAjustlevelApplication")
	wcStaffAjustlevelApplicationRouterWithoutAuth := PublicRouter.Group("wcStaffAjustlevelApplication")

	var wcStaffAjustlevelApplicationApi = v1.ApiGroupApp.EmployeeApiGroup.WcStaffAjustlevelApplicationApi
	{
		wcStaffAjustlevelApplicationRouter.POST("createWcStaffAjustlevelApplication", wcStaffAjustlevelApplicationApi.CreateWcStaffAjustlevelApplication)   // 新建调级申请
		wcStaffAjustlevelApplicationRouter.DELETE("deleteWcStaffAjustlevelApplication", wcStaffAjustlevelApplicationApi.DeleteWcStaffAjustlevelApplication) // 删除调级申请
		wcStaffAjustlevelApplicationRouter.DELETE("deleteWcStaffAjustlevelApplicationByIds", wcStaffAjustlevelApplicationApi.DeleteWcStaffAjustlevelApplicationByIds) // 批量删除调级申请
		wcStaffAjustlevelApplicationRouter.PUT("updateWcStaffAjustlevelApplication", wcStaffAjustlevelApplicationApi.UpdateWcStaffAjustlevelApplication)    // 更新调级申请
	}
	{
		wcStaffAjustlevelApplicationRouterWithoutRecord.GET("findWcStaffAjustlevelApplication", wcStaffAjustlevelApplicationApi.FindWcStaffAjustlevelApplication)        // 根据ID获取调级申请
		wcStaffAjustlevelApplicationRouterWithoutRecord.GET("getWcStaffAjustlevelApplicationList", wcStaffAjustlevelApplicationApi.GetWcStaffAjustlevelApplicationList)  // 获取调级申请列表
	}
	{
	    wcStaffAjustlevelApplicationRouterWithoutAuth.GET("getWcStaffAjustlevelApplicationPublic", wcStaffAjustlevelApplicationApi.GetWcStaffAjustlevelApplicationPublic)  // 获取调级申请列表
	}
}
