package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	获取querystring参数
	querystring指的是URL中?后面携带的参数，
	例如：/user/search?username=小王子&address=沙河

	获取请求的querystring参数的方法如下：
 */
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}
