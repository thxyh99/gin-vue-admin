package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffEmploymentApplicationRouter struct {
}

// InitWcStaffEmploymentApplicationRouter 初始化 入职申请 路由信息
func (s *WcStaffEmploymentApplicationRouter) InitWcStaffEmploymentApplicationRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffEmploymentApplicationRouter := Router.Group("wcStaffEmploymentApplication").Use(middleware.OperationRecord())
	wcStaffEmploymentApplicationRouterWithoutRecord := Router.Group("wcStaffEmploymentApplication")
	wcStaffEmploymentApplicationRouterWithoutAuth := PublicRouter.Group("wcStaffEmploymentApplication")

	var wcStaffEmploymentApplicationApi = v1.ApiGroupApp.EmployeeApiGroup.WcStaffEmploymentApplicationApi
	{
		wcStaffEmploymentApplicationRouter.POST("createWcStaffEmploymentApplication", wcStaffEmploymentApplicationApi.CreateWcStaffEmploymentApplication)   // 新建入职申请
		wcStaffEmploymentApplicationRouter.DELETE("deleteWcStaffEmploymentApplication", wcStaffEmploymentApplicationApi.DeleteWcStaffEmploymentApplication) // 删除入职申请
		wcStaffEmploymentApplicationRouter.DELETE("deleteWcStaffEmploymentApplicationByIds", wcStaffEmploymentApplicationApi.DeleteWcStaffEmploymentApplicationByIds) // 批量删除入职申请
		wcStaffEmploymentApplicationRouter.PUT("updateWcStaffEmploymentApplication", wcStaffEmploymentApplicationApi.UpdateWcStaffEmploymentApplication)    // 更新入职申请
	}
	{
		wcStaffEmploymentApplicationRouterWithoutRecord.GET("findWcStaffEmploymentApplication", wcStaffEmploymentApplicationApi.FindWcStaffEmploymentApplication)        // 根据ID获取入职申请
		wcStaffEmploymentApplicationRouterWithoutRecord.GET("getWcStaffEmploymentApplicationList", wcStaffEmploymentApplicationApi.GetWcStaffEmploymentApplicationList)  // 获取入职申请列表
	}
	{
	    wcStaffEmploymentApplicationRouterWithoutAuth.GET("getWcStaffEmploymentApplicationPublic", wcStaffEmploymentApplicationApi.GetWcStaffEmploymentApplicationPublic)  // 获取入职申请列表
	}
}
