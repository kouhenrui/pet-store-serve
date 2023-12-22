package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/dto/resDto"
	"pet-store-serve/src/msg"
	"pet-store-serve/src/service/account"
	util "pet-store-serve/src/utils"
	"time"
)

var (
	adminService account.AccountInter = &account.AccountService{}
	err          error
)

// @Summary	登录接口
// @Produce	json
// @Tags		account
// @Param		request	body reqDto.AccountLogin true	"参数"
// @Success	200 {object} comDto.ResponseData
// @Router		/auth/login [post]
func Login(c *gin.Context) {
	var adminLogin reqDto.AccountLogin
	if err = c.Bind(&adminLogin); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &adminLogin))
		return
	}
	ip := c.GetString("ip")
	t, err := adminService.Login(adminLogin, ip)
	if err != nil {
		c.Error(err)
		return
	}
	// 将Token设置到Cookie中
	cookieName := "token:" + adminLogin.UserName
	//fmt.Println(cookieName, "000000000000000000000000", t)
	c.SetCookie(cookieName, t.Token, int(time.Hour*24), c.GetString("reqUrl"), ip, false, true)
	c.Set("res", t)
	return
}

// @Summary	登出接口
// @Produce	json
// @Tags		account
// @Success	200 string ok
// @Router		/auth/logout [get]
func LogOut(c *gin.Context) {

	//c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Set("res", "ok")
	return
	//var adminRegister reqDto.AddAccount
	//if err = c.ShouldBindJSON(&adminRegister); err != nil {
	//	c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &adminRegister))
	//	return
	//}

}

// @Summary	详情接口
// @Produce	json
// @Tags		account
// @Success	200 {object} comDto.ResponseData
// @Router		/auth/info [get]
func Info(c *gin.Context) {
	id := c.GetUint("user_id")
	//fmt.Println(id, "id")
	var info = &resDto.AccountInfo{}
	info, err = adminService.Info(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", info)
	return
}
func Register(c *gin.Context) {
	var addAccount reqDto.AddAccount
	if err = c.ShouldBindJSON(&addAccount); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &addAccount))
		return
	}
	err = adminService.Register(addAccount)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", msg.ADD_SUCCESS)

}

func AccountProfile(c *gin.Context) {
	id := c.GetUint("user_id")
	fmt.Println(id, "id")
	var info = &resDto.AccountInfo{}
	info, err = adminService.Info(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", info)
	return
}

func ResetPwdBySelf(c *gin.Context) {
	var reset reqDto.UpdateAccount
	if err = c.ShouldBindJSON(&reset); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &reset))
		return
	}
	err = adminService.ResetPwdBySelf(reset)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", msg.PWD_CHANGE_SUCCESS)
	return
}
