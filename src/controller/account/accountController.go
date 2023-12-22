package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/service/account"
	util "pet-store-serve/src/utils"
	"strings"
	"time"
)

var (
	err            error
	accountService account.AccountInter = &account.AccountService{}
)

func AccountLogin(c *gin.Context) {
	var accountLogin reqDto.AccountLogin
	if err = c.Bind(&accountLogin); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &accountLogin))
		return
	}
	ip := c.GetString("ip")
	t, err := accountService.Login(accountLogin, ip)
	if err != nil {
		c.Error(err)
		return
	}
	// 将Token设置到Cookie中
	cookieName := "token:" + accountLogin.UserName
	c.SetCookie(cookieName, t.Token, int(time.Hour*24), c.GetString("reqUrl"), ip, false, true)
	c.Set("res", t)
	return
}
func RefreshToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	oldToken := strings.Split(header, " ")[1]
	ip := c.GetString("ip")
	t, err := accountService.RefreshToken(oldToken, ip)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", t)
	return
}
