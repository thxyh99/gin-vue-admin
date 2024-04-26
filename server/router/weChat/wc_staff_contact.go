package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffContactRouter struct {
}

// InitWcStaffContactRouter 初始化 紧急联系人 路由信息
func (s *WcStaffContactRouter) InitWcStaffContactRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffContactRouter := Router.Group("wcStaffContact").Use(middleware.OperationRecord())
	wcStaffContactRouterWithoutRecord := Router.Group("wcStaffContact")
	wcStaffContactRouterWithoutAuth := PublicRouter.Group("wcStaffContact")

	var wcStaffContactApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffContactApi
	{
		wcStaffContactRouter.POST("createWcStaffContact", wcStaffContactApi.CreateWcStaffContact)   // 新建紧急联系人
		wcStaffContactRouter.DELETE("deleteWcStaffContact", wcStaffContactApi.DeleteWcStaffContact) // 删除紧急联系人
		wcStaffContactRouter.DELETE("deleteWcStaffContactByIds", wcStaffContactApi.DeleteWcStaffContactByIds) // 批量删除紧急联系人
		wcStaffContactRouter.PUT("updateWcStaffContact", wcStaffContactApi.UpdateWcStaffContact)    // 更新紧急联系人
	}
	{
		wcStaffContactRouterWithoutRecord.GET("findWcStaffContact", wcStaffContactApi.FindWcStaffContact)        // 根据ID获取紧急联系人
		wcStaffContactRouterWithoutRecord.GET("getWcStaffContactList", wcStaffContactApi.GetWcStaffContactList)  // 获取紧急联系人列表
	}
	{
	    wcStaffContactRouterWithoutAuth.GET("getWcStaffContactPublic", wcStaffContactApi.GetWcStaffContactPublic)  // 获取紧急联系人列表
	}
}
