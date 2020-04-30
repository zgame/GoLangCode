package Action

import (
	"github.com/gofiber/fiber"
	"log"
)

type Person struct {
	Name string `json:"username" xml:"name" form:"name" query:"name"`
	Pass string `json:"password" xml:"pass" form:"pass" query:"pass"`
}

func Login(c *fiber.Ctx)  {
	p := new(Person)

	if err := c.BodyParser(p); err != nil {
		log.Fatal(err)
	}

	log.Println(p.Name) // post 获取json数据
	log.Println(p.Pass)
}
