package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffDepartmentRouter struct {
}

// InitWcStaffDepartmentRouter 初始化 员工部门管理 路由信息
func (s *WcStaffDepartmentRouter) InitWcStaffDepartmentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffDepartmentRouter := Router.Group("wcStaffDepartment").Use(middleware.OperationRecord())
	wcStaffDepartmentRouterWithoutRecord := Router.Group("wcStaffDepartment")
	wcStaffDepartmentRouterWithoutAuth := PublicRouter.Group("wcStaffDepartment")

	var wcStaffDepartmentApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffDepartmentApi
	{
		wcStaffDepartmentRouter.POST("createWcStaffDepartment", wcStaffDepartmentApi.CreateWcStaffDepartment)   // 新建员工部门管理
		wcStaffDepartmentRouter.DELETE("deleteWcStaffDepartment", wcStaffDepartmentApi.DeleteWcStaffDepartment) // 删除员工部门管理
		wcStaffDepartmentRouter.DELETE("deleteWcStaffDepartmentByIds", wcStaffDepartmentApi.DeleteWcStaffDepartmentByIds) // 批量删除员工部门管理
		wcStaffDepartmentRouter.PUT("updateWcStaffDepartment", wcStaffDepartmentApi.UpdateWcStaffDepartment)    // 更新员工部门管理
	}
	{
		wcStaffDepartmentRouterWithoutRecord.GET("findWcStaffDepartment", wcStaffDepartmentApi.FindWcStaffDepartment)        // 根据ID获取员工部门管理
		wcStaffDepartmentRouterWithoutRecord.GET("getWcStaffDepartmentList", wcStaffDepartmentApi.GetWcStaffDepartmentList)  // 获取员工部门管理列表
	}
	{
	    wcStaffDepartmentRouterWithoutAuth.GET("getWcStaffDepartmentPublic", wcStaffDepartmentApi.GetWcStaffDepartmentPublic)  // 获取员工部门管理列表
	}
}
