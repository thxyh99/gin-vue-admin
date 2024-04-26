package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffBankRouter struct {
}

// InitWcStaffBankRouter 初始化 银行卡信息 路由信息
func (s *WcStaffBankRouter) InitWcStaffBankRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffBankRouter := Router.Group("wcStaffBank").Use(middleware.OperationRecord())
	wcStaffBankRouterWithoutRecord := Router.Group("wcStaffBank")
	wcStaffBankRouterWithoutAuth := PublicRouter.Group("wcStaffBank")

	var wcStaffBankApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffBankApi
	{
		wcStaffBankRouter.POST("createWcStaffBank", wcStaffBankApi.CreateWcStaffBank)   // 新建银行卡信息
		wcStaffBankRouter.DELETE("deleteWcStaffBank", wcStaffBankApi.DeleteWcStaffBank) // 删除银行卡信息
		wcStaffBankRouter.DELETE("deleteWcStaffBankByIds", wcStaffBankApi.DeleteWcStaffBankByIds) // 批量删除银行卡信息
		wcStaffBankRouter.PUT("updateWcStaffBank", wcStaffBankApi.UpdateWcStaffBank)    // 更新银行卡信息
	}
	{
		wcStaffBankRouterWithoutRecord.GET("findWcStaffBank", wcStaffBankApi.FindWcStaffBank)        // 根据ID获取银行卡信息
		wcStaffBankRouterWithoutRecord.GET("getWcStaffBankList", wcStaffBankApi.GetWcStaffBankList)  // 获取银行卡信息列表
	}
	{
	    wcStaffBankRouterWithoutAuth.GET("getWcStaffBankPublic", wcStaffBankApi.GetWcStaffBankPublic)  // 获取银行卡信息列表
	}
}
