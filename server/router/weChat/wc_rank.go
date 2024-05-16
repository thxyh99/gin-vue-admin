package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcRankRouter struct {
}

// InitWcRankRouter 初始化 职级管理 路由信息
func (s *WcRankRouter) InitWcRankRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wcRankRouter := Router.Group("wcRank").Use(middleware.OperationRecord())
	wcRankRouterWithoutRecord := Router.Group("wcRank")
	wcRankRouterWithoutAuth := PublicRouter.Group("wcRank")

	var wcRankApi = v1.ApiGroupApp.WeChatApiGroup.WcRankApi
	{
		wcRankRouter.POST("createWcRank", wcRankApi.CreateWcRank)             // 新建职级管理
		wcRankRouter.DELETE("deleteWcRank", wcRankApi.DeleteWcRank)           // 删除职级管理
		wcRankRouter.DELETE("deleteWcRankByIds", wcRankApi.DeleteWcRankByIds) // 批量删除职级管理
		wcRankRouter.PUT("updateWcRank", wcRankApi.UpdateWcRank)              // 更新职级管理
	}
	{
		wcRankRouterWithoutRecord.GET("findWcRank", wcRankApi.FindWcRank)                       // 根据ID获取职级管理
		wcRankRouterWithoutRecord.GET("getWcRankList", wcRankApi.GetWcRankList)                 // 获取职级管理列表
		wcRankRouterWithoutRecord.GET("getRankTypeList", wcRankApi.GetRankTypeList)             // 获取职级类型列表
		wcRankRouterWithoutRecord.GET("getRankListByRankType", wcRankApi.GetRankListByRankType) // 获取职级列表
	}
	{
		wcRankRouterWithoutAuth.GET("getWcRankPublic", wcRankApi.GetWcRankPublic) // 获取职级管理列表
	}
}
