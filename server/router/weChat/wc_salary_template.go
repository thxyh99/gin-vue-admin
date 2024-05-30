package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcSalaryTemplateRouter struct {
}

// InitWcSalaryTemplateRouter 初始化 工资单模版 路由信息
func (s *WcSalaryTemplateRouter) InitWcSalaryTemplateRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcSalaryTemplateRouter := Router.Group("wcSalaryTemplate").Use(middleware.OperationRecord())
	wcSalaryTemplateRouterWithoutRecord := Router.Group("wcSalaryTemplate")
	wcSalaryTemplateRouterWithoutAuth := PublicRouter.Group("wcSalaryTemplate")

	var wcSalaryTemplateApi = v1.ApiGroupApp.WeChatApiGroup.WcSalaryTemplateApi
	{
		wcSalaryTemplateRouter.POST("createWcSalaryTemplate", wcSalaryTemplateApi.CreateWcSalaryTemplate)   // 新建工资单模版
		wcSalaryTemplateRouter.DELETE("deleteWcSalaryTemplate", wcSalaryTemplateApi.DeleteWcSalaryTemplate) // 删除工资单模版
		wcSalaryTemplateRouter.DELETE("deleteWcSalaryTemplateByIds", wcSalaryTemplateApi.DeleteWcSalaryTemplateByIds) // 批量删除工资单模版
		wcSalaryTemplateRouter.PUT("updateWcSalaryTemplate", wcSalaryTemplateApi.UpdateWcSalaryTemplate)    // 更新工资单模版
	}
	{
		wcSalaryTemplateRouterWithoutRecord.GET("findWcSalaryTemplate", wcSalaryTemplateApi.FindWcSalaryTemplate)        // 根据ID获取工资单模版
		wcSalaryTemplateRouterWithoutRecord.GET("getWcSalaryTemplateList", wcSalaryTemplateApi.GetWcSalaryTemplateList)  // 获取工资单模版列表
	}
	{
	    wcSalaryTemplateRouterWithoutAuth.GET("getWcSalaryTemplatePublic", wcSalaryTemplateApi.GetWcSalaryTemplatePublic)  // 获取工资单模版列表
	}
}
