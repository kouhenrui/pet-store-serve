package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-store-serve/src/controller/account"
	admin "pet-store-serve/src/controller/admin"
	common "pet-store-serve/src/controller/common"
	"pet-store-serve/src/controller/pet"
	"pet-store-serve/src/controller/system"
	"pet-store-serve/src/middleware"
)

/**
 * @ClassName api
 * @Description TODO
 * @Author khr
 * @Date 2023/7/29 14:18
 * @Version 1.0
 */

func InitApi(route *gin.Engine) {
	api := route.Group("/v1/api")
	{
		api.StaticFS("/static/file", http.Dir("static"))
		api.GET("/captcha", common.GetCaptcha, middleware.LimiterMiddleWare())
		api.POST("/upload/file", common.UploadFile)
		api.POST("/uploads/files", common.UploadFiles)
		api.POST("/upload/video", common.UploadVideo)
		api.GET("/download", common.DownloadFile)
		monitorModule := api.Group("/monitor")
		{
			monitorModule.POST("/listen", system.MonitorData)
		}
		sysModule := api.Group("/sys")
		{
			sysModule.GET("/menu/list", system.GetMenuList)
		}
		authModule := api.Group("/auth")

		{
			authModule.POST("/login", admin.Login)
			authModule.GET("/info", admin.Info)
			//authModule.GET("/refresh/token",admin.RefeshToken)
			authModule.POST("/register", admin.Register)
			authModule.POST("/logout", admin.LogOut)
			authModule.GET("/account/profile", admin.AccountProfile)
			authModule.PUT("/restet/pwd/self", admin.ResetPwdBySelf)

		}
		petModule := api.Group("/pet")
		{
			petModule.POST("/type/list", pet.PetTypeList)
			petModule.GET("/type/info", pet.PetTypeInfo)
		}

		accModule := api.Group("/account")
		{
			accModule.GET("/refresh/token", account.RefreshToken)
		}

	}

}
