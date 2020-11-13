package Action

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type loginData struct {
	Username string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	fmt.Println("user:",c.PostForm("username"))		// 获取form的数据

	var data loginData
	err:= c.BindJSON(&data)					// 获取json数据
	if err == nil{
		fmt.Println("pwd:", data.Password,data.Username)
	}else{
		fmt.Println("",err.Error())
	}

	//c.JSON(200, gin.H{
	//	"message": "pong",
	//})


	reJson :=  gin.H{}
	reJson["code"] = 20000
	reJson["data"] = "token"
	reJson["message"] = data.Username + " 登录成功"
	//c.JSON( fiber.Map{"code":"777", "data": "" , "message":"这个bug" })
	fmt.Printf("reJson: %v\n", reJson)
	c.JSON(200,reJson)

}

type loginInfo struct {
	Roles []string `json:"roles"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}

func Info(c *gin.Context) {
	//fmt.Println("user:",c.PostForm("username"))		// 获取form的数据
	//
	//var data loginData
	//err:= c.BindJSON(&data)					// 获取json数据
	//if err == nil{
	//	fmt.Println("pwd:", data.Password,data.Username)
	//}else{
	//	fmt.Println("",err.Error())
	//}

	//c.JSON(200, gin.H{
	//	"message": "pong",
	//})
	//var info loginInfo
	//info.Avatar = ""
	//info.Name = "name"
	//info.Roles = ["admin","view"]

	roles := []string{"admin","view"}
	info:= loginInfo{Roles:roles, Name:"no name", Avatar:"" }


	reJson :=  gin.H{}
	reJson["code"] = 20000
	reJson["data"] = info
	//reJson["data"] = gin.H{"roles":roles, "name":"no name", "avatar":""}
	reJson["message"] = " 用户信息"
	//c.JSON( fiber.Map{"code":"777", "data": "" , "message":"这个bug" })
	fmt.Printf("reJson: %v\n", reJson)
	c.JSON(200,reJson)

}

func Logout(c *gin.Context) {
	reJson :=  gin.H{}
	reJson["code"] = 20000
	reJson["data"] = ""
	reJson["message"] = " 用户登出"
	fmt.Printf("reJson: %v\n", reJson)
	c.JSON(200,reJson)
}