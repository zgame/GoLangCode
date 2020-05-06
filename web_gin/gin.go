package main

import (
	"./Action"
	"./MiddleWare"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(MiddleWare.Cors()) // 允许使用跨域请求  全局中间件
	Routes(r)

	r.Run("0.0.0.0:8097") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Routes(r *gin.Engine) {
	r.GET("/ping", Action.Ping)
	r.GET("/get", Action.Get)
	//r.GET("/user/login", Action.Login)
	r.POST("/user/login", Action.Login)
	r.GET("/user/info", Action.Info)
	r.POST("/user/logout", Action.Logout)
}
