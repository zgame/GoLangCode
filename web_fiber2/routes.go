package main

import (
	"github.com/gofiber/fiber"
	"web_portia_shop/Action"
)

// 路由
func Routes(app *fiber.App) {
	app.Get("/", Action.Index)
	app.Get("/ping", Action.Ping)
	app.Get("/get", Action.Get)
	app.Post("/user/login", Action.Login)
	app.Get("/user/info", Action.Info)
	app.Post("/user/logout", Action.Logout)

	app.Get("/recharge/list", Action.Recharge)

	// portia shop
	app.Get("/portia_shop/buy_list", Action.GetUserBuyList)


}
