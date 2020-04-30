package main

import (
	"./Action"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	Routes(app)
	app.Listen(8097)
}

func Routes(app *fiber.App) {
	app.Get("/ping", Action.Ping)
	app.Get("/get", Action.Get)
	app.Post("/user/login", Action.Login)
	//app.Get("/user/login", Action.Login)

}
