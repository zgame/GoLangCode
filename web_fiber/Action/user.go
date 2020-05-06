package Action

import (
	"fmt"
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

	fmt.Println("---------------登录申请---------------")
	fmt.Println(p.Name) // post 获取json数据
	fmt.Println(p.Pass)

	reJson := fiber.Map{}
	reJson["code"] = 20000
	reJson["data"] = "token"
	reJson["message"] = p.Name + " 登录成功"
	//c.JSON( fiber.Map{"code":"777", "data": "" , "message":"这个bug" })
	fmt.Printf("reJson: %v\n", reJson)
	c.JSON(reJson)
}

func Info(c *fiber.Ctx)  {
	//p := new(Person)
	//
	//if err := c.BodyParser(p); err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(p.Name) // post 获取json数据
	//log.Println(p.Pass)


	fmt.Println("---------------获取info---------------")
	tokens := c.Query("token")
	fmt.Println("tokens : ",tokens)

	roles := []string{"admin","view"}

	reJson := fiber.Map{}
	reJson["code"] = 20000
	reJson["data"] = fiber.Map{"roles":roles, "name":"no name", "avatar":""}
	reJson["message"] = "用户信息"
	//c.JSON( fiber.Map{"code":"777", "data": "" , "message":"这个bug" })
	fmt.Printf("reJson: %v\n", reJson)
	c.JSON(reJson)
}


func Logout(c *fiber.Ctx) {
	fmt.Println("---------------登出---------------")
	reJson := fiber.Map{}
	reJson["code"] = 20000
	reJson["data"] = ""
	reJson["message"] = " 用户登出"
	fmt.Printf("reJson: %v\n", reJson)
	c.JSON(reJson)
}