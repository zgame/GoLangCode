package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"web_gin/Action"
	"web_gin/MiddleWare"
)

func main() {
	// 记录到文件。
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)			// 如果需要同时将日志写入文件和控制台，请使用以下代码。

	r := gin.Default()
	r.Use(MiddleWare.Cors()) // 允许使用跨域请求  全局中间件
	Routes(r)

	r.Run("0.0.0.0:8097") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Routes(r *gin.Engine) {
	r.GET("/ping", Action.Ping)
	r.GET("/", Action.Welcome)

	r.GET("/get", Action.Get)
	//r.GET("/user/login", Action.Login)


	r.POST("/user/login", Action.Login)
	r.GET("/user/info", Action.Info)
	r.POST("/user/logout", Action.Logout)

	r.GET("/cookie", Action.Cookie)
}
