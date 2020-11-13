package Action

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Cookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "gin cookie ok", 3600, "/", "localhost", true, true)
	}

	fmt.Printf("Cookie value: %s \n", cookie)
	c.JSON(200,  gin.H{
		"cookie": cookie,
	})
}
