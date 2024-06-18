package employee

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffLeaveApplicationRouter struct {
}

// InitWcStaffLeaveApplicationRouter 初始化 离职申请 路由信息
func (s *WcStaffLeaveApplicationRouter) InitWcStaffLeaveApplicationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wcStaffLeaveApplicationRouter := Router.Group("wcStaffLeaveApplication").Use(middleware.OperationRecord())
	wcStaffLeaveApplicationRouterWithoutRecord := Router.Group("wcStaffLeaveApplication")
	wcStaffLeaveApplicationRouterWithoutAuth := PublicRouter.Group("wcStaffLeaveApplication")

	var wcStaffLeaveApplicationApi = v1.ApiGroupApp.EmployeeApiGroup.WcStaffLeaveApplicationApi
	{
		wcStaffLeaveApplicationRouter.POST("createWcStaffLeaveApplication", wcStaffLeaveApplicationApi.CreateWcStaffLeaveApplication)             // 新建离职申请
		wcStaffLeaveApplicationRouter.DELETE("deleteWcStaffLeaveApplication", wcStaffLeaveApplicationApi.DeleteWcStaffLeaveApplication)           // 删除离职申请
		wcStaffLeaveApplicationRouter.DELETE("deleteWcStaffLeaveApplicationByIds", wcStaffLeaveApplicationApi.DeleteWcStaffLeaveApplicationByIds) // 批量删除离职申请
		wcStaffLeaveApplicationRouter.PUT("updateWcStaffLeaveApplication", wcStaffLeaveApplicationApi.UpdateWcStaffLeaveApplication)              // 更新离职申请
	}
	{
		wcStaffLeaveApplicationRouterWithoutRecord.GET("findWcStaffLeaveApplication", wcStaffLeaveApplicationApi.FindWcStaffLeaveApplication) // 根据ID获取离职申请
		wcStaffLeaveApplicationRouterWithoutRecord.GET("getWcStaffLeaveApplicationList", wcStaffLeaveApplicationApi.GetWcStaffLeaveApplicationList)
		// 获取离职申请列表
	}
	{
		wcStaffLeaveApplicationRouterWithoutAuth.GET("getWcStaffLeaveApplicationPublic", wcStaffLeaveApplicationApi.GetWcStaffLeaveApplicationPublic) // 获取离职申请列表
		wcStaffLeaveApplicationRouterWithoutAuth.POST("createOAStaffLeaveApplication", wcStaffLeaveApplicationApi.CreateOAStaffLeaveApplication)      // 创建OA离职申请
	}
}
