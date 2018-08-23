package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

func (c * UserController)Get()  {
	//c.Data["name"] = "d"
	fmt.Println("get:",c.GetString("name") )

	c.Ctx.WriteString("hello get")
}

func (c * UserController)Post()  {
	//c.Data["name"] = "d"
	fmt.Println("post")
	fmt.Println("get:",c.GetString("name") )


	var ob DDT
	var err error
	fmt.Println("",string(c.Ctx.Input.RequestBody))

	// 从json数据中解析到结构中
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err == nil {
		fmt.Println("ddt:", ob.Name)
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()

	c.Ctx.WriteString("hello post")
}

type DDT struct {
	Name string
}
