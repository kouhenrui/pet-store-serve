package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/service/sys"
	util "pet-store-serve/src/utils"
)

var (
	systemService sys.SysInter = &sys.SysService{}
)

func MenuList(c *gin.Context) {
	var list reqDto.MenuList
	if err := c.Bind(&list); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &list))
		return
	}
	resList, err := systemService.MenuList(list)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", resList)
	return
}
func GetMenuList(c *gin.Context) {
	resList, err := systemService.GetMenuList()
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", resList)
	return
}
func MonitorData(c *gin.Context) {
	fmt.Println(c.Request, "**************************")
	body := c.Request.Body
	//var list reqDto.MenuList
	//if err := c.Bind(&list); err != nil {
	//	c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &list))
	//	return
	//}
	//resList, err := systemService.MenuList(list)
	//if err != nil {
	//	c.Error(err)
	//	return
	//}
	c.Set("res", body)
	return
}
