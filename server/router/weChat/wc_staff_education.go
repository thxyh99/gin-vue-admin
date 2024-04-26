package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcStaffEducationRouter struct {
}

// InitWcStaffEducationRouter 初始化 学历信息 路由信息
func (s *WcStaffEducationRouter) InitWcStaffEducationRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wcStaffEducationRouter := Router.Group("wcStaffEducation").Use(middleware.OperationRecord())
	wcStaffEducationRouterWithoutRecord := Router.Group("wcStaffEducation")
	wcStaffEducationRouterWithoutAuth := PublicRouter.Group("wcStaffEducation")

	var wcStaffEducationApi = v1.ApiGroupApp.WeChatApiGroup.WcStaffEducationApi
	{
		wcStaffEducationRouter.POST("createWcStaffEducation", wcStaffEducationApi.CreateWcStaffEducation)   // 新建学历信息
		wcStaffEducationRouter.DELETE("deleteWcStaffEducation", wcStaffEducationApi.DeleteWcStaffEducation) // 删除学历信息
		wcStaffEducationRouter.DELETE("deleteWcStaffEducationByIds", wcStaffEducationApi.DeleteWcStaffEducationByIds) // 批量删除学历信息
		wcStaffEducationRouter.PUT("updateWcStaffEducation", wcStaffEducationApi.UpdateWcStaffEducation)    // 更新学历信息
	}
	{
		wcStaffEducationRouterWithoutRecord.GET("findWcStaffEducation", wcStaffEducationApi.FindWcStaffEducation)        // 根据ID获取学历信息
		wcStaffEducationRouterWithoutRecord.GET("getWcStaffEducationList", wcStaffEducationApi.GetWcStaffEducationList)  // 获取学历信息列表
	}
	{
	    wcStaffEducationRouterWithoutAuth.GET("getWcStaffEducationPublic", wcStaffEducationApi.GetWcStaffEducationPublic)  // 获取学历信息列表
	}
}
