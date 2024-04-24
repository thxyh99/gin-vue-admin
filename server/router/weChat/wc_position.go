package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcPositionRouter struct {
}

// InitWcPositionRouter 初始化 职位管理 路由信息
func (s *WcPositionRouter) InitWcPositionRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcPositionRouter := Router.Group("wcPosition").Use(middleware.OperationRecord())
	wcPositionRouterWithoutRecord := Router.Group("wcPosition")
	wcPositionRouterWithoutAuth := PublicRouter.Group("wcPosition")

	var wcPositionApi = v1.ApiGroupApp.WeChatApiGroup.WcPositionApi
	{
		wcPositionRouter.POST("createWcPosition", wcPositionApi.CreateWcPosition)   // 新建职位管理
		wcPositionRouter.DELETE("deleteWcPosition", wcPositionApi.DeleteWcPosition) // 删除职位管理
		wcPositionRouter.DELETE("deleteWcPositionByIds", wcPositionApi.DeleteWcPositionByIds) // 批量删除职位管理
		wcPositionRouter.PUT("updateWcPosition", wcPositionApi.UpdateWcPosition)    // 更新职位管理
	}
	{
		wcPositionRouterWithoutRecord.GET("findWcPosition", wcPositionApi.FindWcPosition)        // 根据ID获取职位管理
		wcPositionRouterWithoutRecord.GET("getWcPositionList", wcPositionApi.GetWcPositionList)  // 获取职位管理列表
	}
	{
	    wcPositionRouterWithoutAuth.GET("getWcPositionPublic", wcPositionApi.GetWcPositionPublic)  // 获取职位管理列表
	}
}
