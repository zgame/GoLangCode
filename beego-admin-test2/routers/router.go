package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	//admin.Run()
    beego.Router("/", &controllers.MainController{})
}
