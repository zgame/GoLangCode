package log

import (
	"os"
	"log"
	"fmt"
)
// 记录，并且显示，字符串
func PrintLogger(s string, print bool)  {
	file, _ := os.OpenFile("Log.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	logger := log.New(file, "", log.LstdFlags)
	logger.Println("[Log] ",s)
	if print {
		fmt.Println(s)
	}
}

// 记录，并且显示， 带格式的
func PrintfLogger(format string, a ...interface{})   {
	str:=fmt.Sprintf(format,a...)
	PrintLogger(str,true)
}

// 记录，但是不显示
func WritefLogger(format string, a ...interface{})   {
	str:=fmt.Sprintf(format,a...)
	PrintLogger(str,false)
}
