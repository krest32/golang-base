package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Gin渲染
func main() {
	r := gin.Default()
	// 加载该文件夹下所有的模版
	r.LoadHTMLGlob("src/day12-gin/templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8080")
}
