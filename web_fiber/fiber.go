package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"io"
	"log"
	"os"
	"web_fiber/Action"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())		// 跨域设置
	//cfg := basicauth.Config{
	//	Users: map[string]string{
	//		"zsw":   "zsw",				// 用户名和密码
	//		"admin":  "123456",
	//	},
	//}
	//app.Use(basicauth.New(cfg))		// 认证

	//--------------------日志---------------------------------
	//app.Use(logger.New())				//日志
	file, _ := os.Create("fiber.log")
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		//TimeZone:   "America/New_York",
		//Output: os.Stdout ,
		Output: io.MultiWriter(file, os.Stdout) ,
	}))
	//--------------------报错---------------------------------
	cfg := recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}
	app.Use(recover.New(cfg))			// 报错


	Routes(app)
	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
		// => 404 "Not Found"
	})
	log.Fatal(app.Listen(8097))
}


// 路由
func Routes(app *fiber.App) {
	app.Get("/", Action.Index)
	app.Get("/ping", Action.Ping)
	app.Get("/get", Action.Get)
	app.Post("/user/login", Action.Login)
	app.Get("/user/info", Action.Info)
	app.Post("/user/logout", Action.Logout)

	app.Get("/recharge/list", Action.Recharge)


}
