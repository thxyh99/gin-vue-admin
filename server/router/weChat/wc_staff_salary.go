package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffSalaryRouter struct {
}

// InitWcStaffSalaryRouter 初始化 工资单 路由信息
func (s *WcStaffSalaryRouter) InitWcStaffSalaryRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wcStaffSalaryRouter := Router.Group("wcStaffSalary").Use(middleware.OperationRecord())
	wcStaffSalaryRouterWithoutRecord := Router.Group("wcStaffSalary")
	wcStaffSalaryRouterWithoutAuth := PublicRouter.Group("wcStaffSalary")

	var wcStaffSalaryApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffSalaryApi
	{
		wcStaffSalaryRouter.POST("createWcStaffSalary", wcStaffSalaryApi.CreateWcStaffSalary)             // 新建工资单
		wcStaffSalaryRouter.DELETE("deleteWcStaffSalary", wcStaffSalaryApi.DeleteWcStaffSalary)           // 删除工资单
		wcStaffSalaryRouter.DELETE("deleteWcStaffSalaryByIds", wcStaffSalaryApi.DeleteWcStaffSalaryByIds) // 批量删除工资单
		wcStaffSalaryRouter.PUT("updateWcStaffSalary", wcStaffSalaryApi.UpdateWcStaffSalary)              // 更新工资单
	}
	{
		wcStaffSalaryRouterWithoutRecord.GET("findWcStaffSalary", wcStaffSalaryApi.FindWcStaffSalary)       // 根据ID获取工资单
		wcStaffSalaryRouterWithoutRecord.GET("getWcStaffSalaryList", wcStaffSalaryApi.GetWcStaffSalaryList) // 获取工资单列表
	}
	{
		wcStaffSalaryRouterWithoutAuth.GET("getWcStaffSalaryPublic", wcStaffSalaryApi.GetWcStaffSalaryPublic) // 获取工资单列表
	}
}
