package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffPassApplicationRouter struct {
}

// InitWcStaffPassApplicationRouter 初始化 转正申请 路由信息
func (s *WcStaffPassApplicationRouter) InitWcStaffPassApplicationRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffPassApplicationRouter := Router.Group("wcStaffPassApplication").Use(middleware.OperationRecord())
	wcStaffPassApplicationRouterWithoutRecord := Router.Group("wcStaffPassApplication")
	wcStaffPassApplicationRouterWithoutAuth := PublicRouter.Group("wcStaffPassApplication")

	var wcStaffPassApplicationApi = v1.ApiGroupApp.EmployeeApiGroup.WcStaffPassApplicationApi
	{
		wcStaffPassApplicationRouter.POST("createWcStaffPassApplication", wcStaffPassApplicationApi.CreateWcStaffPassApplication)   // 新建转正申请
		wcStaffPassApplicationRouter.DELETE("deleteWcStaffPassApplication", wcStaffPassApplicationApi.DeleteWcStaffPassApplication) // 删除转正申请
		wcStaffPassApplicationRouter.DELETE("deleteWcStaffPassApplicationByIds", wcStaffPassApplicationApi.DeleteWcStaffPassApplicationByIds) // 批量删除转正申请
		wcStaffPassApplicationRouter.PUT("updateWcStaffPassApplication", wcStaffPassApplicationApi.UpdateWcStaffPassApplication)    // 更新转正申请
	}
	{
		wcStaffPassApplicationRouterWithoutRecord.GET("findWcStaffPassApplication", wcStaffPassApplicationApi.FindWcStaffPassApplication)        // 根据ID获取转正申请
		wcStaffPassApplicationRouterWithoutRecord.GET("getWcStaffPassApplicationList", wcStaffPassApplicationApi.GetWcStaffPassApplicationList)  // 获取转正申请列表
	}
	{
	    wcStaffPassApplicationRouterWithoutAuth.GET("getWcStaffPassApplicationPublic", wcStaffPassApplicationApi.GetWcStaffPassApplicationPublic)  // 获取转正申请列表
	}
}
