package route

import (
	"github.com/gin-gonic/gin"
	admin "pet-store-serve/src/controller/admin"
	common "pet-store-serve/src/controller/common"
)

/**
 * @ClassName api
 * @Description TODO
 * @Author khr
 * @Date 2023/7/29 14:18
 * @Version 1.0
 */

func InitApi(route *gin.Engine) {
	api := route.Group("/api")
	{
		api.GET("/captcha", common.GetCaptcha)
		api.POST("/upload/file", common.UploadFile)
		api.POST("/uploads/files", common.UploadFiles)
		api.POST("/upload/video", common.UploadVideo)
		api.GET("/download", common.DownloadFile)

		authModule := api.Group("/auth")
		{
			authModule.POST("/login", admin.Login)
			authModule.GET("/info", admin.Info)
			authModule.POST("/register", admin.Register)
			authModule.GET("/logout", admin.LogOut)
			authModule.GET("/account/profile", admin.AccountProfile)
			authModule.PUT("/restet/pwd/self", admin.ResetPwdBySelf)
		}
	}

}
