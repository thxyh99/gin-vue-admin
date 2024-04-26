package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffMaterialsRouter struct {
}

// InitWcStaffMaterialsRouter 初始化 证件材料 路由信息
func (s *WcStaffMaterialsRouter) InitWcStaffMaterialsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffMaterialsRouter := Router.Group("wcStaffMaterials").Use(middleware.OperationRecord())
	wcStaffMaterialsRouterWithoutRecord := Router.Group("wcStaffMaterials")
	wcStaffMaterialsRouterWithoutAuth := PublicRouter.Group("wcStaffMaterials")

	var wcStaffMaterialsApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffMaterialsApi
	{
		wcStaffMaterialsRouter.POST("createWcStaffMaterials", wcStaffMaterialsApi.CreateWcStaffMaterials)   // 新建证件材料
		wcStaffMaterialsRouter.DELETE("deleteWcStaffMaterials", wcStaffMaterialsApi.DeleteWcStaffMaterials) // 删除证件材料
		wcStaffMaterialsRouter.DELETE("deleteWcStaffMaterialsByIds", wcStaffMaterialsApi.DeleteWcStaffMaterialsByIds) // 批量删除证件材料
		wcStaffMaterialsRouter.PUT("updateWcStaffMaterials", wcStaffMaterialsApi.UpdateWcStaffMaterials)    // 更新证件材料
	}
	{
		wcStaffMaterialsRouterWithoutRecord.GET("findWcStaffMaterials", wcStaffMaterialsApi.FindWcStaffMaterials)        // 根据ID获取证件材料
		wcStaffMaterialsRouterWithoutRecord.GET("getWcStaffMaterialsList", wcStaffMaterialsApi.GetWcStaffMaterialsList)  // 获取证件材料列表
	}
	{
	    wcStaffMaterialsRouterWithoutAuth.GET("getWcStaffMaterialsPublic", wcStaffMaterialsApi.GetWcStaffMaterialsPublic)  // 获取证件材料列表
	}
}
