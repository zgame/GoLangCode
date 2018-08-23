package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{})

	beego.Get("/user/get",func(ctx *context.Context){
		fmt.Println("",ctx.Request)
		ctx.Output.Body([]byte("hello 2222222222world"))
	})
}
