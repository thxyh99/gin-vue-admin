package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffTransferApplicationRouter struct {
}

// InitWcStaffTransferApplicationRouter 初始化 调动申请 路由信息
func (s *WcStaffTransferApplicationRouter) InitWcStaffTransferApplicationRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffTransferApplicationRouter := Router.Group("wcStaffTransferApplication").Use(middleware.OperationRecord())
	wcStaffTransferApplicationRouterWithoutRecord := Router.Group("wcStaffTransferApplication")
	wcStaffTransferApplicationRouterWithoutAuth := PublicRouter.Group("wcStaffTransferApplication")

	var wcStaffTransferApplicationApi = v1.ApiGroupApp.EmployeeApiGroup.WcStaffTransferApplicationApi
	{
		wcStaffTransferApplicationRouter.POST("createWcStaffTransferApplication", wcStaffTransferApplicationApi.CreateWcStaffTransferApplication)   // 新建调动申请
		wcStaffTransferApplicationRouter.DELETE("deleteWcStaffTransferApplication", wcStaffTransferApplicationApi.DeleteWcStaffTransferApplication) // 删除调动申请
		wcStaffTransferApplicationRouter.DELETE("deleteWcStaffTransferApplicationByIds", wcStaffTransferApplicationApi.DeleteWcStaffTransferApplicationByIds) // 批量删除调动申请
		wcStaffTransferApplicationRouter.PUT("updateWcStaffTransferApplication", wcStaffTransferApplicationApi.UpdateWcStaffTransferApplication)    // 更新调动申请
	}
	{
		wcStaffTransferApplicationRouterWithoutRecord.GET("findWcStaffTransferApplication", wcStaffTransferApplicationApi.FindWcStaffTransferApplication)        // 根据ID获取调动申请
		wcStaffTransferApplicationRouterWithoutRecord.GET("getWcStaffTransferApplicationList", wcStaffTransferApplicationApi.GetWcStaffTransferApplicationList)  // 获取调动申请列表
	}
	{
	    wcStaffTransferApplicationRouterWithoutAuth.GET("getWcStaffTransferApplicationPublic", wcStaffTransferApplicationApi.GetWcStaffTransferApplicationPublic)  // 获取调动申请列表
	}
}
