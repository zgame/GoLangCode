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

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
