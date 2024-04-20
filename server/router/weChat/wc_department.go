package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcDepartmentRouter struct {
}

// InitWcDepartmentRouter 初始化 wcDepartment表 路由信息
func (s *WcDepartmentRouter) InitWcDepartmentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcDepartmentRouter := Router.Group("wcDepartment").Use(middleware.OperationRecord())
	wcDepartmentRouterWithoutRecord := Router.Group("wcDepartment")
	wcDepartmentRouterWithoutAuth := PublicRouter.Group("wcDepartment")

	var wcDepartmentApi = v1.ApiGroupApp.WeChatApiGroup.WcDepartmentApi
	{
		wcDepartmentRouter.POST("createWcDepartment", wcDepartmentApi.CreateWcDepartment)   // 新建wcDepartment表
		wcDepartmentRouter.DELETE("deleteWcDepartment", wcDepartmentApi.DeleteWcDepartment) // 删除wcDepartment表
		wcDepartmentRouter.DELETE("deleteWcDepartmentByIds", wcDepartmentApi.DeleteWcDepartmentByIds) // 批量删除wcDepartment表
		wcDepartmentRouter.PUT("updateWcDepartment", wcDepartmentApi.UpdateWcDepartment)    // 更新wcDepartment表
	}
	{
		wcDepartmentRouterWithoutRecord.GET("findWcDepartment", wcDepartmentApi.FindWcDepartment)        // 根据ID获取wcDepartment表
		wcDepartmentRouterWithoutRecord.GET("getWcDepartmentList", wcDepartmentApi.GetWcDepartmentList)  // 获取wcDepartment表列表
	}
	{
	    wcDepartmentRouterWithoutAuth.GET("getWcDepartmentPublic", wcDepartmentApi.GetWcDepartmentPublic)  // 获取wcDepartment表列表
	}
}
