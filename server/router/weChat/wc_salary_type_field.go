package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcSalaryTypeFieldRouter struct {
}

// InitWcSalaryTypeFieldRouter 初始化 工资类型款项 路由信息
func (s *WcSalaryTypeFieldRouter) InitWcSalaryTypeFieldRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcSalaryTypeFieldRouter := Router.Group("wcSalaryTypeField").Use(middleware.OperationRecord())
	wcSalaryTypeFieldRouterWithoutRecord := Router.Group("wcSalaryTypeField")
	wcSalaryTypeFieldRouterWithoutAuth := PublicRouter.Group("wcSalaryTypeField")

	var wcSalaryTypeFieldApi = v1.ApiGroupApp.WeChatApiGroup.WcSalaryTypeFieldApi
	{
		wcSalaryTypeFieldRouter.POST("createWcSalaryTypeField", wcSalaryTypeFieldApi.CreateWcSalaryTypeField)   // 新建工资类型款项
		wcSalaryTypeFieldRouter.DELETE("deleteWcSalaryTypeField", wcSalaryTypeFieldApi.DeleteWcSalaryTypeField) // 删除工资类型款项
		wcSalaryTypeFieldRouter.DELETE("deleteWcSalaryTypeFieldByIds", wcSalaryTypeFieldApi.DeleteWcSalaryTypeFieldByIds) // 批量删除工资类型款项
		wcSalaryTypeFieldRouter.PUT("updateWcSalaryTypeField", wcSalaryTypeFieldApi.UpdateWcSalaryTypeField)    // 更新工资类型款项
	}
	{
		wcSalaryTypeFieldRouterWithoutRecord.GET("findWcSalaryTypeField", wcSalaryTypeFieldApi.FindWcSalaryTypeField)        // 根据ID获取工资类型款项
		wcSalaryTypeFieldRouterWithoutRecord.GET("getWcSalaryTypeFieldList", wcSalaryTypeFieldApi.GetWcSalaryTypeFieldList)  // 获取工资类型款项列表
	}
	{
	    wcSalaryTypeFieldRouterWithoutAuth.GET("getWcSalaryTypeFieldPublic", wcSalaryTypeFieldApi.GetWcSalaryTypeFieldPublic)  // 获取工资类型款项列表
	}
}
