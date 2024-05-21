package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WcFileRouter struct {
}

// InitWcFileRouter 初始化 文件管理 路由信息
func (s *WcFileRouter) InitWcFileRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wcFileRouter := Router.Group("wcFile").Use(middleware.OperationRecord())
	wcFileRouterWithoutRecord := Router.Group("wcFile")
	wcFileRouterWithoutAuth := PublicRouter.Group("wcFile")

	var wcFileApi = v1.ApiGroupApp.WeChatApiGroup.WcFileApi
	{
		wcFileRouter.POST("createWcFile", wcFileApi.CreateWcFile)             // 新建文件管理
		wcFileRouter.DELETE("deleteWcFile", wcFileApi.DeleteWcFile)           // 删除文件管理
		wcFileRouter.DELETE("deleteWcFileByIds", wcFileApi.DeleteWcFileByIds) // 批量删除文件管理
		wcFileRouter.PUT("updateWcFile", wcFileApi.UpdateWcFile)              // 更新文件管理
	}
	{
		wcFileRouterWithoutRecord.GET("findWcFile", wcFileApi.FindWcFile)       // 根据ID获取文件管理
		wcFileRouterWithoutRecord.GET("getWcFileList", wcFileApi.GetWcFileList) // 获取文件管理列表
	}
	{
		wcFileRouterWithoutAuth.GET("getWcFilePublic", wcFileApi.GetWcFilePublic) // 获取文件管理列表
	}
	{
		wcFileRouter.POST("upload", wcFileApi.UploadFile)         // 上传文件
		wcFileRouter.POST("getFileList", wcFileApi.GetFileList)   // 获取上传文件列表
		wcFileRouter.POST("deleteFile", wcFileApi.DeleteFile)     // 删除指定文件
		wcFileRouter.POST("editFileName", wcFileApi.EditFileName) // 编辑文件名或者备注
	}

}
