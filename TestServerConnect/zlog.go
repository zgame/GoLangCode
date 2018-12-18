package main

import (
	"os"
	"log"
	"fmt"
	"time"
)

func debugUid(uid int)  {
	file, _ := os.OpenFile("uid.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	logger := log.New(file, "", log.LstdFlags)
	logger.Println("UID:...",uid)
}

// 格式化字符串日志，带显示出控制台
func PrintfLogger(format string, a...interface{})  {
	str:= fmt.Sprintf(format,a...)
	_logger(str)

	fmt.Println("Log："+str)

}

// 内部函数
func _logger(str string)  {
	//go func() {
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

	// 根据日期进行分文件
	t1:=time.Now()
	t11 := t1.Format("2006-01-02")

	file, _ := os.OpenFile(logDir+"/Logger_"+t11 +".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	logger := log.New(file, "", log.LstdFlags)
	logger.Println("[Log:]", str)
	//}()

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

