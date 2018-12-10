package log

import (
	"os"
	"log"
	"fmt"
	"time"
)




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

// 记录，并且显示，字符串
func PrintLogger(s string, print bool)  {
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

	file, _ := os.OpenFile(logDir+"/Logger_"+t11 +".log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
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
