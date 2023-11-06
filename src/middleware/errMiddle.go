package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pet-store-serve/src/dto/comDto"
	"runtime/debug"
)

/*
* @MethodName
* @Description 全局异常捕捉
* @Author khr
* @Date 2023/7/31 15:21
 */

func GlobalErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			fmt.Println("程序捕捉错误")
			if r := recover(); r != nil {
				fmt.Println("打印错误信息:", r)
				log.Println(string(debug.Stack()))
				// 打印错误堆栈信息
				debug.PrintStack()
				errorMessage := string(debug.Stack())
				fmt.Println("错误")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": errorMessage, "data": ""})
				//c.JSON(http.StatusInternalServerError, &comDto.ResponseData{
				//	Code:    http.StatusInternalServerError,
				//	Message: errorMessage,
				//	Data:    nil,
				//})
				//return
				c.Abort()
			}
		}()

		c.Next()
	}
}

/*
* @MethodName UnifiedResponseMiddleware
* @Description 统一返回正确和错误格式
* @Author khr
* @Date 2023/7/29 9:45
 */

func UnifiedResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			fmt.Println("出现错误", c.Errors)
			err := c.Errors.Last()
			//fmt.Println(err.Type, "定位最新的错误地址")
			errorMessage := err.Err.Error()
			//fmt.Println(c.Writer.Status(), "错误类型")
			c.JSON(http.StatusOK, &comDto.ResponseData{
				Code:    c.Writer.Status() | 460,
				Message: errorMessage,
				Data:    "",
			})
			return

		}
		fmt.Println("格式化返回")
		if c.Writer.Status() == http.StatusOK {
			data, exists := c.Get("res")
			if exists {
				c.JSON(http.StatusOK, &comDto.ResponseData{
					Code:    http.StatusOK,
					Message: "success",
					Data:    data,
				})
				return
			}
		}

	}
}
