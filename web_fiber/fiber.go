package main

import (
	"./Action"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	Routes(app)
	app.Listen(3000)
}

func Routes(app *fiber.App) {
	app.Get("/ping", Action.Ping)
	app.Get("/get", Action.Get)
}
