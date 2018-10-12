package log

import (
	"os"
	"log"
	"fmt"
)

//--------------------------------------------------------------------------------------------------
// 错误日志处理
//--------------------------------------------------------------------------------------------------
var ShowLog bool
func CheckError(e error) bool{
	if e!=nil{
		file, _ := os.OpenFile("error.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		logger := log.New(file, "", log.LstdFlags|log.Llongfile)
		logger.Println("...error:...",e.Error())
		if ShowLog{
			fmt.Println("错误："+e.Error())
		}
		return true
	}
	return false
}


