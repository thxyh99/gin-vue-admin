package ach

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AchTestRouter struct {
}

// InitAchTestRouter 初始化 绩效测试 路由信息
func (s *AchTestRouter) InitAchTestRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	achTestRouter := Router.Group("achTest").Use(middleware.OperationRecord())
	achTestRouterWithoutRecord := Router.Group("achTest")
	achTestRouterWithoutAuth := PublicRouter.Group("achTest")

	var achTestApi = v1.ApiGroupApp.AchApiGroup.AchTestApi
	{
		achTestRouter.POST("createAchTest", achTestApi.CreateAchTest)   // 新建绩效测试
		achTestRouter.DELETE("deleteAchTest", achTestApi.DeleteAchTest) // 删除绩效测试
		achTestRouter.DELETE("deleteAchTestByIds", achTestApi.DeleteAchTestByIds) // 批量删除绩效测试
		achTestRouter.PUT("updateAchTest", achTestApi.UpdateAchTest)    // 更新绩效测试
	}
	{
		achTestRouterWithoutRecord.GET("findAchTest", achTestApi.FindAchTest)        // 根据ID获取绩效测试
		achTestRouterWithoutRecord.GET("getAchTestList", achTestApi.GetAchTestList)  // 获取绩效测试列表
	}
	{
	    achTestRouterWithoutAuth.GET("getAchTestPublic", achTestApi.GetAchTestPublic)  // 获取绩效测试列表
	}
}
