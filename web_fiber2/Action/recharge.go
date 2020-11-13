package Action

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func Recharge(c *fiber.Ctx) {
	fmt.Println("---------------充值表---------------")
	limit := c.Query("limit")
	page := c.Query("page")
	uid := c.Query("uid")
	channel := c.Query("channel")
	serverid := c.Query("serverid")
	starttime := c.Query("starttime")
	endtime := c.Query("endtime")
	time := c.Query("time")    // 前端发过来的是一个数组[starttime, endtime] ， 备用

	fmt.Println("",limit)
	fmt.Println("",page)
	fmt.Println("",uid)
	fmt.Println("",channel)
	fmt.Println("",serverid)
	fmt.Println("",starttime)
	fmt.Println("",endtime)
	fmt.Println("",time)


	reJson := fiber.Map{}
	reJson["code"] = 20000
	reJson["data"] = fiber.Map{"items":[]string{} , "total" : 1}
	reJson["message"] = " 充值"
	fmt.Printf("reJson: %v\n", reJson)
	c.JSON(reJson)
}