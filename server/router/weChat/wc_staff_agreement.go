package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffAgreementRouter struct {
}

// InitWcStaffAgreementRouter 初始化 合同信息 路由信息
func (s *WcStaffAgreementRouter) InitWcStaffAgreementRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffAgreementRouter := Router.Group("wcStaffAgreement").Use(middleware.OperationRecord())
	wcStaffAgreementRouterWithoutRecord := Router.Group("wcStaffAgreement")
	wcStaffAgreementRouterWithoutAuth := PublicRouter.Group("wcStaffAgreement")

	var wcStaffAgreementApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffAgreementApi
	{
		wcStaffAgreementRouter.POST("createWcStaffAgreement", wcStaffAgreementApi.CreateWcStaffAgreement)   // 新建合同信息
		wcStaffAgreementRouter.DELETE("deleteWcStaffAgreement", wcStaffAgreementApi.DeleteWcStaffAgreement) // 删除合同信息
		wcStaffAgreementRouter.DELETE("deleteWcStaffAgreementByIds", wcStaffAgreementApi.DeleteWcStaffAgreementByIds) // 批量删除合同信息
		wcStaffAgreementRouter.PUT("updateWcStaffAgreement", wcStaffAgreementApi.UpdateWcStaffAgreement)    // 更新合同信息
	}
	{
		wcStaffAgreementRouterWithoutRecord.GET("findWcStaffAgreement", wcStaffAgreementApi.FindWcStaffAgreement)        // 根据ID获取合同信息
		wcStaffAgreementRouterWithoutRecord.GET("getWcStaffAgreementList", wcStaffAgreementApi.GetWcStaffAgreementList)  // 获取合同信息列表
	}
	{
	    wcStaffAgreementRouterWithoutAuth.GET("getWcStaffAgreementPublic", wcStaffAgreementApi.GetWcStaffAgreementPublic)  // 获取合同信息列表
	}
}
