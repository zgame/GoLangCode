package Action

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func Index(c *fiber.Ctx) {
	c.Send("Hello, " + c.Hostname())
}

func Ping(c *fiber.Ctx) {
	c.Send("Hello, World!")
	c.JSON( fiber.Map{"hello":"me"})
}

//http://localhost:3000/get?ss=1&name=zsw
func Get(c *fiber.Ctx)  {
	ss:= c.Query("ss")		// 获取get的参数
	name:= c.Query("name")		// 获取get的参数
	fmt.Println(ss +"      "+name )
}
