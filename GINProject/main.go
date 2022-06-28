package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "你好GIN")
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
