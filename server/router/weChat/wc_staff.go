package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffRouter struct {
}

// InitWcStaffRouter 初始化 账号信息 路由信息
func (s *WcStaffRouter) InitWcStaffRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wcStaffRouter := Router.Group("wcStaff").Use(middleware.OperationRecord())
	wcStaffRouterWithoutRecord := Router.Group("wcStaff")
	wcStaffRouterWithoutAuth := PublicRouter.Group("wcStaff")

	var wcStaffApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffApi
	{
		wcStaffRouter.POST("createWcStaff", wcStaffApi.CreateWcStaff)             // 新建账号信息
		wcStaffRouter.DELETE("deleteWcStaff", wcStaffApi.DeleteWcStaff)           // 删除账号信息
		wcStaffRouter.DELETE("deleteWcStaffByIds", wcStaffApi.DeleteWcStaffByIds) // 批量删除账号信息
		wcStaffRouter.PUT("updateWcStaff", wcStaffApi.UpdateWcStaff)              // 更新账号信息
		wcStaffRouter.POST("importExcel", wcStaffApi.ImportExcel)                 // 导入账号信息
	}
	{
		wcStaffRouterWithoutRecord.GET("findWcStaff", wcStaffApi.FindWcStaff)                   // 根据ID获取账号信息
		wcStaffRouterWithoutRecord.GET("getWcStaffList", wcStaffApi.GetWcStaffList)             // 获取账号信息列表
		wcStaffRouterWithoutRecord.GET("getSimpleStaffList", wcStaffApi.GetSimpleStaffList)     // 获取账号信息列表
		wcStaffRouterWithoutRecord.GET("obtainEmployeeRoster", wcStaffApi.ObtainEmployeeRoster) // 获取员工花名册
	}
	{
		wcStaffRouterWithoutAuth.GET("getWcStaffPublic", wcStaffApi.GetWcStaffPublic) // 获取账号信息列表
	}
}
