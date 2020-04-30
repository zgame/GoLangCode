package Action

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Get(c *gin.Context)  {
	fmt.Println("", c.Param("ss"))
}