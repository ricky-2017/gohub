package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	r := gin.New()

	// 中间件注册
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "rubio",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {
		// 获取标头信息的Accpet信息
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// html返回
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认json
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"massage": "请求资源不存在",
			})
		}
	})

	r.Run()
}
