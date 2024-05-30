package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcSalaryRouter struct {
}

// InitWcSalaryRouter 初始化 工资单模版 路由信息
func (s *WcSalaryRouter) InitWcSalaryRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcSalaryRouter := Router.Group("wcSalary").Use(middleware.OperationRecord())
	wcSalaryRouterWithoutRecord := Router.Group("wcSalary")
	wcSalaryRouterWithoutAuth := PublicRouter.Group("wcSalary")

	var wcSalaryApi = v1.ApiGroupApp.WeChatApiGroup.WcSalaryApi
	{
		wcSalaryRouter.POST("createWcSalary", wcSalaryApi.CreateWcSalary)   // 新建工资单模版
		wcSalaryRouter.DELETE("deleteWcSalary", wcSalaryApi.DeleteWcSalary) // 删除工资单模版
		wcSalaryRouter.DELETE("deleteWcSalaryByIds", wcSalaryApi.DeleteWcSalaryByIds) // 批量删除工资单模版
		wcSalaryRouter.PUT("updateWcSalary", wcSalaryApi.UpdateWcSalary)    // 更新工资单模版
	}
	{
		wcSalaryRouterWithoutRecord.GET("findWcSalary", wcSalaryApi.FindWcSalary)        // 根据ID获取工资单模版
		wcSalaryRouterWithoutRecord.GET("getWcSalaryList", wcSalaryApi.GetWcSalaryList)  // 获取工资单模版列表
	}
	{
	    wcSalaryRouterWithoutAuth.GET("getWcSalaryPublic", wcSalaryApi.GetWcSalaryPublic)  // 获取工资单模版列表
	}
}
