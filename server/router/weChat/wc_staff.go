package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffRouter struct {
}

// InitWcStaffRouter 初始化 员工管理 路由信息
func (s *WcStaffRouter) InitWcStaffRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffRouter := Router.Group("wcStaff").Use(middleware.OperationRecord())
	wcStaffRouterWithoutRecord := Router.Group("wcStaff")
	wcStaffRouterWithoutAuth := PublicRouter.Group("wcStaff")

	var wcStaffApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffApi
	{
		wcStaffRouter.POST("createWcStaff", wcStaffApi.CreateWcStaff)   // 新建员工管理
		wcStaffRouter.DELETE("deleteWcStaff", wcStaffApi.DeleteWcStaff) // 删除员工管理
		wcStaffRouter.DELETE("deleteWcStaffByIds", wcStaffApi.DeleteWcStaffByIds) // 批量删除员工管理
		wcStaffRouter.PUT("updateWcStaff", wcStaffApi.UpdateWcStaff)    // 更新员工管理
	}
	{
		wcStaffRouterWithoutRecord.GET("findWcStaff", wcStaffApi.FindWcStaff)        // 根据ID获取员工管理
		wcStaffRouterWithoutRecord.GET("getWcStaffList", wcStaffApi.GetWcStaffList)  // 获取员工管理列表
	}
	{
	    wcStaffRouterWithoutAuth.GET("getWcStaffPublic", wcStaffApi.GetWcStaffPublic)  // 获取员工管理列表
	}
}
