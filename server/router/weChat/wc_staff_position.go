package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffPositionRouter struct {
}

// InitWcStaffPositionRouter 初始化 员工职位管理 路由信息
func (s *WcStaffPositionRouter) InitWcStaffPositionRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffPositionRouter := Router.Group("wcStaffPosition").Use(middleware.OperationRecord())
	wcStaffPositionRouterWithoutRecord := Router.Group("wcStaffPosition")
	wcStaffPositionRouterWithoutAuth := PublicRouter.Group("wcStaffPosition")

	var wcStaffPositionApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffPositionApi
	{
		wcStaffPositionRouter.POST("createWcStaffPosition", wcStaffPositionApi.CreateWcStaffPosition)   // 新建员工职位管理
		wcStaffPositionRouter.DELETE("deleteWcStaffPosition", wcStaffPositionApi.DeleteWcStaffPosition) // 删除员工职位管理
		wcStaffPositionRouter.DELETE("deleteWcStaffPositionByIds", wcStaffPositionApi.DeleteWcStaffPositionByIds) // 批量删除员工职位管理
		wcStaffPositionRouter.PUT("updateWcStaffPosition", wcStaffPositionApi.UpdateWcStaffPosition)    // 更新员工职位管理
	}
	{
		wcStaffPositionRouterWithoutRecord.GET("findWcStaffPosition", wcStaffPositionApi.FindWcStaffPosition)        // 根据ID获取员工职位管理
		wcStaffPositionRouterWithoutRecord.GET("getWcStaffPositionList", wcStaffPositionApi.GetWcStaffPositionList)  // 获取员工职位管理列表
	}
	{
	    wcStaffPositionRouterWithoutAuth.GET("getWcStaffPositionPublic", wcStaffPositionApi.GetWcStaffPositionPublic)  // 获取员工职位管理列表
	}
}
