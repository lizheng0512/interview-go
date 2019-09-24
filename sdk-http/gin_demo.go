package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	group := engine.Group("auth", func(c *gin.Context) {
		fmt.Println("auth")
		param, _ := c.GetQuery("username")
		if param != "lizheng" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"msg": "Unauthorized."})
		}
	})
	group.GET("hello",
		func(c *gin.Context) {
			fmt.Println("1")
			c.String(http.StatusOK, "123")
		}, func(c *gin.Context) {
			fmt.Println("2")
			jsonMap := map[string]interface{}{"msg": "你好2"}
			c.JSON(http.StatusInternalServerError, jsonMap)
		}, func(c *gin.Context) {
			//time.Sleep(10e9)
			fmt.Println("3")
		})
	fmt.Println("gin服务启动失败！", engine.Run("0.0.0.0:8081"))
}
