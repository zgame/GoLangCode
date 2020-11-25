package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"log"
	"os"
	"web_gin/MiddleWare"
	"web_gin/MiddleWare/aliPay"
	"web_gin/MiddleWare/wxPay"
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

	wxPay.Init()
	aliPay.Init()


	//fmt.Println("------------------首先读取命令行参数---------------------------")
	//https := flag.Int("http", 0, "")
	//flag.Parse()
	//if *https > 0 {
		zLog.PrintLogger("===========启动 https 服务器=============")
		// 读取配置文件
		f, err := ini.Load("Setting.ini")
		if err != nil {
			fmt.Println("ini配置文件出错！", err)
			log.Fatal(err)
		}
		ServerPort := f.Section("author").Key("MyPort").Value()
		r.RunTLS(":"+ServerPort,"Crt/1_shop.portia.xyz_bundle.crt", "Crt/2_shop.portia.xyz.key") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//}else {
	//	zLog.PrintLogger("===========启动 http 服务器=============")
	//	r.Run(":8098") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//}
}
