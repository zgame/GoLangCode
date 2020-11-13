package Action

import (
	"github.com/gin-gonic/gin"
	"time"
)

// 错误消息， 一个本地记录， 一个返回客户端
func Error(msgStr string ,c *gin.Context)  {
	//zLog.PrintfLogger(msgStr)
	c.JSON(200, gin.H{"Error": msgStr})
}

// 获取当前时间
func getTime()  string{
	return  time.Now().Format("2006-01-02 15:04:05")
}