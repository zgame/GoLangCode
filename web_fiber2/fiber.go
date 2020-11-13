package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"log"
	"net"
	"os"
	"web_portia_shop/MySql"
)



func main() {
	// 链接数据库
	MySql.DataBaseEngine = MySql.InitDataBase()
	if MySql.DataBaseEngine == nil{
		fmt.Println("-----------数据库启动错误，无法启动服务器----------------")
		os.Exit(0)
	}

	// 启动服务器
	app := fiber.New()
	app.Use(cors.New())		// 跨域设置

	// 日志
	CheckLogDir()
	app.Use(LoggerToFile())

	// 报错
	cfg := recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}
	app.Use(recover.New(cfg))

	//路由
	Routes(app)
	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
		// => 404 "Not Found"
	})

	// 端口
	//log.Fatal(app.Listen(8097))
	ln, _ := net.Listen("tcp", ":8098")
	cer, _:= tls.LoadX509KeyPair("server.crt", "server.key")
	ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})
	log.Fatal(app.Listener(ln))


}


