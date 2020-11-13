package Action

import (
	"github.com/gofiber/fiber"
	"web_portia_shop/MySql"
)

func GetUserBuyList(c *fiber.Ctx) {
	// 获取参数
	uid := c.Query("uid")
	println(uid)
	MySql.GetUserInfoData()

	c.Send("Hello, World!")
	c.JSON( fiber.Map{"hello":"me"})
}

