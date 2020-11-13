package Action

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(200, "Hello Welcome to my site !!")
}
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
//http://localhost:8097/get?ss=1&name=zsw
func Get(c *gin.Context)  {
	ss:= c.Query("ss")		// 获取get的参数
	name:= c.Query("name")		// 获取get的参数
	fmt.Println(ss +"      "+name )
}