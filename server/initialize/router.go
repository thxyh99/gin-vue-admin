package initialize

import (
	"net/http"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	InstallPlugin(Router)
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)})

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{

		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)
		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)

	}
	{
		weChatRouter := router.RouterGroupApp.WeChat
		weChatRouter.InitWcDepartmentRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcPositionRouter(PrivateGroup, PublicGroup)

		weChatRouter.InitWcStaffJobRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcStaffEducationRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcStaffBankRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcStaffContactRouter(PrivateGroup, PublicGroup)

		weChatRouter.InitWcStaffRouter(PrivateGroup, PublicGroup)

		weChatRouter.InitWcStaffPositionRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcStaffDepartmentRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcStaffAgreementRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcRankRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcFileRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcStaffSocialRouter(PrivateGroup, PublicGroup)

		weChatRouter.InitWcSalaryRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcStaffSalaryRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcSalaryTemplateRouter(PrivateGroup, PublicGroup)
		weChatRouter.InitWcSalaryTypeFieldRouter(PrivateGroup, PublicGroup)

	}
	{
		employeeRouter := router.RouterGroupApp.Employee
		employeeRouter.InitWcStaffEmploymentApplicationRouter(PrivateGroup, PublicGroup)
		employeeRouter.InitWcStaffPassApplicationRouter(PrivateGroup, PublicGroup)
		employeeRouter.InitWcStaffAdjustlevelApplicationRouter(PrivateGroup, PublicGroup)
		employeeRouter.InitWcStaffTransferApplicationRouter(PrivateGroup, PublicGroup)
		employeeRouter.InitWcStaffLeaveApplicationRouter(PrivateGroup, PublicGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
