package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

// 日志记录到文件中间件
type LogStr struct {
	Time        string
	Code        int
	Method      string
	Url         string
	LatencyTime string
	Ip          string
}

func LoggerToFile() func(c *gin.Context) {
	return func(c *gin.Context) {

		startTime := time.Now()                              // 开始时间
		timeStr := time.Now().Format("2006-01-02 15:04:05")  // 开始时间string
		c.Next()                                             // 处理请求
		endTime := time.Now()                                // 结束时间
		latencyTime := endTime.Sub(startTime).String() // 执行时间

		statusCode := c.Writer.Status()// 状态码
		reqMethod := c.Request.Method// 请求方式
		reqUrl := c.Request.RequestURI// 请求路由
		clientIP := c.ClientIP()// 请求IP


		//str := LogStr{Time: timeStr, Code: statusCode, Method: reqMethod, Url: reqUrl, LatencyTime: latencyTime, Ip: clientIP}
		//jsonStr, _ := json.Marshal(str)
		//WriteWithIo(string(jsonStr) + "\n")
		str :=  fmt.Sprintf("%s	|%d	 |%s	|%s		|%s		|%s  \n" ,timeStr,statusCode,reqMethod,reqUrl,latencyTime, clientIP)
		WriteWithIo(str )

	}
}

func WriteWithIo(content string) {

	// 根据日期进行分文件
	t1 := time.Now()
	timeShow := t1.Format("2006-01-02")

	filePath := "Logs" + "/Logger_" + timeShow + ".log"

	fileObj, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	if _, err := io.WriteString(fileObj, content); err == nil {
		//fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.",content)
	}

}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 检查日志的目录
func CheckLogDir() {
	// 判断一下目录， 如果没有目录就创建出来
	logDir := "Logs"
	exist, err := PathExists(logDir)
	if err != nil {
		fmt.Printf("get Logs/ dir error![%v]\n", err)
		return
	}
	if !exist {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir %s failed![%v]\n", logDir, err)
		} else {
			fmt.Printf("mkdir %s success!\n", logDir)
		}
	}
	fmt.Println("日志目录检查完成")
}

