package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"web_gin/MiddleWare"
	"web_gin/MiddleWare/zLog"
	"web_gin/MySql"
)

func main() {
	// 记录到文件。
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)			// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	zLog.CheckLogDir()
	zLog.ShowLog = true

	// 链接数据库
	MySql.DataBaseEngine = MySql.InitDataBase()
	MySql.InitSycTables()
	if MySql.DataBaseEngine == nil{
		fmt.Println("-----------数据库启动错误，无法启动服务器----------------")
		os.Exit(0)
	}


	r := gin.Default()
	r.Use(zLog.LoggerToFile())
	r.Use(MiddleWare.Cors()) // 允许使用跨域请求  全局中间件
	Routes(r)

	r.RunTLS("0.0.0.0:8097","server.crt", "server.key") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
