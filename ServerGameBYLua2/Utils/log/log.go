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

// 输出字符串日志，带显示出控制台
func PrintLogger(str string) {
	_logger(str)
	if ShowLog{
		fmt.Println("Log："+str)
	}
}
// 输出字符串日志
func WriteLogger(str string) {
	_logger(str)
}

// 输出字符串日志
func WritefLogger(format string, a ...interface{}) {
	str:= fmt.Sprintf(format,a...)
	_logger(str)
}

// 格式化字符串日志，带显示出控制台
func PrintfLogger(format string, a...interface{})  {
	str:= fmt.Sprintf(format,a...)
	_logger(str)
	if ShowLog{
		fmt.Println("Log："+str)
	}
}

// 内部函数
func _logger(str string)  {
	file, _ := os.OpenFile("Logger.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	logger := log.New(file, "", log.LstdFlags)
	logger.Println("[Log:]", str)
}
