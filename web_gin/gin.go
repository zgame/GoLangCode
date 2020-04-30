package main

import (
	"./Action"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Routes(r)
	r.Run("0.0.0.0:8097") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Routes(r *gin.Engine)  {
	r.GET("/ping", Action.Ping)
	r.GET("/get", Action.Get)
}